package models

import (
	"database/sql"
	"fmt"
	"simple-api/payload"
)

type LogBarang struct {
	id     string
	name   Barang
	action string
	date   string
	amount int64
	admin  User
}

func (lb *LogBarang) Set(item *payload.LogBarang) {
	var barang Barang
	var admin User

	barang.Set(&item.Name)
	admin.Set(&item.Admin)

	lb.id = item.ID
	lb.date = item.Date
	lb.name = barang
	lb.admin = admin
	lb.action = item.Action
	lb.amount = item.Amount
}

func (lb *LogBarang) Get() payload.LogBarang {

	return payload.LogBarang{
		ID:     lb.id,
		Name:   lb.name.Get(),
		Action: lb.action,
		Date:   lb.date,
		Amount: lb.amount,
		Admin:  lb.admin.Get(),
	}
}

// func (lb *LogBarang) createID() {
// 	time := time.Now()
// 	year := strconv.Itoa(time.Year())
// 	month := strings.ToUpper(time.Month().String())
// 	nano := strconv.Itoa(time.Nanosecond())

// 	var productID []string
// 	productID = append(productID, "Log")
// 	productID = append(productID, year)
// 	productID = append(productID, month[0:3])
// 	productID = append(productID, nano)

// 	ID := strings.Join(productID, "-")
// 	encode := base64.StdEncoding.EncodeToString([]byte(ID))
// 	lb.id = encode
// }

func (lb *LogBarang) Insert(db *sql.DB) (int64, error) {
	// lb.createID()
	query := fmt.Sprintf(`INSERT INTO public.log_inventory(action, amount, id_product, admin) VALUES ('%s', %d, '%s', '%s');`, lb.action, lb.amount, lb.name.id, lb.admin.id)
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
