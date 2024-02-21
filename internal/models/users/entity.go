package users

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

type UserCreateRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type Response struct {
	StatusCode int         `json:"status_code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     string      `json:"errors,omitempty"`
}

type UserResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

func pwdCompare(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	confirmField := fl.Parent().FieldByName("Password").String()
	return field == confirmField
}
