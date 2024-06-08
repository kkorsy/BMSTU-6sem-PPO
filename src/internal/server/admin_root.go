package server

import (
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func (s *srv) AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.session.Get(r, "sname")
		if err != nil {
			return
		}
		_, success := session.Values["admin"]
		if !success {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *srv) HandleAdminCabinet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id != "0" {
			switch id {
			case "1":
				{
					s.adminCabinetTemplate(w, "Cериал успешно добавлен")
					return
				}
			case "2":
				{
					s.adminCabinetTemplate(w, "Сериал успешо обновлен")
					return
				}
			case "3":
				{
					s.adminCabinetTemplate(w, "Сериал успешно удален")
					return
				}
			case "4":
				{
					s.adminCabinetTemplate(w, "Режиссер успешно добавлен")
					return
				}
			case "5":
				{
					s.adminCabinetTemplate(w, "Режиссер успешно обновлен")
					return
				}
			case "6":
				{
					s.adminCabinetTemplate(w, "Режиссер успешно удален")
					return
				}
			case "7":
				{
					s.adminCabinetTemplate(w, "Актер успешно добавлен")
					return
				}
			case "8":
				{
					s.adminCabinetTemplate(w, "Актер успешно обновлен")
					return
				}
			case "9":
				{
					s.adminCabinetTemplate(w, "Актер успешно удален")
					return
				}
			case "10":
				{
					s.adminCabinetTemplate(w, "Права администратора успешно назначены")
					return
				}
			case "11":
				{
					s.adminCabinetTemplate(w, "Пользователь успешно удален")
					return
				}
			}
		}
		s.adminCabinetTemplate(w, "")
	}
}

func (s *srv) adminCabinetTemplate(w http.ResponseWriter, msg string) {
	type Success struct {
		Msg string
	}
	tmpl, _ := template.ParseFiles("templates/admin/cabinet.html")
	scc := &Success{Msg: msg}
	tmpl.Execute(w, scc)
}

func (s *srv) HandleAddSerial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptAddSerial(w, r)
			return
		}
		s.addSerialTemplate(w, "")
	}
}

func (s *srv) addSerialTemplate(w http.ResponseWriter, err string) {
	type addSerialErr struct {
		Err       string
		Producers []*models.Producers
	}
	tmpl, _ := template.ParseFiles("templates/admin/addSerial.html")
	cerr := &addSerialErr{Err: err}
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producers, _ := ctrl.GetProducers()
	cerr.Producers = producers
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptAddSerial(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serial := &models.Serial{
		S_name:        r.FormValue("name"),
		S_description: r.FormValue("description"),
		S_genre:       r.FormValue("genre"),
		S_state:       r.FormValue("state"),
		S_img:         r.FormValue("img"),
	}
	idProducer, err := strconv.Atoi(r.FormValue("idProducer"))
	if err != nil {
		s.addSerialTemplate(w, "Продюссер не выбран")
		return
	}
	serial.SetIdProducer(idProducer)
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		s.addSerialTemplate(w, "Год выхода должен быть числом")
		return
	}
	serial.SetYear(year)
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 32)
	if err != nil {
		s.addSerialTemplate(w, "Рейтинг должен быть числом")
		return
	}
	serial.SetRating(float32(rating))
	serial.SetSeasons(0)

	err = ctrl.CreateSerial(serial)
	if err != nil {
		s.addSerialTemplate(w, "Ошибка создания сериала")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/1", http.StatusSeeOther)
}

func (s *srv) HandleUpdateSerial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("id") != "" {
			s.AcceptUpdateSerial(w, r)
			return
		} else if r.Method == http.MethodPost && r.FormValue("serial") != "" {
			s.ChoosenSerial(w, r)
			return
		}
		s.updateSerialTemplate(w, "", nil)
	}
}

func (s *srv) updateSerialTemplate(w http.ResponseWriter, err string, serial *models.Serial) {
	type updateSerialErr struct {
		S         *models.Serial
		Serials   []*models.Serial
		Producers []*models.Producers
		Err       string
	}
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, _ := ctrl.GetSerials()
	ctrlProducer := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producers, _ := ctrlProducer.GetProducers()
	tmpl, _ := template.ParseFiles("templates/admin/updateSerial.html")
	cerr := &updateSerialErr{Err: err, Serials: serials, S: serial, Producers: producers}
	tmpl.Execute(w, cerr)
}

