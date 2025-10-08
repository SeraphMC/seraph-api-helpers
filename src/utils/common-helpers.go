package utils

import (
	"os"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// IsDevelopmentMode checks if the application is running in development mode by validating if the "APP_ENV" environment variable is set to "DEV". It returns true if the mode is "DEV", otherwise false.
func IsDevelopmentMode() bool {
	devMode := os.Getenv("APP_ENV") == "DEV"
	return devMode
}

// PermissionNode represents a type for defining permission identifiers.
type PermissionNode string

// CheckPermissions verifies if a given permission is included in the user's grants stored in the request context.
// It checks the "grants" local variable, ensuring it's a string, and compares it case-insensitively against the requested permission.
// Returns true if the permission exists, otherwise false.
func CheckPermissions(ctx *fiber.Ctx, permissionName PermissionNode) bool {
	grantsLocal := ctx.Locals("grants")
	if grantsLocal == nil {
		return false
	}

	if grants, ok := grantsLocal.(string); !ok {
		return false
	} else {
		if slices.Contains(strings.Split(strings.ToLower(grants), ","), permissionName.GetName()) {
			return true
		}
		return false
	}
}
