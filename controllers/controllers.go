package controllers

import "github.com/gofiber/fiber/v2"

type ControllerError struct {
	code    uint
	message string
}

var BodyError = ControllerError{
	code:    fiber.StatusBadRequest,
	message: "Error parsing the body content",
}

var ParamsError = ControllerError{
	code:    fiber.StatusBadRequest,
	message: "Error parsing params or query values",
}

var DatabaseError = ControllerError{
	code:    fiber.StatusInternalServerError,
	message: "Something went wrong, please try it later",
}
