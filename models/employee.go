package models

import (
	"database/sql"
	"errors"
	"fmt"
	"simple-api/payload"

	"github.com/rs/xid"
)

type Employee struct {
	id       string
	name     string
	username string
	password string
	position string
}

func (e *Employee) Set(item *payload.Employee) {
	e.id = item.ID
	e.name = item.Name
	e.position = item.Position
	e.username = item.Username
	e.password = item.Password
}

func (e *Employee) Get() payload.Employee {
	return payload.Employee{
		ID:       e.id,
		Name:     e.name,
		Username: e.username,
		Position: e.position,
	}
}

func (e *Employee) Insert(db *sql.DB) (int64, error) {
	id := xid.New().String()
	query := fmt.Sprintf(`INSERT INTO public.employees(id_employee, name, billet, username, password)
		VALUES ('%s', '%s', '%s', '%s', '%s');`, id, e.name, e.position, e.username, e.password)
	fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	effect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	e.id = id
	return effect, nil
}

func (e *Employee) UserValidation(db *sql.DB) (string, error) {
	query := fmt.Sprintf(`SELECT id_employee, password FROM public.employees WHERE username = '%s';`, e.username)
	row := db.QueryRow(query)
	var res payload.Employee

	if err := row.Scan(&res.ID, &res.Password); err != nil {
		return "", err
	}

	if res.Password != e.password {
		notvalid := errors.New("Password wrong")
		return "", notvalid
	}

	return res.ID, nil
}
