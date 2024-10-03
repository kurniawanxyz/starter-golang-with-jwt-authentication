package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandleResponse(c *fiber.Ctx, status int, message string, data any) error {

	var success bool = false

	if status >= 200 && status <= 299 {
		success = true
	}

	return c.Status(status).JSON(fiber.Map{
		"success": success,
		"message": message,
		"data":    data,
	})
}

func HandleValidationMessage(c *fiber.Ctx,err error) error {

	errorMessages := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages[err.Field()] = GetValidationMessages(err.Tag(), err.Field(), err.Param())
	}
	return c.Status(422).JSON(fiber.Map{
		"success": false,
		"message": "Validation error",
		"data":    errorMessages,
	})
}

func GetValidationMessages(tag string, field string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Field %s harus diisi.", field)
	case "email":
		return fmt.Sprintf("Field %s harus berupa alamat email yang valid.", field)
	case "min":
		return fmt.Sprintf("Field %s harus memiliki minimal %s karakter.", field, param)
	case "max":
		return fmt.Sprintf("Field %s harus memiliki maksimal %s karakter.", field, param)
	case "len":
		return fmt.Sprintf("Field %s harus memiliki panjang %s karakter.", field, param)
	case "eqfield":
		return fmt.Sprintf("Field %s harus sama dengan field %s.", field, param)
	case "nefield":
		return fmt.Sprintf("Field %s tidak boleh sama dengan field %s.", field, param)
	case "lt":
		return fmt.Sprintf("Field %s harus kurang dari %s.", field, param)
	case "lte":
		return fmt.Sprintf("Field %s harus kurang dari atau sama dengan %s.", field, param)
	case "gt":
		return fmt.Sprintf("Field %s harus lebih dari %s.", field, param)
	case "gte":
		return fmt.Sprintf("Field %s harus lebih dari atau sama dengan %s.", field, param)
	case "oneof":
		return fmt.Sprintf("Field %s harus salah satu dari %s.", field, param)
	case "numeric":
		return fmt.Sprintf("Field %s harus berupa angka.", field)
	case "alphanum":
		return fmt.Sprintf("Field %s harus berupa karakter alfanumerik.", field)
	case "url":
		return fmt.Sprintf("Field %s harus berupa URL yang valid.", field)
	case "uuid":
		return fmt.Sprintf("Field %s harus berupa UUID yang valid.", field)
	case "uuid4":
		return fmt.Sprintf("Field %s harus berupa UUID4 yang valid.", field)
	case "uuid5":
		return fmt.Sprintf("Field %s harus berupa UUID5 yang valid.", field)
	case "ip":
		return fmt.Sprintf("Field %s harus berupa alamat IP yang valid.", field)
	case "ipv4":
		return fmt.Sprintf("Field %s harus berupa alamat IPv4 yang valid.", field)
	case "ipv6":
		return fmt.Sprintf("Field %s harus berupa alamat IPv6 yang valid.", field)
	case "mac":
		return fmt.Sprintf("Field %s harus berupa alamat MAC yang valid.", field)
	case "e164":
		return fmt.Sprintf("Field %s harus berupa nomor telepon yang valid eg: +62XXX", field)
	// Tambahkan case lainnya sesuai kebutuhan
	default:
		return fmt.Sprintf("Field %s memiliki kesalahan: %s.", field, tag)
	}
}