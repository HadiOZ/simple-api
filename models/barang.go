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

type Barang struct {
	id        string
	name      string
	code      sql.NullString
	price     int64
	pathImage sql.NullString
}

func (b *Barang) Set(item *payload.Barang) {
	b.id = item.ID
	b.name = item.Name
	b.price = item.Price
	b.code.Scan(item.Code)
	b.pathImage.Scan(item.PathImage)
}

func (b *Barang) Get() payload.Barang {
	return payload.Barang{
		ID:        b.id,
		Name:      b.name,
		Code:      b.code.String,
		Price:     b.price,
		PathImage: b.pathImage.String,
	}
}

func (b *Barang) createID() {
	time := time.Now()
	year := strconv.Itoa(time.Year())
	month := strings.ToUpper(time.Month().String())
	nano := strconv.Itoa(time.Nanosecond())

	var productID []string
	productID = append(productID, "Product")
	productID = append(productID, year)
	productID = append(productID, month[0:3])
	productID = append(productID, nano)

	ID := strings.Join(productID, "-")
	encode := base64.StdEncoding.EncodeToString([]byte(ID))
	b.id = encode
}

func (b *Barang) Insert(db *sql.DB) (int64, error) {
	b.createID()
	query := fmt.Sprintf("INSERT INTO public.product(id, name, price) VALUES ('%s', '%s', %d)", b.id, b.name, b.price)
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
