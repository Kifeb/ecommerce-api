package web

type ProductCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Picture  string `json:"picture" validate:"required,min=1,max=100"`
	Price    int    `json:"price" validate:"required,min=1,max=100"`
	Category string `json:"category" validate:"required,min=1,max=100"`
	Quantity int    `json:"quantity" validate:"required,min=1,max=100"`
	User_Id  int    `json:"user_id" validate:"required,min=1,max=100"`
}

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
	Role     string `json:"role" validate:"required,min=1,max=100"`
	Phone    string `json:"phone" validate:"required,min=1,max=100"`
}
