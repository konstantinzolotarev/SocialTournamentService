package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"context"
	"SocialTournamentService/models"
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

	fmt.Println(configuration.DbUser)

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configuration.DbUser, configuration.DbPassword, configuration.DbName)
	db, err := models.ConnectToDB(dbInfo)
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/take", &ContextAdapter{ctx, ContextHandlerFunc(take)})
	http.Handle("/fund", &ContextAdapter{ctx, ContextHandlerFunc(fund)})
	http.Handle("/announceTournament", &ContextAdapter{ctx, ContextHandlerFunc(announceTournament)})
	http.Handle("/balance", &ContextAdapter{ctx, ContextHandlerFunc(balance)})
	http.Handle("/resultTournament", &ContextAdapter{ctx, ContextHandlerFunc(resultTournament)})
	http.ListenAndServe(":8081", nil)
}

func take(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playerId := r.URL.Query().Get("playerId")
	points := r.URL.Query().Get("points")
	if len(playerId) == 0 || len(points) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	count, err := models.Take(ctx, playerId, points)
	if err != nil {
		println(err.Error())
		if count != 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func fund(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playerId := r.URL.Query().Get("playerId")
	points := r.URL.Query().Get("points")
	if len(playerId) == 0 || len(points) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := models.Fund(ctx, playerId, points)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func announceTournament(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tournamentId := r.URL.Query().Get("tournamentId")
	deposit := r.URL.Query().Get("deposit")
	if len(tournamentId) == 0 || len(deposit) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := models.AnnounceTournament(ctx, tournamentId, deposit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func balance(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	playerId := r.URL.Query().Get("playerId")
	if len(playerId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := models.Balance(ctx, playerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func resultTournament(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}