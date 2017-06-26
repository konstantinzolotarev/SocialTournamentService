package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"encoding/json"
)

type Player struct {
	id int64
	PlayerName string `json:"playerName"`
	Points float64    `json:"balance"`
}

func Take(ctx context.Context, playerName string, points string) (error) {
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

func PrintPlayers(ctx context.Context) (error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return errors.New("models: could not get database connection pool from context")
	}

	rows, err := db.Query("SELECT id, \"playerName\", points FROM game.\"Players\"")
	if err != nil {
		return err
	}
	defer rows.Close()

	players := make([]*Player, 0)
	for rows.Next() {
		player := new(Player)
		err := rows.Scan(&player.id, &player.PlayerName, &player.Points)
		if err != nil {
			return err
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	for i := 0; i < len(players); i++ {
		fmt.Println("Id: ", players[i].id, "playerName: ", players[i].PlayerName, "points: ", players[i].Points)
	}

	return nil
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