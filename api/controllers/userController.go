package controllers

import (
	"errors"

	"github.com/Kamva/mgm/v2"
	"github.com/YonchevSimeon/fiber-mongodb-vue/models"
	"github.com/YonchevSimeon/fiber-mongodb-vue/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {
	params := new(struct {
		Username string
		Password string
	})

	ctx.BodyParser(&params)

	user := &models.User{}

	collection := mgm.Coll(user)

	usernameErr := collection.FindOne(ctx.Context(), bson.M{"username": params.Username}).Decode(&user)

	match := checkPasswordHash(params.Password, user.Password)

	if !match || usernameErr != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Wrong username or password, please try again!",
		})
	}

	token, err := utils.GenerateJwtToken()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok":       true,
		"token":    token,
		"username": user.Username,
		"userId":   user.ID,
	})

}

func Register(ctx *fiber.Ctx) error {
	params := new(struct {
		Email    string
		Username string
		Password string
	})

	ctx.BodyParser(&params)

	emailErr := isValidEmail(params.Email)
	usernameErr := isValidUsername(params.Username)
	passwordErr := isValidPassword(params.Password)

	if emailErr != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": emailErr.Error(),
		})
	} else if usernameErr != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": usernameErr.Error(),
		})
	} else if passwordErr != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": passwordErr.Error(),
		})
	}

	hashPassword, _ := hashPassword(params.Password)

	user := models.Register(params.Email, params.Username, hashPassword)

	err := mgm.Coll(user).Create(user)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"ok": true,
	})

}

func isValidEmail(email string) error {
	// TODO: check if email is valid // may be regex
	if len(email) == 0 {
		return errors.New("please enter a valid email")
	}
	return nil
}

func isValidUsername(username string) error {
	// TODO: check if there is an user with this username

	return nil
}

func isValidPassword(password string) error {
	// TODO: check if it contains numbers
	// TODO: check if it contains symbols

	passwordLength := len(password)
	if passwordLength < 8 || passwordLength > 16 {
		return errors.New("the password should be between 8 and 16 characters long")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
