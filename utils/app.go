package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func GetApp() fiber.Map {
	return fiber.Map{
		"name":   "Rent-N-Go Backend",
		"slogan": "Your journey, our priority",
	}
}

func ShouldPanic(err error) {
	if IsProduction() {
		log.Fatalf("An error occurred: %v\n", err.Error())
	} else {
		panic(err)
	}
}

func RecordLog(err error) {
	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err.Error())
	}
}

func IsProduction() bool {
	return viper.GetString("APP_ENV") == "production"
}

func SafeThrow(w *fiber.Ctx, err error) error {
	errorMessage := "Can't proceed your request"

	if !IsProduction() {
		errorMessage = err.Error()
	}

	return w.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Something went wrong",
		"error":   errorMessage,
	})
}
