package web

type ProductCreateRequest struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    int    `json:"phone"`
}
