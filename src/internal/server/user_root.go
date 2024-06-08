package server

import (
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func (s *srv) UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		_, success := session.Values["user"]
		if !success {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *srv) HandleUserCabinet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id != "0" {
			switch id {
			case "1":
				{
					s.userCabinetTemplate(w, r, "Комментарий добавлен")
					return
				}
			case "2":
				{
					s.userCabinetTemplate(w, r, "Комментарий обновлен")
					return
				}
			case "3":
				{
					s.userCabinetTemplate(w, r, "Комментарий удален")
					return
				}
			case "4":
				{
					s.userCabinetTemplate(w, r, "Профиль изменен")
					return
				}
			}
		}
		s.userCabinetTemplate(w, r, "")
	}
}

func (s *srv) userCabinetTemplate(w http.ResponseWriter, r *http.Request, msg string) {
	type Data struct {
		Msg  string
		User *models.Users
	}
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	id := session.Values["user"].(int)
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user, err := ctrl.GetUserById(id)
	if err != nil {
		return
	}

	d := &Data{Msg: msg,
		User: user,
	}

	tmpl, _ := template.ParseFiles("templates/user/cabinet.html")
	tmpl.Execute(w, d)
}

func (s *srv) HandleCompareSerials() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Data struct {
			Serials []*models.Serial
			Compare []*models.Serial
		}
		d := &Data{}

		if r.Method == http.MethodPost {
			r.ParseForm()
			ids := r.Form["serial"]
			ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
			for _, id := range ids {
				id, _ := strconv.Atoi(id)
				serial, err := ctrl.GetSerialById(id)
				if err != nil {
					return
				}
				d.Compare = append(d.Compare, serial)
			}
		} else {
			ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
			serials, err := ctrl.GetSerials()
			if err != nil {
				return
			}
			d.Serials = serials
		}
		tmpl, _ := template.ParseFiles("templates/user/compare.html")
		tmpl.Execute(w, d)
	}
}

func (s *srv) HandleFavourites() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Data struct {
			Favourites []*models.Serial
		}
		d := &Data{}

		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		id := session.Values["user"].(int)
		ctrl := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(s.DB, s.Log))
		serials, err := ctrl.GetSerialsByFavouriteId(id)
		if err != nil {
			return
		}
		for _, serial := range serials {
			ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
			s, err := ctrl.GetSerialById(serial.GetIdSerial())
			if err != nil {
				return
			}
			d.Favourites = append(d.Favourites, s)
		}

		tmpl, _ := template.ParseFiles("templates/user/favourites.html")
		tmpl.Execute(w, d)
	}
}

func (s *srv) HandleDeleteFavourite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tid := r.FormValue("serial")
		id, _ := strconv.Atoi(tid)

		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		iduser := session.Values["user"].(int)
		ctrlUser := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
		user, err := ctrlUser.GetUserById(iduser)
		if err != nil {
			return
		}

		ctrl := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(s.DB, s.Log))
		err = ctrl.DeleteSerialById(user.GetIdFavourites(), id)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/user/favourites", http.StatusSeeOther)
	}
}

func (s *srv) HandleAddComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptAddComment(w, r)
			return
		}
		s.addCommentTemplate(w, "")
	}
}

func (s *srv) addCommentTemplate(w http.ResponseWriter, msg string) {
	type Data struct {
		Err     string
		Serials []*models.Serial
	}
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, err := ctrl.GetSerials()
	if err != nil {
		return
	}
	d := &Data{Err: msg, Serials: serials}

	tmpl, _ := template.ParseFiles("templates/user/addComment.html")
	tmpl.Execute(w, d)
}

