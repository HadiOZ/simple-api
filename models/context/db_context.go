package context

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbContext struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (db *DbContext) ConnectDB() (*sql.DB, error) {
	user := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.DbName)
	conn, err := sql.Open("postgres", user)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
