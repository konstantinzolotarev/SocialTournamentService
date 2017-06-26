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
	http.HandleFunc("/fund", fund)
	http.HandleFunc("/announceTournament", announceTournament)
	http.HandleFunc("/balance", balance)
	http.HandleFunc("/resultTournament", resultTournament)
	http.ListenAndServe(":8081", nil)
}

func take(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	models.PrintPlayers(ctx)
	fmt.Println("Get query params: ", r.URL.Query())

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

	err := models.Take(ctx, playerId, points)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func fund(w http.ResponseWriter, r *http.Request) {

}

func announceTournament(w http.ResponseWriter, r *http.Request) {

}

func balance(w http.ResponseWriter, r *http.Request) {

}

func resultTournament(w http.ResponseWriter, r *http.Request) {

}