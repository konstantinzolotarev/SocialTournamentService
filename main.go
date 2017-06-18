package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"log"
)

type Config struct {
	DbUser string
	DbPassword string
	DbName string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
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
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err)
	defer db.Close()

	http.HandleFunc("/take", take)
	http.HandleFunc("/fund", fund)
	http.HandleFunc("/announceTournament", announceTournament)
	http.HandleFunc("/balance", balance)
	http.HandleFunc("/resultTournament", resultTournament)
	http.ListenAndServe(":8081", nil)
}

func take(w http.ResponseWriter, r *http.Request) {
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