func (s *srv) AcceptAddComment(w http.ResponseWriter, r *http.Request) {
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	id := session.Values["user"].(int)

	idserial, err := strconv.Atoi(r.FormValue("idserial"))
	if err != nil {
		s.addCommentTemplate(w, "Сериал не выбран")
		return
	}
	ctrlComment := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	if ctrlComment.CheckComment(id, idserial) {
		s.addCommentTemplate(w, "Вы уже оставляли комментарий к этому сериалу")
		return

	}

	comment := r.FormValue("comment")
	if comment == "" {
		s.addCommentTemplate(w, "Текст комментария не может быть пустым")
		return
	}

	err = ctrlComment.CreateComment(&models.Comments{
		C_idUser:   id,
		C_idSerial: idserial,
		C_text:     comment,
		C_date:     time.Now().Format("2006-01-02"),
	})
	if err != nil {
		return
	}
	http.Redirect(w, r, "/user/cabinet/1", http.StatusSeeOther)
}

func (s *srv) HandleUpdateComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("id") != "" {
			s.AcceptUpdateComment(w, r)
			return
		} else if r.Method == http.MethodPost && r.FormValue("idcomment") != "" {
			s.ChoosenComment(w, r)
			return
		}
		s.updateCommentTemplate(w, r, "", nil)
	}
}

func (s *srv) updateCommentTemplate(w http.ResponseWriter, r *http.Request, msg string, comment_prev *models.Comments) {
	type SerialsComments struct {
		Serial  *models.Serial
		Comment *models.Comments
	}
	type Data struct {
		Err string
		SC  []*SerialsComments
		C   *models.Comments
	}
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	id := session.Values["user"].(int)
	ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	comments, err := ctrl.GetCommentsByUserId(id)
	if err != nil {
		return
	}
	serialscomments := []*SerialsComments{}
	ctrlS := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	for _, comment := range comments {
		serial, err := ctrlS.GetSerialById(comment.GetIdSerial())
		if err != nil {
			return
		}
		sc := &SerialsComments{Serial: serial, Comment: comment}
		serialscomments = append(serialscomments, sc)
	}
	d := &Data{Err: msg, SC: serialscomments, C: comment_prev}

	tmpl, _ := template.ParseFiles("templates/user/updateComment.html")
	tmpl.Execute(w, d)
}

func (s *srv) ChoosenComment(w http.ResponseWriter, r *http.Request) {
	c_id, _ := strconv.Atoi(r.FormValue("idcomment"))
	ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	comment, _ := ctrl.GetCommentById(c_id)
	s.updateCommentTemplate(w, r, "", comment)

}

func (s *srv) AcceptUpdateComment(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	c_id, _ := strconv.Atoi(r.FormValue("id"))
	comment_prev, _ := ctrl.GetCommentById(c_id)

	comment := r.FormValue("comment")
	if comment == "" {
		s.updateCommentTemplate(w, r, "Текст комментария не может быть пустым", comment_prev)
		return
	}

	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	iduser := session.Values["user"].(int)

	err = ctrl.UpdateComment(&models.Comments{
		C_id:       c_id,
		C_idUser:   iduser,
		C_idSerial: comment_prev.GetIdSerial(),
		C_text:     comment,
		C_date:     time.Now().Format("2006-01-02"),
	})
	if err != nil {
		return
	}
	http.Redirect(w, r, "/user/cabinet/2", http.StatusSeeOther)
}

func (s *srv) HandleDeleteComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptDeleteComment(w, r)
			return
		}
		s.deleteCommentTemplate(w, r, "")
	}
}

func (s *srv) deleteCommentTemplate(w http.ResponseWriter, r *http.Request, msg string) {
	type deleteCommentErr struct {
		Err     string
		Serials []*models.Serial
	}

	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	iduser := session.Values["user"].(int)
	ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	comments, err := ctrl.GetCommentsByUserId(iduser)
	if err != nil {
		return
	}
	serials := []*models.Serial{}
	ctrlS := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	for _, comment := range comments {
		serial, err := ctrlS.GetSerialById(comment.GetIdSerial())
		if err != nil {
			return
		}
		serials = append(serials, serial)
	}
	d := &deleteCommentErr{Err: msg, Serials: serials}

	tmpl, _ := template.ParseFiles("templates/user/deleteComment.html")
	tmpl.Execute(w, d)
}

