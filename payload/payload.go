package payload

type Product struct {
	ID        string
	Name      string
	Price     int64
	Code      string
	PathImage string
	Stock     int64
}

type Log struct {
	ID     string
	Name   Product
	Action string
	Date   string
	Amount int64
	Admin  Employee
}

type Employee struct {
	ID       string
	Name     string
	Username string
	Password string
	Position string
}
