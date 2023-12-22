package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tecolotedev/stori_back/db"
	"github.com/tecolotedev/stori_back/db/sqlc_code"
)

func ListAccounts(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return fmt.Errorf("Wrong limit value")
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return fmt.Errorf("Wrong offset value")
	}

	params := sqlc_code.ListAccountsParams{
		UserID: pgtype.Int4{Int32: userId, Valid: true},
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	accounts, err := db.Queries.ListAccounts(context.Background(), params)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(accounts)

}

func GetAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	accountId, err := strconv.Atoi(c.Params("account_id"))
	if err != nil {
		return fmt.Errorf("Wrong account_id value")
	}

	account, err := db.Queries.GetAccount(context.Background(), int32(accountId))
	if err != nil {
		log.Println(err)
		return err
	}

	if account.UserID.Int32 != userId {
		return fmt.Errorf("User not authorized to access this account")
	}

	return c.JSON(account)

}

type createAccountRequest struct {
	Balance  float64 `json:"balance" form:"balance"`
	Currency string  `json:"currency" form:"currency"`
}

func CreateAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	createAccountBody := new(createAccountRequest)

	if err := c.BodyParser(createAccountBody); err != nil {
		log.Println(err)
		return err
	}

	params := sqlc_code.CreateAccountParams{
		Balance:  pgtype.Float8{Float64: createAccountBody.Balance, Valid: true},
		Currency: createAccountBody.Currency,
		UserID:   pgtype.Int4{Int32: userId, Valid: true},
	}

	accountCreated, err := db.Queries.CreateAccount(context.Background(), params)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(accountCreated)

}
