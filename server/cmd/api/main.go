package main

import (
	"github.com/tmplam/movseek/config"
	"github.com/tmplam/movseek/internal/appconfig/mongo"
	"github.com/tmplam/movseek/internal/httpserver"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	client, err := mongo.Connect(cfg.Mongo.URI)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(client)

	db := client.Database(cfg.Mongo.DBName)

	srv := httpserver.New(httpserver.Config{
		Port: cfg.HTTPServer.Port,
		DB:   db,
	})
	srv.Run()
}
