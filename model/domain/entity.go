package domain

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Picture  string `json:"picture"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	User_Id  int    `json:"user_id"`
}
