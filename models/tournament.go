package models

import (
	"context"
	"database/sql"
	"errors"
)

type tournament struct {
	Id uint64       `json:"id"`
	Deposit float64 `json:"deposit"`
	statusId uint64
}

const (
	announcement  = 1
	finished  = 2
)

func AnnounceTournament(ctx context.Context, id string, deposit string) (error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return errors.New("models: could not get database connection pool from context")
	}

	_, err := db.Exec("INSERT INTO game.\"Tournaments\" (\"id\", \"deposit\", \"statusId\") " +
		"VALUES ($1, $2, $3);", id, deposit, announcement)
	if err != nil {
		return errors.New("models: could not write Players to database")
	}

	return nil
}