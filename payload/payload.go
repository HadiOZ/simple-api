package payload

type Barang struct {
	ID        string
	Name      string
	Price     int64
	Code      string
	PathImage string
}

type LogBarang struct {
	ID     string
	Name   Barang
	Action string
	Date   string
	Amount int64
	Admin  User
}

type User struct {
	ID       string
	Name     string
	Username string
	Password string
	Role     string
}