func (s *srv) ChoosenSerial(w http.ResponseWriter, r *http.Request) {
	s_id, _ := strconv.Atoi(r.FormValue("serial"))
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serial, _ := ctrl.GetSerialById(s_id)
	s.updateSerialTemplate(w, "", serial)
}

func (s *srv) AcceptUpdateSerial(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	s_id, _ := strconv.Atoi(r.FormValue("id"))
	serial_prev, _ := ctrl.GetSerialById(s_id)
	serial := &models.Serial{
		S_id:          s_id,
		S_name:        r.FormValue("name"),
		S_description: r.FormValue("description"),
		S_genre:       r.FormValue("genre"),
		S_state:       r.FormValue("state"),
		S_img:         r.FormValue("img"),
	}
	idProducer, err := strconv.Atoi(r.FormValue("idProducer"))
	if err != nil {
		s.updateSerialTemplate(w, "id продюсера должен быть числом", serial_prev)
		return
	}
	serial.SetIdProducer(idProducer)
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		s.updateSerialTemplate(w, "Год выхода должен быть числом", serial_prev)
		return
	}
	serial.SetYear(year)
	rating, err := strconv.ParseFloat(r.FormValue("rating"), 32)
	if err != nil {
		s.updateSerialTemplate(w, "Рейтинг должен быть числом", serial_prev)
		return
	}
	serial.SetRating(float32(rating))
	serial.SetSeasons(serial_prev.S_seasons)
	serial.S_duration = serial_prev.S_duration

	err = ctrl.UpdateSerial(serial)
	if err != nil {
		s.updateSerialTemplate(w, "Ошибка обновления сериала", serial_prev)
		return
	}
	http.Redirect(w, r, "/admin/cabinet/2", http.StatusSeeOther)
}

func (s *srv) HandleDeleteSerial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptDeleteSerial(w, r)
			return
		}
		s.deleteSerialTemplate(w, "")
	}
}

func (s *srv) deleteSerialTemplate(w http.ResponseWriter, err string) {
	type deleteSerialErr struct {
		Err     string
		Serials []*models.Serial
	}
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, _ := ctrl.GetSerials()
	tmpl, _ := template.ParseFiles("templates/admin/deleteSerial.html")
	cerr := &deleteSerialErr{Err: err, Serials: serials}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptDeleteSerial(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	s_id, err := strconv.Atoi(r.FormValue("serial"))
	if err != nil {
		s.deleteSerialTemplate(w, "Сериал не выбран")
		return
	}
	err = ctrl.DeleteSerial(s_id)
	if err != nil {
		s.deleteSerialTemplate(w, "Ошибка удаления сериала")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/3", http.StatusSeeOther)
}

func (s *srv) HandleAddProducer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptAddProducer(w, r)
			return
		}
		s.addProducerTemplate(w, "")
	}
}

func (s *srv) addProducerTemplate(w http.ResponseWriter, err string) {
	type addProducerErr struct {
		Err string
	}
	tmpl, _ := template.ParseFiles("templates/admin/addProducer.html")
	cerr := &addProducerErr{Err: err}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptAddProducer(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producer := &models.Producers{
		P_name:    r.FormValue("name"),
		P_surname: r.FormValue("surname"),
	}
	if !producer.Validate() {
		s.addProducerTemplate(w, "Имя и фамилия не могут быть пустыми")
		return
	}
	err := ctrl.CreateProducer(producer)
	if err != nil {
		s.addProducerTemplate(w, "Ошибка создания продюссера")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/4", http.StatusSeeOther)
}

func (s *srv) HandleUpdateProducer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("id") != "" {
			s.AcceptUpdateProducer(w, r)
			return
		} else if r.Method == http.MethodPost && r.FormValue("producer") != "" {
			s.ChoosenProducer(w, r)
			return
		}
		s.updateProducerTemplate(w, "", nil)
	}
}

func (s *srv) updateProducerTemplate(w http.ResponseWriter, err string, producer *models.Producers) {
	type updateProducerErr struct {
		P         *models.Producers
		Producers []*models.Producers
		Err       string
	}
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producers, _ := ctrl.GetProducers()
	tmpl, _ := template.ParseFiles("templates/admin/updateProducer.html")
	cerr := &updateProducerErr{Err: err, Producers: producers, P: producer}
	tmpl.Execute(w, cerr)
}

