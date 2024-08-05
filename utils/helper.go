package utils

import (
	"regexp"

	"github.com/gofiber/fiber/v2"

)

func EmailRegex(email string) bool {
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return regexpEmail.MatchString(email)
}

func PhoneRegex(phoneNumber string) bool {
	regexPhone := regexp.MustCompile(`^\+[0-9]{1,3}[\s.-]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}([-\s\.]?[0-9]){4,6}$`)
	return regexPhone.MatchString(phoneNumber)
}

func GetLanguage(ctx *fiber.Ctx) string {
	lang := ctx.Get("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	return lang
}
