package web

type ProductCreateRequest struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	User_Id  int    `json:"user_id"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
}
