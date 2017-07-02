package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"context"
	"SocialTournamentService/models"
	"SocialTournamentService/handlers"
)

type Config struct {
	DbUser string
	DbPassword string
	DbName string
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)

	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configuration.DbUser, configuration.DbPassword, configuration.DbName)
	db, err := models.ConnectToDB(dbInfo)
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/take", &ContextAdapter{ctx, ContextHandlerFunc(handlers.Take)})
	http.Handle("/fund", &ContextAdapter{ctx, ContextHandlerFunc(handlers.Fund)})
	http.Handle("/announceTournament", &ContextAdapter{ctx, ContextHandlerFunc(handlers.AnnounceTournament)})
	http.Handle("/balance", &ContextAdapter{ctx, ContextHandlerFunc(handlers.Balance)})
	http.Handle("/resultTournament", &ContextAdapter{ctx, ContextHandlerFunc(handlers.ResultTournament)})
	http.Handle("/joinTournament", &ContextAdapter{ctx, ContextHandlerFunc(handlers.JoinTournament)})
	http.ListenAndServe(":8081", nil)
}
