package integration_test

import (
	"app/config"
	"app/internal/controllers"
	"app/internal/models"
	"app/internal/repositories"
	"app/logger"
	"log"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetActors(t *testing.T) {
	cfg := config.Config{}
	_, err := toml.DecodeFile("D:/BMSTU/PPO/src/config/config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log, err := logger.InitLog("D:/BMSTU/PPO/src/logger/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewActorsRepo(db, log)
	actCtrl := controllers.NewActorsCtrl(repo)

	actors, err := actCtrl.GetActors()

	assert.NoError(t, err)
	assert.NotNil(t, actors)
}

func TestGetCommentsById(t *testing.T) {
	cfg := config.Config{}
	_, err := toml.DecodeFile("D:/BMSTU/PPO/src/config/config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log, err := logger.InitLog("D:/BMSTU/PPO/src/logger/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewCommentsRepo(db, log)
	comCtrl := controllers.NewCommentsCtrl(repo)

	comment, err := comCtrl.GetCommentById(1)

	assert.NoError(t, err)
	assert.NotNil(t, comment)
}

func TestCreateSerial(t *testing.T) {
	cfg := config.Config{}
	_, err := toml.DecodeFile("D:/BMSTU/PPO/src/config/config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log, err := logger.InitLog("D:/BMSTU/PPO/src/logger/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewSerialsRepo(db, log)
	serCtrl := controllers.NewSerialsCtrl(repo)

	err = serCtrl.CreateSerial(&models.Serial{
		S_id:          100,
		S_name:        "Test",
		S_year:        2021,
		S_description: "test descr",
		S_genre:       "test genre",
		S_rating:      1.0,
		S_seasons:     1,
		S_state:       "завершен",
		S_idProducer:  1,
		S_img:         "test img",
		S_duration:    "10:00:00",
	})

	assert.NoError(t, err)
}

func TestUpdateSeasons(t *testing.T) {
	cfg := config.Config{}
	_, err := toml.DecodeFile("D:/BMSTU/PPO/src/config/config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log, err := logger.InitLog("D:/BMSTU/PPO/src/logger/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewSeasonsRepo(db, log)
	ssCtrl := controllers.NewSeasonsCtrl(repo)

	err = ssCtrl.UpdateSeason(&models.Seasons{Ss_id: 1, Ss_name: "Test", Ss_date: "2021-01-01", Ss_idSerial: 1, Ss_num: 1, Ss_cntEpisodes: 1})

	assert.NoError(t, err)
}

func TestDeleteEpisode(t *testing.T) {
	cfg := config.Config{}
	_, err := toml.DecodeFile("D:/BMSTU/PPO/src/config/config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log, err := logger.InitLog("D:/BMSTU/PPO/src/logger/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.Db_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewEpisodesRepo(db, log)
	epCtrl := controllers.NewEpisodesCtrl(repo)

	err = epCtrl.DeleteEpisode(1)

	assert.NoError(t, err)
}
