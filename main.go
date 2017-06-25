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
	//http.HandleFunc("/take", take)
	http.HandleFunc("/fund", fund)
	http.HandleFunc("/announceTournament", announceTournament)
	http.HandleFunc("/balance", balance)
	http.HandleFunc("/resultTournament", resultTournament)
	http.ListenAndServe(":8081", nil)
}

func take(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	models.PrintPlayers(ctx)
	fmt.Println("Get query params: ", r.URL.Query())
	b, err := json.Marshal("Hello client")
	if err != nil {
		panic(err)
	}

	w.Write(b)

}

func fund(w http.ResponseWriter, r *http.Request) {

}

func announceTournament(w http.ResponseWriter, r *http.Request) {

}

func balance(w http.ResponseWriter, r *http.Request) {

}

func resultTournament(w http.ResponseWriter, r *http.Request) {

}