package handler

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"simple-api/models"
	"simple-api/payload"
	"simple-api/utility"
	"strings"

	"github.com/google/uuid"
)

func Halloworld(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Write([]byte("Halo World"))
}

func SignIn(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method POST")
		return
	}

	decoder := json.NewDecoder(r.Body)
	payload := payload.Employee{}

	if err := decoder.Decode(&payload); err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data structure wrong")
		return
	}

	employee := models.Employee{}
	employee.Set(&payload)

	account, err := employee.UserValidation(db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Username not found")
		return
	}

	utility.ResponJSON(w, http.StatusOK, map[string]string{"user-id": account})
}

func SelectAllProduct(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method GET")
		return
	}

	products, err := models.SelectAllProduct(db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusInternalServerError, "Data product can't be presened")
	}

	utility.ResponJSON(w, http.StatusOK, products)
}

func InsertProduct(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "POST" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method POST")
		return
	}

	decoder := json.NewDecoder(r.Body)
	payload := payload.Product{}

	if err := decoder.Decode(&payload); err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data structure wrong")
		return
	}

	productID, err := utility.InsertProduct(&payload, db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data faild to add")
		return
	}

	utility.ResponJSON(w, http.StatusOK, map[string]string{"id-product": productID})
}

func InsertLog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method POST")
		return
	}

	decoder := json.NewDecoder(r.Body)
	payload := payload.Log{}

	if err := decoder.Decode(&payload); err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data structure wrong")
		return
	}

	logID, err := utility.InsertLog(&payload, db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data faild to add")
		return
	}

	utility.ResponJSON(w, http.StatusOK, map[string]string{"id-log": logID})
}

func SelectLog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method GET")
		return
	}

	productID := r.URL.Query().Get("id-product")

	logs, err := models.SelectLogByIDProduct(productID, db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusInternalServerError, "Data product can't be presened")
	}

	utility.ResponJSON(w, http.StatusOK, logs)
}

func SelectLogHistory(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method GET")
		return
	}

	logID := r.URL.Query().Get("id-log")

	logs, err := models.SelectLogHistory(logID, db)
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusInternalServerError, "Log history can't be presened")
	}

	utility.ResponJSON(w, http.StatusOK, logs)
}

func UpdateLog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method POST")
		return
	}

	decoder := json.NewDecoder(r.Body)
	payload := payload.Log{}

	if err := decoder.Decode(&payload); err != nil {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Data structure wrong")
		return
	}

	if err := utility.UpdateLog(&payload, db); err != nil {
		utility.ResponErrorJSON(w, http.StatusInternalServerError, "Data Failed to update")
		return
	}

	utility.ResponJSON(w, http.StatusOK, map[string]string{"message": "Data was Updated"})
}

func UploadImageProduct(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		utility.ResponErrorJSON(w, http.StatusBadRequest, "Only allow method POST")
		return
	}
	productID := r.URL.Query().Get("id-product")
	reader, err := r.MultipartReader()
	if err != nil {
		utility.ResponErrorJSON(w, http.StatusInternalServerError, "internal server error")
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		filename := uuid.New().String()
		ext := filepath.Ext(part.FileName())

		payload := payload.Product{
			ID:        productID,
			PathImage: strings.Join([]string{filename, ext}, ""),
		}

		location := "/home/hadioz/development/golang/simple-api/image"

		filelocation := filepath.Join(location, payload.PathImage)
		file, err := os.OpenFile(filelocation, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		if err != nil {
			log.Panic(err)
			return
		}

		if _, err := io.Copy(file, part); err != nil {
			utility.ResponErrorJSON(w, http.StatusInternalServerError, "Can't save file")
			return
		}

		if err := utility.UpdateImagePath(&payload, db); err != nil {
			utility.ResponErrorJSON(w, http.StatusInternalServerError, "Image failed to save  chack your product id")
			return
		}
	}

	utility.ResponJSON(w, http.StatusOK, map[string]string{"message": "Data was saved"})

}
