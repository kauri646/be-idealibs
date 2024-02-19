package handlers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/be-idealibs/config"
	"github.com/kauri646/be-idealibs/internal/models/users"
	"github.com/kauri646/be-idealibs/utils"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []users.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)

}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(users.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := users.User{
		Username: user.Username,
		Email:    user.Email,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "iternal server error",
		})

	}

	newUser.Password = hashedPassword

	errCreateUser := config.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user users.User
	err := config.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	Name: user.Name,
	// 	Email: user.Email,
	//     Address: user.Address,
	//     Phone: user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {

	userRequest := new(users.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user users.User

	userId := ctx.Params("id")
	err := config.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userRequest.Username != "" {
		user.Username = userRequest.Username
	}

	errUpdate := config.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {

	userRequest := new(users.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user users.User
	var isEmailUserExist users.User

	userId := ctx.Params("id")
	err := config.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errCheckEmail := config.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail != nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "user already used.",
		})
	}

	user.Email = userRequest.Email

	errUpdate := config.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	var user users.User

	err := config.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	errDelete := config.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}

func LoginHandler(ctx *fiber.Ctx) error {
	LoginRequest := new(users.LoginRequest)

	if err := ctx.BodyParser(LoginRequest); err != nil {
		return err
	}
	log.Println(LoginRequest)

	validate := validator.New()
	errValidate := validate.Struct(LoginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user users.User
	err := config.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	claims := jwt.MapClaims{}
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "atra@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

// func RegisterHandler(ctx *fiber.Ctx) error {
// 	user := new(users.RegisterRequest)

// 	if err := ctx.BodyParser(user); err != nil {
// 		return err
// 	}

// 	validate := validator.New()
// 	errValidate := validate.Struct(user)
// 	if errValidate != nil {
// 		return ctx.Status(400).JSON(fiber.Map{
// 			"message": "failed",
// 			"error":   errValidate.Error(),
// 		})
// 	}

// 	newUser := users.User{
// 		Username: user.Username,
// 		Email:    user.Email,
// 	}

// 	hashedPassword, err := utils.HashingPassword(user.Password)
// 	if err != nil {
// 		log.Println(err)
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "iternal server error",
// 		})

// 	}

// 	newUser.Password = hashedPassword

// 	errCreateUser := config.DB.Create(&newUser).Error
// 	if errCreateUser != nil {
// 		return ctx.Status(500).JSON(fiber.Map{
// 			"message": "failed to store data",
// 		})
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"message": "success",
// 		"data":    newUser,
// 	})
// }

func RegisterHandler(ctx *fiber.Ctx) error {
	user := new(users.RegisterRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Manual validation to check password match
	if user.Password != user.ConfirmPassword {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   "Passwords do not match",
		})
	}

	newUser := users.User{
		Username: user.Username,
		Email:    user.Email,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := config.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
