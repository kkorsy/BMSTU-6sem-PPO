package main

import (
	"app/config"
	"app/internal/server"
	"app/logger"
	"context"

	"net/http"
	"time"

	// "app/tech_ui"

	log_default "log"

	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Config{}
	_, err := toml.DecodeFile("config/config.toml", &cfg)
	if err != nil {
		log_default.Fatal(err)
	}
	log, err := logger.InitLog(cfg.Log_path)
	if err != nil {
		log.Fatal(err)
	}

	var db interface{}
	switch cfg.Db_type {
	case "postgres":
		{
			client, err := sqlx.Connect(cfg.Db_type, cfg.Db_url)
			if err != nil {
				log.Fatal(err)
			}
			db = client
			defer client.Close()
			log.Info("Successfully connected to Postgres")
		}
	case "mongo":
		{
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://kkorsy:5454038mmm@cluster0.xje41xy.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if err = client.Disconnect(ctx); err != nil {
					log.Fatal(err)
				}
			}()

			err = client.Ping(ctx, nil)
			if err != nil {
				log.Fatal(err)
			}
			db = client
			log.Info("Successfully connected to MongoDB")
		}
	default:
		{
			log.Fatal("Unknown db type")
		}
	}

	// tech_ui.Run(db, log)

	s := server.NewServer(log, db, cfg.SessionKey)

	err = http.ListenAndServe(cfg.Port, s)
	if err != nil {
		log.Fatal(err)
	}
}
