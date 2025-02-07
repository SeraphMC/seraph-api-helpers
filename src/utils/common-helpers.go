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
	grantsLocal := ctx.Locals("grants")
	if grantsLocal == nil {
		return false
	}

	if grants, ok := grantsLocal.(string); !ok {
		return false
	} else {
		if slices.Contains(strings.Split(strings.ToLower(grants), ","), strings.ToLower(permissionName)) {
			return true
		}
		return false
	}
}
