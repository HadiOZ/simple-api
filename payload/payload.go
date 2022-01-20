package payload

type Product struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Price     int64  `json:"price" bson:"price"`
	Code      string `json:"code" bson:"code"`
	PathImage string `json:"path-image" bson:"path-image"`
	Stock     int64  `json:"stock" bson:"stock"`
}

type Log struct {
	ID      string   `json:"id" bson:"id"`
	Product Product  `json:"product"  bson:"product"`
	Action  string   `json:"action" bson:"action"`
	Date    string   `json:"date" bson:"date"`
	Amount  int64    `json:"amount" bson:"amount"`
	Admin   Employee `json:"admin" bson:"admin"`
}

type Employee struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Position string `json:"position" bson:"position"`
}
