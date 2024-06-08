package server

import (
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type srv struct {
	Router  *mux.Router
	Log     *logrus.Logger
	DB      interface{}
	session *sessions.CookieStore
}

func NewServer(logger *logrus.Logger, db interface{}, session string) *srv {
	s := &srv{
		Router:  mux.NewRouter(),
		Log:     logger,
		DB:      db,
		session: sessions.NewCookieStore([]byte(session)),
	}
	s.InitRouter()
	return s
}

func (s *srv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *srv) InitRouter() {
	s.Router.HandleFunc("/", s.HandleStart())
	s.Router.HandleFunc("/search", s.HandleSearch())
	s.Router.HandleFunc("/login", s.HandleLogin())
	s.Router.HandleFunc("/createUser", s.HandleCreateUser())

	serial_root := s.Router.PathPrefix("/serial").Subrouter()
	serial_root.HandleFunc("/{id:[0-9]+}", s.HandleSerial())

	user_root := s.Router.PathPrefix("/user").Subrouter()
	user_root.Use(s.UserAuth)
	user_root.HandleFunc("/cabinet/{id:[0-9]+}", s.HandleUserCabinet())
	user_root.HandleFunc("/exit", s.HandleExit())
	user_root.HandleFunc("/addComment", s.HandleAddComment())
	user_root.HandleFunc("/updateComment", s.HandleUpdateComment())
	user_root.HandleFunc("/deleteComment", s.HandleDeleteComment())
	user_root.HandleFunc("/favourites", s.HandleFavourites())
	user_root.HandleFunc("/deleteFavourite", s.HandleDeleteFavourite())
	user_root.HandleFunc("/history", s.HandleHistory())
	user_root.HandleFunc("/clearHistory", s.HandleClearHistory())
	user_root.HandleFunc("/compareSerials", s.HandleCompareSerials())
	user_root.HandleFunc("/changeProfile", s.HandleUpdateProfile())

	admin_root := s.Router.PathPrefix("/admin").Subrouter()
	admin_root.Use(s.AdminAuth)
	admin_root.HandleFunc("/cabinet/{id:[0-9]+}", s.HandleAdminCabinet())
	admin_root.HandleFunc("/exit", s.HandleExit())
	admin_root.HandleFunc("/addSerial", s.HandleAddSerial())
	admin_root.HandleFunc("/changeSerial", s.HandleUpdateSerial())
	admin_root.HandleFunc("/deleteSerial", s.HandleDeleteSerial())
	admin_root.HandleFunc("/addSerialActor", s.HandleAddSerialActor())
	admin_root.HandleFunc("/addActor", s.HandleAddActor())
	admin_root.HandleFunc("/changeActor", s.HandleUpdateActor())
	admin_root.HandleFunc("/deleteActor", s.HandleDeleteActor())
	admin_root.HandleFunc("/addProducer", s.HandleAddProducer())
	admin_root.HandleFunc("/changeProducer", s.HandleUpdateProducer())
	admin_root.HandleFunc("/deleteProducer", s.HandleDeleteProducer())
	admin_root.HandleFunc("/showUsers", s.HandleShowUsers())
	admin_root.HandleFunc("/deleteUser", s.HandleDeleteUser())
	admin_root.HandleFunc("/grantAdmin", s.HandleGrantAdmin())
	admin_root.HandleFunc("/showStatistics", s.HandleShowStatistics())

}

func (s *srv) HandleExit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.session.Get(r, "sname")
		delete(session.Values, "user")
		delete(session.Values, "admin")
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (s *srv) HandleSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
		serials, err := ctrl.GetSerialByTitle(r.FormValue("search"))
		if err != nil {
			s.Log.Error(err)
			return
		}
		tmpl, _ := template.ParseFiles("templates/search.html")
		tmpl.Execute(w, serials)
	}
}

func (s *srv) HandleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
		serials, err := ctrl.GetSerials()
		if err != nil {
			log.Fatal(err)
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, serials)
	}
}

func (s *srv) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.session.Get(r, "sname")
		if session.Values["user"] != nil {
			http.Redirect(w, r, "/user/cabinet/0", http.StatusMovedPermanently)
			return
		} else if session.Values["admin"] != nil {
			http.Redirect(w, r, "/admin/cabinet/0", http.StatusMovedPermanently)
			return
		}
		if r.Method == http.MethodPost {
			s.Auth(w, r)
			return
		}
		s.loginTemplate(w, nil)
	}
}

