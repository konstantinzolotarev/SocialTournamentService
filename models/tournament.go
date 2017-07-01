package models

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
	"math/rand"
	"encoding/json"
	"fmt"
)

const (
	announcement  = 1
	finished  = 2
)

type winner struct {
	PlayerName string `json:"playerId"`
	Share float64     `json:"prize"`
}

type responseWinner struct {
	TournamentName uint64   `json:"tournamentId"`
	Winners        []winner `json:"winners"`
}

func AnnounceTournament(ctx context.Context, id string, deposit string) (error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return errors.New("models: could not get database connection pool from context")
	}

	_, err := db.Exec("INSERT INTO game.\"Tournaments\" (\"tournamentNumber\", \"deposit\", \"statusId\") " +
		"VALUES ($1, $2, $3);", id, deposit, announcement)
	if err != nil {
		println(err.Error())
		return errors.New("models: could not write Players to database")
	}

	return nil
}

func JoinTournament(ctx context.Context, values map[string][]string) (error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return errors.New("models: could not get database connection pool from context")
	}

	transaction, err := db.Begin()
	if err != nil {
		println(err.Error())
		return errors.New("models: error get transaction")
	}

	tournamentIds, ok := values["tournamentId"]
	if !ok {
		return errors.New("In query not found tournamentIds")
	} else if len(tournamentIds) > 1 {
		return errors.New("More than one parameter tournamentIds")
	}

	playerIds, ok := values["playerId"]
	if !ok {
		return errors.New("In query not found playerIds")
	} else if len(playerIds) > 1 {
		return errors.New("More than one parameter playerIds")
	}

	backerIds, ok := values["backerId"]
	countPlayers := 1
	if ok {
		countPlayers += len(backerIds)
	}

	row := transaction.QueryRow("SELECT deposit / $1 FROM game.\"Tournaments\" WHERE \"id\" = " +
		"(SELECT id FROM game.\"Tournaments\" WHERE \"tournamentNumber\" = $2)",
		countPlayers, tournamentIds[0])

	var share float64
	err = row.Scan(&share)
	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return errors.New("model: can't calculate share")
	}

	_, err = transaction.Exec("INSERT INTO game.\"TournamentPlayers\" (\"playerId\", \"tournamentId\", share) " +
		"VALUES ((SELECT id FROM game.\"Players\" WHERE \"playerName\" = $1), " +
		"(SELECT id FROM game.\"Tournaments\" WHERE \"tournamentNumber\" = $2), $3)",
		playerIds[0], tournamentIds[0], share)
	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return errors.New("model: error insert into TournamentPlayers")
	}

	result, err := transaction.Exec("UPDATE game.\"Players\" SET points = points - $1 " +
		"WHERE \"playerName\" = $2 AND points - $1 >= 0", share, playerIds[0])
	c, err := result.RowsAffected()
	if c != 1 {
		transaction.Rollback()
		return errors.New("model: error update points on player. Error row affected")
	}

	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return errors.New("model: error update points on player")
	}

	if countPlayers > 1 {
		sql := "INSERT INTO game.\"Backers\" (\"playerId\", \"supportPlayerId\", \"tournamentId\", share) VALUES "
		for _, id := range backerIds {
			values := " ((SELECT id FROM game.\"Players\" WHERE \"playerName\" = $1" +
				"), (SELECT id FROM game.\"Players\" WHERE \"playerName\" = $2), " +
				"(SELECT id FROM game.\"Tournaments\" WHERE \"tournamentNumber\" = $3), $4)"

			_, err = transaction.Exec(sql + values, id, playerIds[0], tournamentIds[0],
				strconv.FormatFloat(share, 'f', -1, 64))
			if err != nil {
				println(err.Error())
				transaction.Rollback()
				return errors.New("model: error insert into Backers")
			}

			result, err := transaction.Exec("UPDATE game.\"Players\" SET points = points - $1 " +
				"WHERE \"playerName\" = $2 AND points - $1 >= 0", share, id)
			if err != nil {
				println(err.Error())
				transaction.Rollback()
				return errors.New("model: error update points on backers")
			}

			c, err := result.RowsAffected()
			if c != 1 {
				transaction.Rollback()
				return errors.New("model: error update points on backers. Error row affected")
			}
		}
	}

	transaction.Commit()
	return nil
}

