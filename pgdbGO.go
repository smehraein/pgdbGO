/*
Utility package for managing Postgres database interactions on a basic Go server.
*/
package pgdbGO

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type PGConnection struct {
	DB_USER, DB_PASSWORD, DB_NAME, DB_HOST string
}

// Opens a DB connection when passed an appropriate set of parameters.
// Checks the connection and returns an error if it's faulty.
func Connect(c PGConnection) (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_HOST)
	db, _ = sql.Open("postgres", dbinfo)
	err = db.Ping()
	return
}

// Creates a new user in Postgres when provided a db connection and appropriate parameters.
func CreateNewUser(db *sql.DB, name, password, email, sportsTeam string, gender int) (err error) {
	var lastInsertId int
	err = db.QueryRow(`INSERT INTO userinfo(name, password, email, sportsTeam, gender, created) VALUES($1, $2, $3, $4, $5, $6)
										returning uid;`, name, password, email, sportsTeam, gender, time.Now()).Scan(&lastInsertId)
	return
}
