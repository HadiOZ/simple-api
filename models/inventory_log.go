package models

import (
	"database/sql"
	"fmt"
	"simple-api/payload"

	"github.com/rs/xid"
)

type Log struct {
	id     string
	name   Product
	action string
	date   string
	amount int64
	admin  Employee
}

func (lb *Log) Set(item *payload.Log) {
	var name Product
	var admin Employee

	name.Set(&item.Name)
	admin.Set(&item.Admin)

	lb.id = item.ID
	lb.date = item.Date
	lb.name = name
	lb.admin = admin
	lb.action = item.Action
	lb.amount = item.Amount
}

func (lb *Log) Get() payload.Log {

	return payload.Log{
		ID:     lb.id,
		Name:   lb.name.Get(),
		Action: lb.action,
		Date:   lb.date,
		Amount: lb.amount,
		Admin:  lb.admin.Get(),
	}
}

// func (lb  Log) createID() {
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

func (lb *Log) Insert(db *sql.DB) (int64, error) {
	id := xid.New().String()
	lb.id = id
	query := fmt.Sprintf(`INSERT INTO public.inventory_log(id_log, id_product_fkey, id_employee_fkey, action, amount) VALUES ('%s', '%s', '%s', '%s', %d);`, id, lb.name.id, lb.admin.id, lb.action, lb.amount)
	fmt.Print(query)
	res, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	effect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	lb.id = id
	return effect, nil
}

// func SelectByProduct()
