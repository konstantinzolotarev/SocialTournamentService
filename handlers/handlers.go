package handlers

import (
	"net/http"
	"SocialTournamentService/models"
	"context"
)

func Take(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func Fund(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func AnnounceTournament(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func Balance(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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

func ResultTournament(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response, err := models.ResultTournament(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func JoinTournament(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := models.JoinTournament(ctx, r.URL.Query())
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
}
