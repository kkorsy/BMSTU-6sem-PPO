package server

import (
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"errors"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func (s *srv) AddToFavorite(w http.ResponseWriter, r *http.Request) error {
	session, err := s.session.Get(r, "sname")
	if err != nil {
		return err
	}
	id_str := session.Values["admin"]
	if id_str != nil {
		return errors.New("Администратор не может добавлять сериалы в избранное")
	}
	id_str = session.Values["user"]
	if id_str == nil {
		return errors.New("Пользователь не авторизирован")
	}
	id := id_str.(int)
	ctrlUser := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user, err := ctrlUser.GetUserById(id)
	if err != nil {
		return err
	}

	idSerial, _ := strconv.Atoi(r.FormValue("serial_id"))
	ctrlF := controllers.NewFavouritesCtrl(repositories.NewFavouritesRepo(s.DB, s.Log))
	fav, err := ctrlF.GetFavouriteById(user.U_idFavourites)
	if err != nil {
		return err
	}
	serialFav := &models.SerialsFavourites{
		Sf_idSerial:    idSerial,
		Sf_idFavourite: fav.GetId(),
	}

	ctrlSF := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(s.DB, s.Log))
	if ctrlSF.CheckSerialInFavourite(serialFav) {
		return errors.New("Сериал уже добавлен в избранное")
	}
	err = ctrlSF.CreateSerialsFavourites(serialFav)
	if err != nil {
		return err
	}
	fav.SetCntSerials(fav.GetCntSerials() + 1)
	ctrlF.UpdateFavourite(fav)
	return nil
}

func (s *srv) HandleSerial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := s.AddToFavorite(w, r)
			if err != nil {
				s.serialTemplate(w, r, err.Error())
				return
			} else {
				s.serialTemplate(w, r, "Сериал успешно добавлен в избранное")
				return
			}
		}
		s.serialTemplate(w, r, "")
	}
}

func (s *srv) serialTemplate(w http.ResponseWriter, r *http.Request, msg string) {
	type seasonEpisodes struct {
		Season   *models.Seasons
		Episodes []*models.Episodes
	}

	type CommentsData struct {
		C_text string
		U_name string
	}

	type Data struct {
		Serial   *models.Serial
		Seasons  []*seasonEpisodes
		Err      string
		Comments []*CommentsData
		Actors   []*models.Actors
		Producer *models.Producers
	}

	d := &Data{Err: msg}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctrlSerials := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(s.DB, s.Log))
	serial, err := ctrlSerials.GetSerialById(id)
	if err != nil {
		return
	}
	d.Serial = serial

	ctrlSeasons := controllers.NewSeasonsCtrl(repositories.NewSeasonsRepo(s.DB, s.Log))
	seasons, err := ctrlSeasons.GetSeasonsBySerialId(id)
	if err != nil {
		return
	}

	ctrlEpisodes := controllers.NewEpisodesCtrl(repositories.NewEpisodesRepo(s.DB, s.Log))
	for _, season := range seasons {
		episodes, err := ctrlEpisodes.GetEpisodesBySeasonId(season.GetId())
		if err != nil {
			return
		}
		SesEp := &seasonEpisodes{
			Season:   season,
			Episodes: episodes,
		}
		d.Seasons = append(d.Seasons, SesEp)
	}

	ctrlComments := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(s.DB, s.Log))
	comments, err := ctrlComments.GetCommentsBySerialId(id)
	if err != nil {
		return
	}
	ctrlUser := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	for _, comment := range comments {
		user, err := ctrlUser.GetUserById(comment.GetIdUser())
		if err != nil {
			return
		}
		c := &CommentsData{
			C_text: comment.GetText(),
			U_name: user.GetName(),
		}
		d.Comments = append(d.Comments, c)
	}
	ctrlSa := controllers.NewSerialsActorsCtrl(repositories.NewSerialsActorsRepo(s.DB, s.Log))
	actors, err := ctrlSa.GetActorsBySerialId(id)
	if err != nil {
		return
	}
	ctrlA := controllers.NewActorsCtrl(repositories.NewActorsRepo(s.DB, s.Log))
	for _, actor := range actors {
		a, err := ctrlA.GetActorById(actor.GetIdActor())
		if err != nil {
			return
		}
		d.Actors = append(d.Actors, a)
	}
	ctrlP := controllers.NewProducersCtrl(repositories.NewProducersRepo(s.DB, s.Log))
	producer, err := ctrlP.GetProducerById(serial.GetIdProducer())
	if err != nil {
		return
	}
	d.Producer = producer

	session, err := s.session.Get(r, "sname")
	if err != nil {
		return
	}
	iduser := session.Values["user"]
	if iduser != nil {
		s.addToHistory(iduser.(int), id)
	}

	tmpl, _ := template.ParseFiles("templates/serial/serial.html")
	tmpl.Execute(w, d)
}

func (s *srv) addToHistory(iduser int, idserial int) {
	ctrlUser := controllers.NewUsersCtrl(repositories.NewUsersRepo(s.DB, s.Log), repositories.NewFavouritesRepo(s.DB, s.Log))
	user, err := ctrlUser.GetUserById(iduser)
	if err != nil {
		return
	}
	ctrlHistory := controllers.NewSerialsUsersCtrl(repositories.NewSerialsUsersRepo(s.DB, s.Log))
	su, err := ctrlHistory.GetSerialUserByIds(idserial, user.GetId())
	if err == nil {
		su.SetLastSeen(time.Now().Format("2006-01-02"))
		err = ctrlHistory.UpdateSerialsUsers(su)
		if err != nil {
			return
		}
	} else {
		history := &models.SerialsUsers{
			Su_idSerial: idserial,
			Su_idUser:   user.GetId(),
			Su_lastSeen: time.Now().Format("2006-01-02"),
		}
		err = ctrlHistory.CreateSerialsUsers(history)
		if err != nil {
			return
		}
	}
}
