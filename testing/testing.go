package main

import (
	"database/sql"
	"fmt"
	"log"
	"simple-api/models"
	ctx "simple-api/models/context"
	"simple-api/payload"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

var context ctx.DbContext

func insertEmployee(item *payload.Employee, db *sql.DB) string {
	var employe models.Employee
	employe.Set(item)
	row, err := employe.Insert(db)
	if err != nil {
		log.Panic("Data Employe gagal ditambahkan")
	}
	if row > 0 {
		fmt.Println("Data Employee berhasil ditambahan")
	}

	item.ID = employe.Get().ID
	return employe.Get().ID
}

func insertProduct(item *payload.Product, db *sql.DB) string {
	var product models.Product
	product.Set(item)
	row, err := product.Insert(db)
	if err != nil {
		log.Panic("Data Product gagal ditambahkan")
		return ""
	}
	if row > 0 {
		fmt.Println("Data Product berhasil ditambahan")
	}

	item.ID = product.Get().ID
	return product.Get().ID
}

func insertLog(item *payload.Log, db *sql.DB) string {
	var loging models.Log
	loging.Set(item)

	row, err := loging.Insert(db)
	if err != nil {
		log.Panic("Data log gagal ditambahkan")
		return ""
	}
	if row > 0 {
		fmt.Println("Data log berhasil ditambahan")
	}

	item.ID = loging.Get().ID
	return loging.Get().ID
}

func updateLog(item *payload.Log, db *sql.DB) {
	var loging models.Log
	loging.Set(item)

	row, err := loging.Update(db)
	if err != nil {
		log.Panic("Data log gagal diupdate")
	}
	if row > 0 {
		fmt.Println("Data log berhasil diupdate")
	}
}

func main() {

	context = ctx.DbContext{

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
		Username: "salimSekfdadfli",
		Position: "Staff Gudang",
	}

	userID := insertEmployee(&user, db)
	fmt.Println(userID)

	baso := payload.Product{
		Name:  "Baso Urat",
		Code:  "ramsikud",
		Price: 3500,
	}

	productID := insertProduct(&baso, db)
	fmt.Println(productID)

	loging := payload.Log{
		Product: baso,
		Action:  "i",
		Amount:  54,
		Admin:   user,
	}

	logID := insertLog(&loging, db)
	fmt.Println(logID)

	products, err := models.SelectAllProduct(db)
	if err != nil {
		log.Panic("Tidak dapat menemukan data")
	}
	fmt.Println(products)

	logs, err := models.SelectLogByIDProduct(productID, db)
	if err != nil {
		log.Panic("Tidak dapat menemukan data")
	}
	fmt.Println(logs)

	newUser := payload.Employee{
		Name:     "Alif budiono",
		Username: "sdsdasdhgs",
		Position: "Kepala Gudang",
	}

	insertEmployee(&newUser, db)

	loging.Amount = 76
	loging.Admin = newUser

	updateLog(&loging, db)

	id := uuid.New()
	xid := xid.New()

	fmt.Println(id.String())
	fmt.Println(xid.String())
}