func (s *srv) loginTemplate(w http.ResponseWriter, err error) {
	type loginErr struct {
		Err string
	}
	tmpl, _ := template.ParseFiles("templates/login.html")
	lerr := &loginErr{}
	switch err {
	case nil:
		lerr.Err = ""
	case controllers.ErrUserNotFound:
		lerr.Err = "Пользователя с таким<br>логином не существует"
	case controllers.ErrInvalidPass:
		lerr.Err = "Неверный пароль"
	}
	tmpl.Execute(w, lerr)
}

func (s *srv) Auth(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user, err := ctrl.AuthUser(login, password)
	if err != nil {
		s.loginTemplate(w, err)
		return
	}
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	session.Values[user.U_role] = user.U_id
	s.session.Save(r, w, session)

	if user.U_role == "user" {
		http.Redirect(w, r, "/user/cabinet/0", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/admin/cabinet/0", http.StatusMovedPermanently)
	}
}

func (s *srv) HandleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptUserCreation(w, r)
			return
		}
		s.creationTemplate(w, "")
	}
}

func (s *srv) creationTemplate(w http.ResponseWriter, err string) {
	type creationErr struct {
		Err string
	}
	tmpl, _ := template.ParseFiles("templates/user/create.html")
	cerr := &creationErr{Err: err}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptUserCreation(w http.ResponseWriter, r *http.Request) {
	user := &models.Users{
		U_name:    r.FormValue("name"),
		U_surname: r.FormValue("surname"),
		U_login:   r.FormValue("login"),
		U_gender:  r.FormValue("gender"),
		U_bdate:   r.FormValue("bdate"),
		U_role:    "user",
	}
	password1 := r.FormValue("password1")
	password2 := r.FormValue("password2")
	if password1 != password2 {
		s.creationTemplate(w, "Пароли не совпадают")
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		s.creationTemplate(w, "bcrypt internal error")
		return
	}
	user.SetPassword(string(pass))
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	err = ctrl.CreateUser(user)
	if err != nil {
		if err == models.ErrInvalidModel {
			s.creationTemplate(w, "Заполнены не все поля")
		} else if err == controllers.ErrUserExists {
			s.creationTemplate(w, "Пользователь с таким<br>логином уже существует")
		}
		return
	}
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	session.Values[user.U_role] = user.U_id
	s.session.Save(r, w, session)

	// update statistics
	s.UpdateStatisticUp(user)

	http.Redirect(w, r, "/user/cabinet/0", http.StatusMovedPermanently)
}

func (s *srv) UpdateStatisticUp(user *models.Users) {
	ctrl := controllers.NewStatisticCtrl(repositories.NewStatisticRepo(s.DB, s.Log))
	stat, err := ctrl.GetStatistic()
	if err != nil {
		s.Log.Error(err)
		return
	}
	if user.U_gender == "мужской" {
		stat.IncreaseGenderMale()
	} else {
		stat.IncreaseGenderFemale()
	}
	if user.U_role == "user" {
		stat.IncreaseRoleUser()
	} else {
		stat.IncreaseRoleAdmin()
	}
	age := user.GetAge()
	switch {
	case age <= 18:
		stat.IncreaseAge0_18()
	case age <= 30:
		stat.IncreaseAge19_30()
	case age <= 50:
		stat.IncreaseAge31_50()
	case age > 50:
		stat.IncreaseAge51_100()
	}
	err = ctrl.UpdateStatistic(stat)
	if err != nil {
		s.Log.Error(err)
	}
}

func (s *srv) UpdateStatisticDown(user *models.Users) {
	ctrl := controllers.NewStatisticCtrl(repositories.NewStatisticRepo(s.DB, s.Log))
	stat, err := ctrl.GetStatistic()
	if err != nil {
		s.Log.Error(err)
		return
	}
	if user.U_gender == "мужской" {
		stat.DecreaseGenderMale()
	} else {
		stat.DecreaseGenderFemale()
	}
	if user.U_role == "user" {
		stat.DecreaseRoleUser()
	} else {
		stat.DecreaseRoleAdmin()
	}
	age := user.GetAge()
	switch {
	case age <= 18:
		stat.DecreaseAge0_18()
	case age <= 30 && age > 18:
		stat.DecreaseAge19_30()
	case age <= 50 && age > 30:
		stat.DecreaseAge31_50()
	case age > 50:
		stat.DecreaseAge51_100()
	}
	err = ctrl.UpdateStatistic(stat)
	if err != nil {
		s.Log.Error(err)
	}
}
