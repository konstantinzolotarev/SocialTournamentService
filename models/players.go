package models

import (
	"context"
	"database/sql"
	"errors"
	"encoding/json"
)

type Player struct {
	id int64
	PlayerName string `json:"playerName"`
	Points float64    `json:"balance"`
}

func Fund(ctx context.Context, playerName string, points string) (error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return errors.New("models: could not get database connection pool from context")
	}

	_, err := db.Exec("INSERT INTO game.\"Players\" (\"playerName\", \"points\") " +
		"VALUES ($1, $2) ON CONFLICT (\"playerName\") DO UPDATE SET points = excluded.points;", playerName, points)
	if err != nil {
		return errors.New("models: could not write Players to database")
	}

	return nil
}

func Take(ctx context.Context, playerName string, points string) (int64, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return 0, errors.New("models: could not get database connection pool from context")
	}

	result, err := db.Exec("UPDATE game.\"Players\" SET points = points - $1 WHERE \"playerName\" = $2 AND points - $1 > 0", points, playerName)
	if err != nil {
		return 0, errors.New("models: could not write Players to database")
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("models: error response count row affected")
	}

	if row == 1 {
		return row, nil
	} else {
		return row, errors.New("models: wrong count row affected")
	}
}

func Balance(ctx context.Context, playerName string) ([]byte, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("models: could not get database connection pool from context")
	}

	rows, err := db.Query("SELECT \"playerName\", points FROM game.\"Players\" WHERE \"playerName\" = $1",
		playerName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		p := new(Player)
		err := rows.Scan(&p.PlayerName, &p.Points)
		if err != nil {
			return nil, err
		}

		return json.Marshal(p)
	} else {
		return nil, errors.New("models: wrong player name")
	}
}