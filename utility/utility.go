package utility

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-api/models"
	"simple-api/payload"
)

func HandelRequest(db *sql.DB, handel func(w http.ResponseWriter, r *http.Request, db *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handel(w, r, db)
	}
}

func ResponJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

}

func ResponErrorJSON(w http.ResponseWriter, status int, message string) {
	ResponJSON(w, status, map[string]string{"message": message})
}

func InsertProduct(item *payload.Product, db *sql.DB) (string, error) {
	var product models.Product
	product.Set(item)
	row, err := product.Insert(db)
	if err != nil {
		log.Panic("Data Product gagal ditambahkan")
		return "", err
	}
	if row > 0 {
		fmt.Println("Data Product berhasil ditambahan")
	}

	item.ID = product.Get().ID
	return product.Get().ID, nil
}

func InsertLog(item *payload.Log, db *sql.DB) (string, error) {
	var loging models.Log
	loging.Set(item)

	row, err := loging.Insert(db)
	if err != nil {
		log.Panic("Data log gagal ditambahkan")
		return "", err
	}
	if row > 0 {
		fmt.Println("Data log berhasil ditambahan")
	}

	item.ID = loging.Get().ID
	return loging.Get().ID, nil
}

func UpdateLog(item *payload.Log, db *sql.DB) error {
	var loging models.Log
	loging.Set(item)

	row, err := loging.Update(db)
	if err != nil || row <= 0 {
		newerr := errors.New("Data failed to update")
		return newerr
	}

	return nil

}

func UpdateImagePath(item *payload.Product, db *sql.DB) error {
	var product models.Product
	product.Set(item)

	row, err := product.UpdatePath(db)
	if err != nil || row <= 0 {
		newerr := errors.New("Data failed to update")
		return newerr
	}

	return nil
}