func (s *srv) AcceptDeleteComment(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	_s_id := r.FormValue("serial")
	if _s_id == "" {
		s.deleteCommentTemplate(w, r, "Сериал не выбран")
		return
	}
	s_id, _ := strconv.Atoi(_s_id)
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	iduser := session.Values["user"].(int)
	comment, _ := ctrl.GetCommentsBySerialIdUserId(s_id, iduser)
	err = ctrl.DeleteComment(comment.GetId())
	if err != nil {
		s.deleteCommentTemplate(w, r, "Ошибка удаления комментария")
		return
	}
	http.Redirect(w, r, "/user/cabinet/3", http.StatusSeeOther)
}

func (s *srv) HandleUpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptUpdateProfile(w, r)
			return
		}
		s.updateProfileTemplate(w, r, "")
	}
}

func (s *srv) updateProfileTemplate(w http.ResponseWriter, r *http.Request, msg string) {
	type Data struct {
		Err  string
		User *models.Users
	}
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	id := session.Values["user"].(int)
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user, err := ctrl.GetUserById(id)
	if err != nil {
		return
	}
	d := &Data{Err: msg, User: user}

	tmpl, _ := template.ParseFiles("templates/user/updateProfile.html")
	tmpl.Execute(w, d)
}

func (s *srv) AcceptUpdateProfile(w http.ResponseWriter, r *http.Request) {
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	id := session.Values["user"].(int)
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user_prev, err := ctrl.GetUserById(id)
	if err != nil {
		return
	}

	name := r.FormValue("name")
	if name == "" {
		s.updateProfileTemplate(w, r, "Имя не может быть пустым")
		return
	}
	surname := r.FormValue("surname")
	if surname == "" {
		s.updateProfileTemplate(w, r, "Фамилия не может быть пустой")
		return
	}
	login := r.FormValue("login")
	if login == "" {
		s.updateProfileTemplate(w, r, "Логин не может быть пустым")
		return
	}
	password := r.FormValue("password")
	var pass []byte
	if password == "" {
		pass = []byte(user_prev.GetPassword())
	} else {
		pass, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			s.creationTemplate(w, "bcrypt internal error")
			return
		}
	}
	bdate := r.FormValue("bdate")
	if bdate == "" {
		s.updateProfileTemplate(w, r, "Дата рождения не может быть пустой")
		return
	}
	gender := r.FormValue("gender")
	if gender == "" {
		s.updateProfileTemplate(w, r, "Пол не может быть пустым")
		return
	}

	user := &models.Users{
		U_id:           id,
		U_name:         name,
		U_surname:      surname,
		U_login:        login,
		U_password:     string(pass),
		U_bdate:        bdate,
		U_gender:       gender,
		U_role:         "user",
		U_idFavourites: user_prev.GetIdFavourites(),
	}
	err = ctrl.UpdateUser(user)
	if err != nil {
		s.Log.Error(err)
		return
	}
	s.UpdateStatisticDown(user_prev)
	s.UpdateStatisticUp(user)
	http.Redirect(w, r, "/user/cabinet/4", http.StatusSeeOther)
}

func (s *srv) HandleHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type History struct {
			Serial *models.Serial
			Date   string
		}
		type Data struct {
			History []*History
		}
		d := &Data{}

		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		id := session.Values["user"].(int)
		ctrl := controllers.NewSerialsUsersCtrl(repositories.NewSerialsUsersRepo(s.DB, s.Log))
		history, err := ctrl.GetSerialsByUserId(id)
		if err != nil {
			return
		}
		ctrlS := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
		for _, h := range history {
			serial, err := ctrlS.GetSerialById(h.GetIdSerial())
			if err != nil {
				return
			}
			dat := &History{Serial: serial, Date: h.GetLastSeen()}
			d.History = append(d.History, dat)
		}

		tmpl, _ := template.ParseFiles("templates/user/history.html")
		tmpl.Execute(w, d)
	}
}

func (s *srv) HandleClearHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		id := session.Values["user"].(int)
		ctrl := controllers.NewSerialsUsersCtrl(repositories.NewSerialsUsersRepo(s.DB, s.Log))
		err = ctrl.DeleteSerialsByUserId(id)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/user/history", http.StatusSeeOther)
	}
}
