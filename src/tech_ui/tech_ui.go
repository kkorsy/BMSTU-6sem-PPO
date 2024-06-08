package tech_ui

import (
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Run(db *sqlx.DB, log *logrus.Logger) {
	var curUser *models.Users = nil // replace to cookie in ui?
	for {
		req := Menu()
		switch req {
		// выйти из программы
		case 0:
			log.Info("Exit from program")
			return
		// зарегистрироваться
		case 1:
			{
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				user := &models.Users{}
				fmt.Println("Введите логин:")
				fmt.Scan(&user.U_login)
				fmt.Println("Введите пароль:")
				fmt.Scan(&user.U_password)
				fmt.Println("Введите имя:")
				fmt.Scan(&user.U_name)
				fmt.Println("Введите фамилию:")
				fmt.Scan(&user.U_surname)
				fmt.Println("Введите пол:")
				fmt.Scan(&user.U_gender)
				fmt.Println("Введите дату рождения:")
				fmt.Scan(&user.U_bdate)
				user.SetRole("user")

				err := ctrl.CreateUser(user)
				if err != nil {
					log.Error(err)
					break
				}
				curUser = user
				log.Info("User " + user.U_login + "registered")
			}
		// войти в аккаунт
		case 2:
			{
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				var login, password string
				fmt.Println("Введите логин:")
				fmt.Scan(&login)
				fmt.Println("Введите пароль:")
				fmt.Scan(&password)
				user, err := ctrl.AuthUser(login, password)
				if err != nil {
					log.Error(err)
				} else {
					curUser = user
					log.Info("User logged in")
				}
			}
		// выйти из аккаунта
		case 3:
			{
				curUser = nil
				log.Info("User logged out")
			}
		// посмотреть список сериалов
		case 4:
			{
				ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(db, log))
				serials, err := ctrl.GetSerials()
				if err != nil {
					log.Error(err)
					break
				}
				for _, s := range serials {
					fmt.Println(*s)
				}
				log.Info("User checked serials list")
			}
		// найти сериал
		case 5:
			{
				ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(db, log))
				fmt.Println("Введите название сериала:")
				var title string
				fmt.Scan(&title)
				serials, err := ctrl.SerialsService.GetSerialsByTitle(title)
				if err != nil {
					log.Error(err)
					break
				}
				for _, s := range serials {
					fmt.Println(*s)
				}
				log.Info("User searched serial with title: " + title)
			}
		// добавить отзыв
		case 6:
			{
				if curUser == nil {
					log.Warn("User tried to add comment without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to add comment")
					fmt.Println("Администратор не может оставлять отзывы")
					break
				}
				ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(db, log))
				comment := &models.Comments{}
				fmt.Println("Введите id сериала:")
				fmt.Scan(&comment.C_idSerial)
				fmt.Println("Введите текст отзыва:")
				fmt.Scan(&comment.C_text)
				comment.SetIdUser(curUser.GetId())
				comment.SetDate(time.Now().Format("2006-01-02"))
				fmt.Println(comment)
				err := ctrl.CreateComment(comment)
				if err != nil {
					log.Error(err)
					break
				}
				log.Info("User added comment")
			}
		// изменить отзыв
		case 7:
			{
				if curUser == nil {
					log.Warn("User tried to update comment without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to update comment")
					fmt.Println("Администратор не может изменять отзывы")
					break
				}
				ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(db, log))
				var id int
				fmt.Println("Введите id отзыва:")
				fmt.Scan(&id)
				comment, err := ctrl.GetCommentById(id)
				if err != nil {
					log.Error(err)
					break
				}
				fmt.Println("Введите новый текст отзыва:")
				fmt.Scan(&comment.C_text)
				err = ctrl.UpdateComment(comment)
				if err != nil {
					log.Error(err)
					break
				}
				log.Info("User updated comment")
			}
		// удалить отзыв
		case 8:
			{
				if curUser == nil {
					log.Warn("User tried to delete comment without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to delete comment")
					fmt.Println("Администратор не может удалять отзывы")
					break
				}
				ctrl := controllers.NewCommentsCtrl(repositories.NewCommentsRepo(db, log))
				comment := &models.Comments{}
				fmt.Println("Введите id отзыва:")
				fmt.Scan(&comment.C_id)
				comment.SetIdUser(curUser.GetId())
				err := ctrl.DeleteComment(comment.GetId())
				if err != nil {
					log.Error(err)
					break
				}
				log.Info("User deleted comment")
			}
		// сохранить сериал в избранное
		case 9:
			{
				if curUser == nil {
					log.Warn("User tried to add serial to favourites without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to add serial to favourites")
					fmt.Println("Администратор не может сохранять сериалы в избранное")
					break
				}
				ctrl := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(db, log))
				serialFav := &models.SerialsFavourites{}
				fmt.Println("Введите id сериала:")
				fmt.Scan(&serialFav.Sf_idSerial)
				serialFav.SetIdFavourite(curUser.GetIdFavourites())
				err := ctrl.CreateSerialsFavourites(serialFav)
				if err != nil {
					log.Error(err)
					break
				}

				ctrlFav := controllers.NewFavouritesCtrl(repositories.NewFavouritesRepo(db, log))
				fav, err := ctrlFav.GetFavouriteById(curUser.GetIdFavourites())
				if err != nil {
					log.Error(err)
				} else {
					fav.SetCntSerials(fav.GetCntSerials() + 1)
					err = ctrlFav.UpdateFavourite(fav)
					if err != nil {
						log.Error(err)
					} else {
						log.Info("User added serial to favourites")
					}
				}
			}
		// посмотреть список избранных сериалов
		case 10:
			{
				if curUser == nil {
					log.Warn("User tried to check favourites without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to check favourites")
					fmt.Println("Администратор не может просматривать список избранных сериалов")
					break
				}
				ctrl := controllers.NewSerialsFavouritesCtrl(repositories.NewSerialsFavouritesRepo(db, log))
				serialsFav, err := ctrl.GetSerialsByFavouriteId(curUser.GetIdFavourites())
				if err != nil {
					log.Error(err)
				} else {
					for _, s := range serialsFav {
						fmt.Println(*s)
					}
					log.Info("User checked favourites")
				}
			}
		// посмотреть профиль
		case 11:
			{
				if curUser == nil {
					log.Warn("User tried to check profile without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to check profile")
					fmt.Println("Администратор не может просматривать профиль")
					break
				}
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				user, err := ctrl.GetUserById(curUser.GetId())
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println(*user)
					log.Info("User checked profile")
				}
			}
		// редактировать профиль
		case 12:
			{
				if curUser == nil {
					log.Warn("User tried to update profile without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to update profile")
					fmt.Println("Администратор не может редактировать профиль")
					break
				}
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				fmt.Println("Поля, доступные для изменения:\n" +
					"1. Логин\n" +
					"2. Пароль\n" +
					"3. Имя\n" +
					"4. Фамилия\n" +
					"5. Пол\n" +
					"6. Дата рождения\n\n" +
					"Введите номер поля, которое хотите изменить:")
				var field int
				fmt.Scan(&field)
				switch field {
				case 1:
					{
						fmt.Println("Введите новый логин:")
						fmt.Scan(&curUser.U_login)
					}
				case 2:
					{
						fmt.Println("Введите новый пароль:")
						fmt.Scan(&curUser.U_password)
					}
				case 3:
					{
						fmt.Println("Введите новое имя:")
						fmt.Scan(&curUser.U_name)
					}
				case 4:
					{
						fmt.Println("Введите новую фамилию:")
						fmt.Scan(&curUser.U_surname)
					}
				case 5:
					{
						fmt.Println("Введите новый пол:")
						fmt.Scan(&curUser.U_gender)
					}
				case 6:
					{
						fmt.Println("Введите новую дату рождения:")
						fmt.Scan(&curUser.U_bdate)
					}
				}
				err := ctrl.UpdateUser(curUser)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User updated profile")
				}
			}
		// посмотреть историю просмотров
		case 13:
			{
				if curUser == nil {
					log.Warn("User tried to check history without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() == "admin" {
					log.Warn("Admin tried to check history")
					fmt.Println("Администратор не может просматривать историю просмотров")
					break
				}
				ctrl := controllers.NewSerialsUsersCtrl(repositories.NewSerialsUsersRepo(db, log))
				serialsUsers, err := ctrl.GetSerialsByUserId(curUser.GetId())
				if err != nil {
					log.Error(err)
				} else {
					for _, s := range serialsUsers {
						fmt.Println(*s)
					}
					log.Info("User checked history")
				}
			}
		// добавить сериал
		case 14:
			{
				if curUser == nil {
					log.Warn("User tried to add serial without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to add serial without admin rights")
					fmt.Println("Только администратор может добавлять сериалы")
					break
				}
				ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(db, log))
				serial := &models.Serial{}
				fmt.Println("Введите название сериала:")
				fmt.Scan(&serial.S_name)
				fmt.Println("Введите описание сериала:")
				fmt.Scan(&serial.S_description)
				fmt.Println("Введите жанр сериала:")
				fmt.Scan(&serial.S_genre)
				fmt.Println("Введите год выхода сериала:")
				fmt.Scan(&serial.S_year)
				fmt.Println("Введите рейтинг сериала:")
				fmt.Scan(&serial.S_rating)
				fmt.Println("Введите статус сериала:")
				fmt.Scan(&serial.S_state)
				fmt.Println("Введите id продюсера сериала:")
				fmt.Scan(&serial.S_idProducer)
				serial.SetSeasons(0)
				err := ctrl.CreateSerial(serial)
				if err != nil {
					log.Error(err)
				}
			}
		// изменить сериал
		case 15:
			{
				if curUser == nil {
					log.Warn("User tried to update serial without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to update serial without admin rights")
					fmt.Println("Только администратор может изменять сериалы")
					break
				}
				ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(db, log))
				var id int
				fmt.Println("Введите id сериала:")
				fmt.Scan(&id)
				serial, err := ctrl.GetSerialById(id)
				if err != nil {
					log.Error(err)
					break
				}
				fmt.Println("Поля, доступные для изменения:\n" +
					"1. Название\n" +
					"2. Описание\n" +
					"3. Жанр\n" +
					"4. Год выхода\n" +
					"5. Рейтинг\n" +
					"6. Статус\n" +
					"7. Продюсер\n\n" +
					"Введите номер поля, которое хотите изменить:")
				var field int
				fmt.Scan(&field)
				switch field {
				case 1:
					{
						fmt.Println("Введите новое название:")
						fmt.Scan(&serial.S_name)
					}
				case 2:
					{
						fmt.Println("Введите новое описание:")
						fmt.Scan(&serial.S_description)
					}
				case 3:
					{
						fmt.Println("Введите новый жанр:")
						fmt.Scan(&serial.S_genre)
					}
				case 4:
					{
						fmt.Println("Введите новый год выхода:")
						fmt.Scan(&serial.S_year)
					}
				case 5:
					{
						fmt.Println("Введите новый рейтинг:")
						fmt.Scan(&serial.S_rating)
					}
				case 6:
					{
						fmt.Println("Введите новый статус:")
						fmt.Scan(&serial.S_state)
					}
				case 7:
					{
						fmt.Println("Введите новый id продюсера:")
						fmt.Scan(&serial.S_idProducer)
					}
				}
				err = ctrl.UpdateSerial(serial)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User updated serial")
				}
			}
		// удалить сериал
		case 16:
			{
				if curUser == nil {
					log.Warn("User tried to delete serial without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to delete serial without admin rights")
					fmt.Println("Только администратор может удалять сериалы")
					break
				}
				ctrl := controllers.NewSerialsCtrl(repositories.NewSerialsRepo(db, log))
				serial := &models.Serial{}
				fmt.Println("Введите id сериала:")
				fmt.Scan(&serial.S_id)
				err := ctrl.DeleteSerial(serial.GetId())
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User deleted serial")
				}
			}
		// добавить актера
		case 17:
			{
				if curUser == nil {
					log.Warn("User tried to add actor without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to add actor without admin rights")
					fmt.Println("Только администратор может добавлять актеров")
					break
				}
				ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(db, log))
				actor := &models.Actors{}
				fmt.Println("Введите имя актера:")
				fmt.Scan(&actor.A_name)
				fmt.Println("Введите фамилию актера:")
				fmt.Scan(&actor.A_surname)
				fmt.Println("Введите дату рождения актера:")
				fmt.Scan(&actor.A_bdate)
				fmt.Println("Введите пол актера:")
				fmt.Scan(&actor.A_gender)
				err := ctrl.CreateActor(actor)
				if err != nil {
					log.Error(err)
					break
				}
				fmt.Println("Введите сериал, в котором снимался актер:")
				var serial int
				fmt.Scan(&serial)
				ctrlSerialsActors := controllers.NewSerialsActorsCtrl(repositories.NewSerialsActorsRepo(db, log))
				serialActor := &models.SerialsActors{}
				serialActor.SetIdActor(actor.GetId())
				serialActor.SetIdSerial(serial)
				err = ctrlSerialsActors.CreateSerialsActors(serialActor)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User added actor")
				}
			}
		// изменить актера
		case 18:
			{
				if curUser == nil {
					log.Warn("User tried to update actor without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to update actor without admin rights")
					fmt.Println("Только администратор может изменять актеров")
					break
				}
				ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(db, log))
				var id int
				fmt.Println("Введите id актера:")
				fmt.Scan(&id)
				actor, err := ctrl.GetActorById(id)
				if err != nil {
					log.Error(err)
					break
				}
				fmt.Println("Поля, доступные для изменения:\n" +
					"1. Имя\n" +
					"2. Фамилия\n" +
					"3. Дата рождения\n" +
					"4. Пол\n\n" +
					"Введите номер поля, которое хотите изменить:")
				var field int
				fmt.Scan(&field)
				switch field {
				case 1:
					{
						fmt.Println("Введите новое имя:")
						fmt.Scan(&actor.A_name)
					}
				case 2:
					{
						fmt.Println("Введите новую фамилию:")
						fmt.Scan(&actor.A_surname)
					}
				case 3:
					{
						fmt.Println("Введите новую дату рождения:")
						fmt.Scan(&actor.A_bdate)
					}
				case 4:
					{
						fmt.Println("Введите новый пол:")
						fmt.Scan(&actor.A_gender)
					}
				}
				err = ctrl.UpdateActor(actor)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User updated actor")
				}
			}
		// удалить актера
		case 19:
			{
				if curUser == nil {
					log.Warn("User tried to delete actor without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to delete actor without admin rights")
					fmt.Println("Только администратор может удалять актеров")
					break
				}
				ctrl := controllers.NewActorsCtrl(repositories.NewActorsRepo(db, log))
				actor := &models.Actors{}
				fmt.Println("Введите id актера:")
				fmt.Scan(&actor.A_id)
				err := ctrl.DeleteActor(actor.GetId())
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User deleted actor")
				}
			}
		// добавить режиссера
		case 20:
			{
				if curUser == nil {
					log.Warn("User tried to add producer without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to add producer without admin rights")
					fmt.Println("Только администратор может добавлять режиссеров")
					break
				}
				ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(db, log))
				producer := &models.Producers{}
				fmt.Println("Введите имя режиссера:")
				fmt.Scan(&producer.P_name)
				fmt.Println("Введите фамилию режиссера:")
				fmt.Scan(&producer.P_surname)
				err := ctrl.CreateProducer(producer)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User added producer")
				}
			}
		// изменить режиссера
		case 21:
			{
				if curUser == nil {
					log.Warn("User tried to update producer without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to update producer without admin rights")
					fmt.Println("Только администратор может изменять режиссеров")
					break
				}
				ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(db, log))
				var id int
				fmt.Println("Введите id режиссера:")
				fmt.Scan(&id)
				producer, err := ctrl.GetProducerById(id)
				if err != nil {
					log.Error(err)
					break
				}
				fmt.Println("Поля, доступные для изменения:\n" +
					"1. Имя\n" +
					"2. Фамилия\n\n" +
					"Введите номер поля, которое хотите изменить:")
				var field int
				fmt.Scan(&field)
				switch field {
				case 1:
					{
						fmt.Println("Введите новое имя:")
						fmt.Scan(&producer.P_name)
					}
				case 2:
					{
						fmt.Println("Введите новую фамилию:")
						fmt.Scan(&producer.P_surname)
					}
				}
				err = ctrl.UpdateProducer(producer)
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User updated producer")
				}
			}
		// удалить режиссера
		case 22:
			{
				if curUser == nil {
					log.Warn("User tried to delete producer without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to delete producer without admin rights")
					fmt.Println("Только администратор может удалять режиссеров")
					break
				}
				ctrl := controllers.NewProducersCtrl(repositories.NewProducersRepo(db, log))
				producer := &models.Producers{}
				fmt.Println("Введите id режиссера:")
				fmt.Scan(&producer.P_id)
				err := ctrl.DeleteProducer(producer.GetId())
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User deleted producer")
				}
			}
		// посмотреть список пользователей
		case 23:
			{
				if curUser == nil {
					log.Warn("User tried to check users list without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to check users list without admin rights")
					fmt.Println("Только администратор может просматривать список пользователей")
					break
				}
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				users, err := ctrl.GetUsers()
				if err != nil {
					log.Error(err)
				} else {
					for _, u := range users {
						fmt.Println(*u)
					}
					log.Info("User checked users list")
				}
			}
		// удалить пользователя
		case 24:
			{
				if curUser == nil {
					log.Warn("User tried to delete user without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to delete user without admin rights")
					fmt.Println("Только администратор может удалять пользователей")
					break
				}
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				user := &models.Users{}
				fmt.Println("Введите id пользователя:")
				fmt.Scan(&user.U_id)
				err := ctrl.DeleteUser(user.GetId())
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User deleted user")
				}
			}
		// выдать права администратора
		case 25:
			{
				if curUser == nil {
					log.Warn("User tried to grant admin rights without login")
					fmt.Println("Необходимо войти в аккаунт")
					break
				}
				if curUser.GetRole() != "admin" {
					log.Warn("User tried to grant admin rights without admin rights")
					fmt.Println("Только администратор может выдавать права администратора")
					break
				}
				ctrl := controllers.NewUsersCtrl(repositories.NewUsersRepo(db, log), repositories.NewFavouritesRepo(db, log))
				user := &models.Users{}
				fmt.Println("Введите id пользователя:")
				fmt.Scan(&user.U_id)
				err := ctrl.GrantAdmin(user.GetId())
				if err != nil {
					log.Error(err)
				} else {
					log.Info("User granted admin rights")
				}
			}
		}
	}
}

