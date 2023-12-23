package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/db"
	"github.com/tecolotedev/stori_back/db/sqlc_code"
	"github.com/tecolotedev/stori_back/utils"
)

type loginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

type loginResponse struct {
	ID        int    `json:"id" form:"id"`
	Email     string `json:"email" form:"email"`
	Username  string `json:"username"  form:"username"`
	CreatedAt string `json:"created_at"  form:"created_at"`
}

func Login(c *fiber.Ctx) error {
	loginBody := new(loginRequest)

	if err := c.BodyParser(loginBody); err != nil {
		fmt.Println(err)
		return fmt.Errorf("bad entities")
	}

	user, err := db.Queries.GetUser(context.Background(), loginBody.Email)
	if err != nil {
		return err
	}

	err = utils.CheckPassword(loginBody.Password, user.Password)
	if err != nil {
		return err
	}

	userResponse := loginResponse{
		ID:        int(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Time.String(),
	}

	token, err := utils.CreateToken(user.ID, time.Hour)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.JSON(userResponse)

}

func VerifyToken(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	payload, err := utils.VerifyToken(accessToken)
	if err != nil {
		return err
	}
	return c.JSON(payload)
}

type createUserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	newUserBody := new(createUserRequest)

	if err := c.BodyParser(newUserBody); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newUserBody.Password)
	if err != nil {
		return err
	}

	params := sqlc_code.CreateUserParams{
		Username: newUserBody.Username,
		Email:    newUserBody.Email,
		Password: hashedPassword,
	}

	userCreated, err := db.Queries.CreateUser(context.Background(), params)
	if err != nil {

		return err
	}

	return c.JSON(userCreated)
}