func (s *srv) ChoosenProducer(w http.ResponseWriter, r *http.Request) {
	p_id, _ := strconv.Atoi(r.FormValue("producer"))
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producer, _ := ctrl.GetProducerById(p_id)
	s.Log.Println(producer)
	s.updateProducerTemplate(w, "", producer)
}

func (s *srv) AcceptUpdateProducer(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	p_id, _ := strconv.Atoi(r.FormValue("id"))
	prev_producer, _ := ctrl.GetProducerById(p_id)
	producer := &models.Producers{
		P_id:      p_id,
		P_name:    r.FormValue("name"),
		P_surname: r.FormValue("surname"),
	}
	if !producer.Validate() {
		s.updateProducerTemplate(w, "Имя и фамилия не могут быть пустыми", prev_producer)
		return
	}
	err := ctrl.UpdateProducer(producer)
	if err != nil {
		s.updateProducerTemplate(w, "Ошибка обновления продюссера", prev_producer)
		return
	}
	http.Redirect(w, r, "/admin/cabinet/5", http.StatusSeeOther)
}

func (s *srv) HandleDeleteProducer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptDeleteProducer(w, r)
			return
		}
		s.deleteProducerTemplate(w, "")
	}
}

func (s *srv) deleteProducerTemplate(w http.ResponseWriter, err string) {
	type deleteProducerErr struct {
		Err       string
		Producers []*models.Producers
	}
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producers, _ := ctrl.GetProducers()
	tmpl, _ := template.ParseFiles("templates/admin/deleteProducer.html")
	cerr := &deleteProducerErr{Err: err, Producers: producers}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptDeleteProducer(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	p_id, err := strconv.Atoi(r.FormValue("producer"))
	if err != nil {
		s.deleteProducerTemplate(w, "Режиссер не выбран")
		return
	}
	err = ctrl.DeleteProducer(p_id)
	if err != nil {
		s.deleteProducerTemplate(w, "Ошибка удаления продюссера")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/6", http.StatusSeeOther)
}

func (s *srv) HandleAddActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptAddActor(w, r)
			return
		}
		s.addActorTemplate(w, "")
	}
}

func (s *srv) addActorTemplate(w http.ResponseWriter, err string) {
	type addActorErr struct {
		Err     string
		Serials []*models.Serial
	}
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, _ := ctrl.GetSerials()
	tmpl, _ := template.ParseFiles("templates/admin/addActor.html")
	cerr := &addActorErr{Err: err, Serials: serials}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptAddActor(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	ctrlSA := controllers.NewSerialsActorsCtrl(repositories.NewSerialsActorsRepo(s.DB, s.Log))
	actor := &models.Actors{
		A_name:    r.FormValue("name"),
		A_surname: r.FormValue("surname"),
		A_gender:  r.FormValue("gender"),
		A_bdate:   r.FormValue("bdate"),
	}
	if !actor.Validate() {
		s.addActorTemplate(w, "Данные заполнены некорректно")
		return
	}
	if ctrl.CheckActor(actor) {
		s.addActorTemplate(w, "Актер уже существует")
		return
	}
	err := ctrl.CreateActor(actor)
	if err != nil {
		s.addActorTemplate(w, "Ошибка создания актера")
		return
	}

	_s_id := r.FormValue("serial")
	s_id, err := strconv.Atoi(_s_id)
	if err != nil {
		s.addActorTemplate(w, "Сериал не выбран")
		return
	}
	sa := &models.SerialsActors{
		Sa_idSerial: s_id,
		Sa_idActor:  actor.GetId(),
	}
	err = ctrlSA.CreateSerialsActors(sa)
	if err != nil {
		s.addActorTemplate(w, "Ошибка добавления актера в сериал")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/7", http.StatusSeeOther)
}

func (s *srv) HandleUpdateActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("id") != "" {
			s.AcceptUpdateActor(w, r)
			return
		} else if r.Method == http.MethodPost && r.FormValue("actor") != "" {
			s.ChoosenActor(w, r)
			return
		}
		s.updateActorTemplate(w, "", nil)
	}
}

