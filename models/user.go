package models

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"simple-api/payload"
	"strconv"
	"strings"
	"time"
)

type User struct {
	id       string
	name     string
	username string
	password string
	role     string
}

func (u *User) Set(item *payload.User) {
	u.id = item.ID
	u.name = item.Name
	u.role = item.Role
	u.username = item.Username
	u.password = item.Password
}

func (u *User) Get() payload.User {
	return payload.User{
		ID:       u.id,
		Name:     u.name,
		Username: u.username,
		Role:     u.role,
	}
}

func (u *User) createID() {
	time := time.Now()
	year := strconv.Itoa(time.Year())
	month := strings.ToUpper(time.Month().String())
	nano := strconv.Itoa(time.Nanosecond())

	var productID []string
	productID = append(productID, "User")
	productID = append(productID, year)
	productID = append(productID, month[0:3])
	productID = append(productID, nano)

	ID := strings.Join(productID, "-")
	encode := base64.StdEncoding.EncodeToString([]byte(ID))
	u.id = encode
}

func (u *User) Insert(db *sql.DB) (int64, error) {
	u.createID()

	query := fmt.Sprintf(`INSERT INTO public.users(id, name, role, username, password)
		VALUES ('%s', '%s', '%s', '%s', '%s');`, u.id, u.name, u.role, u.username, u.password)
	// fmt.Print(query)
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	effect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return effect, nil
}
