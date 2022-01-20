package models

import (
	"database/sql"
	"fmt"
	"simple-api/payload"

	"github.com/rs/xid"
)

type Product struct {
	id        string
	name      string
	code      sql.NullString
	price     int64
	pathImage sql.NullString
	stock     int64
}

func (b *Product) Set(item *payload.Product) {
	b.id = item.ID
	b.name = item.Name
	b.price = item.Price
	b.code.Scan(item.Code)
	b.pathImage.Scan(item.PathImage)
	b.stock = item.Stock

}

func (b *Product) Get() payload.Product {
	return payload.Product{
		ID:        b.id,
		Name:      b.name,
		Code:      b.code.String,
		Price:     b.price,
		PathImage: b.pathImage.String,
		Stock:     b.stock,
	}
}

func (b *Product) Insert(db *sql.DB) (int64, error) {
	id := xid.New().String()
	query := fmt.Sprintf(`INSERT INTO public.product(id_product, name, code, price) 
		VALUES ('%s', '%s', '%s', %d)`, id, b.name, b.code.String, b.price)
	fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	effect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	b.id = id
	return effect, nil
}

func SelectAllProduct(db *sql.DB) ([]payload.Product, error) {
	var result []payload.Product
	query := "SELECT id_product, name, price, code, stock, image_path FROM public.product;"
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		return result, err
	}

	for row.Next() {
		item := payload.Product{}
		if err := row.Scan(&item.ID, &item.Name, &item.Price, &item.Code, &item.Stock, &item.PathImage); err != nil {
			return result, err
		}
		result = append(result, item)
	}

	return result, nil
}

// func SelectByName()