func (s *srv) updateActorTemplate(w http.ResponseWriter, err string, actor *models.Actors) {
	type updateActorErr struct {
		A       *models.Actors
		Actors  []*models.Actors
		Serials []*models.Serial
		Err     string
	}
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	actors, _ := ctrl.GetActors()
	ctrlSerial := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, _ := ctrlSerial.GetSerials()
	tmpl, _ := template.ParseFiles("templates/admin/updateActor.html")
	cerr := &updateActorErr{Err: err, Actors: actors, A: actor, Serials: serials}
	tmpl.Execute(w, cerr)
}

func (s *srv) ChoosenActor(w http.ResponseWriter, r *http.Request) {
	a_id, _ := strconv.Atoi(r.FormValue("actor"))
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	actor, _ := ctrl.GetActorById(a_id)
	s.updateActorTemplate(w, "", actor)
}

func (s *srv) AcceptUpdateActor(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	a_id, _ := strconv.Atoi(r.FormValue("id"))
	prev_actor, _ := ctrl.GetActorById(a_id)
	actor := &models.Actors{
		A_id:      a_id,
		A_name:    r.FormValue("name"),
		A_surname: r.FormValue("surname"),
		A_gender:  r.FormValue("gender"),
		A_bdate:   r.FormValue("bdate"),
	}
	if !actor.Validate() {
		s.updateActorTemplate(w, "Данные заполнены некорректно", prev_actor)
		return
	}
	err := ctrl.UpdateActor(actor)
	if err != nil {
		s.updateActorTemplate(w, "Ошибка обновления актера", prev_actor)
		return
	}
	http.Redirect(w, r, "/admin/cabinet/8", http.StatusSeeOther)
}

func (s *srv) HandleDeleteActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptDeleteActor(w, r)
			return
		}
		s.deleteActorTemplate(w, "")
	}
}

func (s *srv) deleteActorTemplate(w http.ResponseWriter, err string) {
	type deleteActorErr struct {
		Err    string
		Actors []*models.Actors
	}
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	actors, _ := ctrl.GetActors()
	tmpl, _ := template.ParseFiles("templates/admin/deleteActor.html")
	cerr := &deleteActorErr{Err: err, Actors: actors}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptDeleteActor(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	a_id, err := strconv.Atoi(r.FormValue("actor"))
	if err != nil {
		s.deleteActorTemplate(w, "Актер не выбран")
		return
	}
	ctrlSA := controllers.NewSerialsActorsCtrl(repositories.NewSerialsActorsRepo(s.DB, s.Log))
	serials, _ := ctrlSA.GetSerialsByActorId(a_id)
	for _, serial := range serials {
		err = ctrlSA.DeleteSerialsActors(serial.GetId())
		if err != nil {
			s.deleteActorTemplate(w, "Ошибка удаления актера из сериала")
			return
		}
	}

	err = ctrl.DeleteActor(a_id)
	if err != nil {
		s.deleteActorTemplate(w, "Ошибка удаления актера")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/9", http.StatusSeeOther)
}

func (s *srv) HandleShowUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
		users, _ := ctrl.GetUsers()
		tmpl, _ := template.ParseFiles("templates/admin/showUsers.html")
		tmpl.Execute(w, users)
	}
}

func (s *srv) HandleGrantAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptGrantAdmin(w, r)
			return
		}
		s.grantAdminTemplate(w, "")
	}
}

func (s *srv) grantAdminTemplate(w http.ResponseWriter, err string) {
	type grantAdminErr struct {
		Err   string
		Users []*models.Users
	}
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	users, _ := ctrl.GetUsers()
	tmpl, _ := template.ParseFiles("templates/admin/grantAdmin.html")
	cerr := &grantAdminErr{Err: err, Users: users}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptGrantAdmin(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	u_id, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		s.grantAdminTemplate(w, "Пользователь не выбран")
		return
	}
	user, _ := ctrl.GetUserById(u_id)
	s.UpdateStatisticDown(user)

	err = ctrl.GrantAdmin(u_id)
	if err != nil {
		s.grantAdminTemplate(w, "Ошибка назначения администратора")
		return
	}
	// update statistics
	ctrlStat := controllers.NewStatisticCtrl(repositories.NewStatisticRepo(s.DB, s.Log))
	statistic, _ := ctrlStat.GetStatistic()
	statistic.IncreaseRoleAdmin()
	err = ctrlStat.UpdateStatistic(statistic)
	if err != nil {
		s.grantAdminTemplate(w, "Ошибка обновления статистики")
		return
	}

	http.Redirect(w, r, "/admin/cabinet/10", http.StatusSeeOther)
}

