package web

type ProductResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Picture  string `json:"picture"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	// User_Id  int    `json:"user_id"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}