func Menu() int {
	menu := "0. Выйти из программы\n\n" +
		// all
		"1. Зарегистрироваться\n" +
		"2. Войти в аккаунт\n" +
		"3. Выйти из аккаунта\n" +
		"4. Посмотреть список сериалов\n" +
		"5. Найти сериал\n\n" +
		// user
		"6. Добавить отзыв\n" +
		"7. Изменить отзыв\n" +
		"8. Удалить отзыв\n" +
		"9. Сохранить сериал в избранное\n" +
		"10. Посмотреть список избранных сериалов\n" +
		"11. Посмотреть профиль\n" +
		"12. Редактировать профиль\n" +
		"13. Посмотреть историю просмотров\n\n" +
		// admin
		"14. Добавить сериал\n" +
		"15. Изменить сериал\n" +
		"16. Удалить сериал\n" +
		"17. Добавить актера\n" +
		"18. Изменить актера\n" +
		"19. Удалить актера\n" +
		"20. Добавить режиссера\n" +
		"21. Изменить режиссера\n" +
		"22. Удалить режиссера\n" +
		"23. Посмотреть список пользователей\n" +
		"24. Удалить пользователя\n" +
		"25. Выдать права администратора\n\n"
	fmt.Println(menu)
	var req int
	fmt.Scan(&req)
	return req
}
