package main

import (
	"fmt"
	"log"
	"simple-api/models"
	ctx "simple-api/models/context"
	"simple-api/payload"
)

func main() {
	context := ctx.DbContext{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "hadioz",
		Password: "ozmonday",
		DbName:   "gudang",
	}

	db, err := context.ConnectDB()
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	barang := payload.Barang{
		Name:  "Mie Ayam",
		Price: 12000,
	}

	user := payload.User{
		Name:     "gofursasa",
		Username: "gofurur",
		Password: "haloukoni",
		Role:     "admin",
	}

	var pdct models.Barang
	var usr models.User
	usr.Set(&user)
	pdct.Set(&barang)
	row, err := pdct.Insert(db)
	if err != nil {
		log.Fatalf("Tidak dapat menambah data : %s", err)
	}
	fmt.Println(row)
	row, err = usr.Insert(db)
	if err != nil {
		log.Fatalf("Tidak dapat menambah data : %s", err)
	}
	fmt.Println(row)
	// fmt.Println(usr)

	loge := payload.LogBarang{
		Name:   pdct.Get(),
		Action: "i",
		Amount: 344,
		Admin:  usr.Get(),
	}

	// fmt.Println(loge)
	var loging models.LogBarang
	loging.Set(&loge)
	row, err = loging.Insert(db)

	if err != nil {
		log.Fatalf("Tidak dapat menambah data : %s", err)
	}

	fmt.Println(row)
	defer db.Close()
}
