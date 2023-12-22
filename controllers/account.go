package controllers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"time"

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

func UpdateBalanceAccount(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Locals("userId").(int32)

	accountId, err := strconv.Atoi(c.Params("account_id"))
	if err != nil {
		return fmt.Errorf("Wrong account_id value")
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error with the file")
	}

	open, err := file.Open()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error with the file")
	}

	csvReader := csv.NewReader(open)

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range data {
		if i > 0 { // omit header line
			date, err := time.Parse("2006-01-02", line[1])
			if err != nil {
				continue
			}
			amount, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				continue
			}
			reason := line[3]

			//transaction block
			err = db.MakeTx(context.Background(), func() error {
				createTransferParams := sqlc_code.CreateTransferParams{
					Amount:    amount,
					Reason:    pgtype.Text{String: reason, Valid: true},
					AccountID: pgtype.Int4{Int32: int32(accountId), Valid: true},
					CreatedAt: pgtype.Timestamp{Time: date, Valid: true},
				}
				_, err := db.Queries.CreateTransfer(ctx, createTransferParams)
				if err != nil {
					fmt.Println("err CreateTransfer: ", err)
					return err
				}

				account, err := db.Queries.GetAccountForUpdate(context.Background(), int32(accountId))
				if err != nil {
					log.Println(err)
					return err
				}

				if account.UserID.Int32 != userId {
					return fmt.Errorf("User not authorized to access this account")
				}

				updateAccountParmas := sqlc_code.UpdateAccountParams{
					ID:      int32(accountId),
					Balance: pgtype.Float8{Float64: account.Balance.Float64 + amount, Valid: true},
				}

				_, err = db.Queries.UpdateAccount(ctx, updateAccountParmas)
				if err != nil {
					fmt.Println("err UpdateAccount: ", err)
					return err
				}
				return nil
			})
			fmt.Println("err MakeTx: ", err)

		}
	}

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"ok": "created"})

}
