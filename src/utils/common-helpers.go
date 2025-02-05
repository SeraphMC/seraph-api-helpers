package utils

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"slices"
	"strings"
)

func IsDevelopmentMode() bool {
	devMode := os.Getenv("APP_ENV") == "DEV"
	return devMode
}

func CheckPermissions(ctx *fiber.Ctx, permissionName string) bool {
	grants := ctx.Locals("grants").(string)

	if slices.Contains(strings.Split(strings.ToLower(grants), ","), strings.ToLower(permissionName)) {
		return true
	}
	return false
}
