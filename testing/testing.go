package main

import (
	"fmt"
	"log"
	"simple-api/models"
	ctx "simple-api/models/context"
	"simple-api/payload"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

func main() {
	context := ctx.DbContext{

		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "hadioz",
		Password: "ozmonday",
		DbName:   "inventory",
	}

	db, err := context.ConnectDB()
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	user := payload.Employee{
		Name:     "Salim jusuf",
		Username: "salmjsdsfkssdfgagsdurs",
		Position: "Staff Gudang",
	}

	cimol := payload.Product{
		Name:  "Cimol Pak Muh",
		Code:  "sadadadhsa",
		Price: 3500,
	}

	var usr models.Employee
	usr.Set(&user)

	var prd models.Product
	prd.Set(&cimol)

	var loge models.Log

	row, err := prd.Insert(db)
	if err != nil {
		log.Fatalf("Tidak dapat menambah data : %s", err)
	}
	fmt.Println(row)
	fmt.Println(prd)
	cimol.ID = prd.Get().ID
	row, err = usr.Insert(db)
	if err != nil {
		log.Fatalf("Tidak dapat menambah data : %s", err)
	}
	fmt.Println(row)
	fmt.Println(usr)
	user.ID = usr.Get().ID

	loging := payload.Log{
		Name:   cimol,
		Action: "i",
		Amount: 54,
		Admin:  user,
	}

	loge.Set(&loging)
	row, err = loge.Insert(db)
	if err != nil {
		log.Fatalf("Tidak dapat menambah data log : %s", err)
	}
	fmt.Println(row)

	products, err := models.SelectAllProduct(db)
	if err != nil {
		log.Fatalf("Tidak dapat menemukan data : %s", err)
	}
	fmt.Println(products)

	id := uuid.New()
	xid := xid.New()

	fmt.Println(id.String())
	fmt.Println(xid.String())
}