func (s *srv) HandleDeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptDeleteUser(w, r)
			return
		}
		s.deleteUserTemplate(w, "")
	}
}

func (s *srv) deleteUserTemplate(w http.ResponseWriter, err string) {
	type deleteUserErr struct {
		Err   string
		Users []*models.Users
	}
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	users, _ := ctrl.GetUsers()
	tmpl, _ := template.ParseFiles("templates/admin/deleteUser.html")
	cerr := &deleteUserErr{Err: err, Users: users}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	u_id, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		s.deleteUserTemplate(w, "Пользователь не выбран")
		return
	}
	user, _ := ctrl.GetUserById(u_id)
	ctrlSF := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(s.DB, s.Log))
	favourites, _ := ctrlSF.GetSerialsByFavouriteId(user.GetIdFavourites())
	for _, favourite := range favourites {
		err = ctrlSF.DeleteSerialsFavourites(favourite.GetId())
		if err != nil {
			s.deleteUserTemplate(w, "Ошибка удаления сериала из избранного")
			return
		}
	}

	ctrlC := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	comments, _ := ctrlC.GetCommentsByUserId(user.GetId())
	for _, comment := range comments {
		err = ctrlC.DeleteComment(comment.GetId())
		if err != nil {
			s.deleteUserTemplate(w, "Ошибка удаления комментариев")
			return
		}
	}

	ctrlSu := controllers.NewSerialsUsersCtrl(repositories.NewSerialsUsersRepo(s.DB, s.Log))
	err = ctrlSu.DeleteSerialsByUserId(user.GetId())
	if err != nil {
		s.deleteUserTemplate(w, "Ошибка удаления сериалов пользователя")
		return
	}

	err = ctrl.DeleteUser(u_id)
	if err != nil {
		s.deleteUserTemplate(w, "Ошибка удаления пользователя")
		return
	}

	// update statistics
	s.UpdateStatisticDown(user)

	http.Redirect(w, r, "/admin/cabinet/11", http.StatusSeeOther)
}

func (s *srv) HandleAddSerialActor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.AcceptAddSerialActor(w, r)
			return
		}
		s.addSerialActorTemplate(w, "")
	}
}

func (s *srv) addSerialActorTemplate(w http.ResponseWriter, err string) {
	type addSerialActorErr struct {
		Err    string
		Serial []*models.Serial
		Actors []*models.Actors
	}
	ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serials, _ := ctrl.GetSerials()
	ctrlA := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	actors, _ := ctrlA.GetActors()
	tmpl, _ := template.ParseFiles("templates/admin/addSerialActor.html")
	cerr := &addSerialActorErr{Err: err, Serial: serials, Actors: actors}
	tmpl.Execute(w, cerr)
}

func (s *srv) AcceptAddSerialActor(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewSerialsActorsCtrl(repositories.NewSerialsActorsRepo(s.DB, s.Log))
	s_id, err := strconv.Atoi(r.FormValue("serial"))
	if err != nil {
		s.addSerialActorTemplate(w, "Сериал не выбран")
		return
	}
	a_id, err := strconv.Atoi(r.FormValue("actor"))
	if err != nil {
		s.addSerialActorTemplate(w, "Актер не выбран")
		return
	}
	sa := &models.SerialsActors{
		Sa_idSerial: s_id,
		Sa_idActor:  a_id,
	}
	err = ctrl.CreateSerialsActors(sa)
	if err != nil {
		s.addSerialActorTemplate(w, "Ошибка добавления актера в сериал")
		return
	}
	http.Redirect(w, r, "/admin/cabinet/7", http.StatusSeeOther)
}

func (s *srv) HandleShowStatistics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.NewStatisticCtrl(repositories.NewStatisticRepo(s.DB, s.Log))
		statistic, _ := ctrl.GetStatistic()
		tmpl, _ := template.ParseFiles("templates/admin/showStatistics.html")
		tmpl.Execute(w, statistic)
	}
}
