package models

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
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
		println(err)
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

	row := transaction.QueryRow("SELECT deposit / $1 FROM game.\"Tournaments\" WHERE \"id\" = $2",
		countPlayers, tournamentIds[0])

	var share float64
	err = row.Scan(&share)
	if err != nil {
		println(err.Error())
		transaction.Rollback()
		return errors.New("model: can't calculate share")
	}

	_, err = transaction.Exec("INSERT INTO game.\"TournamentPlayers\" (\"playerId\", \"tournamentId\", share) " +
		"VALUES ((SELECT id FROM game.\"Players\" WHERE \"playerName\" = $1), $2, $3)", playerIds[0], tournamentIds[0], share)
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
				"), (SELECT id FROM game.\"Players\" WHERE \"playerName\" = $2), $3, $4)"

			_, err = transaction.Exec(sql + values, id, playerIds[0], tournamentIds[0],
				strconv.FormatFloat(share, 'f', -1, 64))
			if err != nil {
				println(err.Error())
				transaction.Rollback()
				return errors.New("model: error insert into Backers")
			}

			result, err := transaction.Exec("UPDATE game.\"Players\" SET points = points - $1 " +
				"WHERE \"playerName\" = $2 AND points - $1 > 0", share, id)
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