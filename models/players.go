package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Player struct {
	Id int64
	playerName string
	points float64
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
		err := rows.Scan(&player.Id, &player.playerName, &player.points)
		if err != nil {
			return err
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	for i := 0; i < len(players); i++ {
		fmt.Println("Id: ", players[i].Id, "playerName: ", players[i].playerName, "points: ", players[i].points)
	}

	return nil
}