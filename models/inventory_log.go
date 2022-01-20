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

	name.Set(&item.Product)
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
		ID:      lb.id,
		Product: lb.name.Get(),
		Action:  lb.action,
		Date:    lb.date,
		Amount:  lb.amount,
		Admin:   lb.admin.Get(),
	}
}

func (lb *Log) Insert(db *sql.DB) (int64, error) {
	id := xid.New().String()
	lb.id = id
	query := fmt.Sprintf(`INSERT INTO public.inventory_log(id_log, id_product_fkey, id_employee_fkey, action, amount) VALUES ('%s', '%s', '%s', '%s', %d);`, id, lb.name.id, lb.admin.id, lb.action, lb.amount)
	fmt.Println(query)
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

func (lb *Log) Update(db *sql.DB) (int64, error) {
	query := fmt.Sprintf(`UPDATE public.inventory_log SET amount = %d, id_employee_fkey = '%s' WHERE id_log ='%s';`, lb.amount, lb.admin.id, lb.id)
	fmt.Println(query)
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

func SelectLogByIDProduct(id string, db *sql.DB) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	query := fmt.Sprintf(`SELECT id_log, admin, action, date, amount FROM view_log WHERE id_product = '%s';`, id)
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := payload.Log{}
		if err := row.Scan(&item.ID, &item.Admin.Name, &item.Action, &item.Date, &item.Amount); err != nil {
			return result, err
		}
		result = append(result, map[string]interface{}{"id": item.ID, "admin": item.Admin.Name, "action": item.Action, "date": item.Date, "amount": item.Amount})
	}
	return result, nil
}

func SelectLogHistory(id string, db *sql.DB) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var time string
	query := fmt.Sprintf(`SELECT * FROM view_history_log WHERE id_log = '%s';`, id)
	fmt.Println(query)
	row, err := db.Query(query)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := payload.Log{}
		if err := row.Scan(&item.ID, &item.Admin.Name, &item.Date, &time, &item.Amount); err != nil {
			return result, err
		}
		result = append(result, map[string]interface{}{"id": item.ID, "admin": item.Admin.Name, "date": item.Date, "time": time, "amount": item.Amount})
	}
	return result, nil
}
