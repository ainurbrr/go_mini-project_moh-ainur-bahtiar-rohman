package payload

type CreateUserRequest struct {
	Name       string `json:"name" form:"name" validate:"required,max=20"`
	Occupation string `json:"occupation" form:"occupation" binding:"required"`
	Email      string `json:"email" form:"email" validate:"required,email"`
	Password   string `json:"password" form:"password" validate:"required,min=5"`
}

type CreateUserResponse struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}