func ResultTournament(ctx context.Context) ([]byte, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("models: could not get database connection pool from context")
	}

	transaction, err := db.Begin()
	if err != nil {
		println(err.Error())
		return nil, errors.New("models: error get transaction")
	}

	var status uint64
	var tournamentNumber uint64
	var winnerId uint64

	//check tournament finish
	row := transaction.QueryRow("SELECT \"statusId\", \"tournamentNumber\" " +
		"FROM game.\"Tournaments\" ORDER BY id DESC limit 1")
	err = row.Scan(&status, &tournamentNumber)
	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return nil, err
	}

	if status == announcement {
		//calculate prize and count players
		rows, err := db.Query("SELECT \"playerId\", \"deposit\", \"playerName\" " +
			"FROM game.\"PlayersInTournament\" WHERE \"tournamentNumber\" = $1", tournamentNumber)

		if err != nil {
			println(err)
			return nil, err
		}

		var players []Player
		var prize float64

		for rows.Next() {
			p := new(Player)
			rows.Scan(&p.id, &p.Points, &p.PlayerName)

			players = append(players, *p)
			prize += p.Points
		}

		rand.Seed(time.Now().UTC().UnixNano())
		winIndex := rand.Intn(len(players))

		winnerId = players[winIndex].id
		//load backers info
		rows, err = transaction.Query("SELECT \"playerId\", \"playerName\", \"tournamentId\", " +
			" \"supportPlayerId\" FROM game.\"PlayerBackers\" WHERE \"supportPlayerId\" = $1", players[winIndex].id)

		if err != nil {
			println(err.Error())
			transaction.Rollback()
			return nil, err
		}

		var backers = []Backer{}
		for rows.Next() {
			b := new(Backer)
			rows.Scan(&b.PlayerId, &b.PlayerName, &b.TournamentId, &b.SupportPlayerId)
			backers = append(backers, *b)
		}

		//add points
		bCount := len(backers)
		gain := prize / (float64(bCount) + 1)
		for i:=0; i < bCount; i++ {
			_, err = transaction.Exec("UPDATE game.\"Players\" SET \"points\" = \"points\" + $2" +
				"WHERE \"id\" = $1", backers[i].PlayerId, gain)

			if err != nil {
				println(err.Error())
				transaction.Rollback()
				return nil, err
			}
		}

		_, err = transaction.Exec("UPDATE game.\"Players\" SET \"points\" = \"points\" + $2" +
			"WHERE \"id\" = $1", players[winIndex].id, gain)

		if err != nil {
			println(err.Error())
			transaction.Rollback()
			return nil, err
		}

		_, err = transaction.Exec("UPDATE game.\"Tournaments\" SET \"statusId\" = $1, \"winnerId\" = $2 " +
			"WHERE \"tournamentNumber\" = $3", finished, players[winIndex].id, tournamentNumber)

		if err != nil {
			println(err.Error())
			transaction.Rollback()
			return nil, err
		}
	} else if status == finished {
		//check tournament finish
		row := transaction.QueryRow("SELECT  \"winnerId\" " +
			"FROM game.\"Tournaments\" ORDER BY id DESC limit 1")
		err = row.Scan(&winnerId)
		if err != nil {
			println(err.Error())
			transaction.Rollback()
			return nil, err
		}
	}

	rows, err := transaction.Query("SELECT \"playerName\", \"share\" FROM game.\"Winners\" " +
		"WHERE \"winnerId\" = $1", winnerId)

	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return nil, err
	}

	var playerWinners = []winner{}
	for rows.Next() {
		w := new(winner)
		rows.Scan(&w.PlayerName, &w.Share)
		playerWinners = append(playerWinners, *w)
	}

	var response responseWinner
	response.TournamentName = tournamentNumber
	response.Winners = playerWinners

	transaction.Commit()
	return json.Marshal(response)
}
