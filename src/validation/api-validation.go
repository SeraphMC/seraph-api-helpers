package validation

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// ErrorHandlerType constants representing common error messages.
type ErrorHandlerType = string

const (
	// InvalidApiKey indicates an invalid API key error.
	InvalidApiKey ErrorHandlerType = "Invalid API Key"
	// InvalidUuid indicates an invalid UUID error.
	InvalidUuid ErrorHandlerType = "Invalid UUID"
	// InvalidToken indicates an invalid token error.
	InvalidToken ErrorHandlerType = "Invalid Token"
	// NotEnoughPermission indicates a permission-related error.
	NotEnoughPermission ErrorHandlerType = "Not Enough Permission"
)

// HandleError generates and sends a JSON error response based on the given status, cause, and error details.
// It handles specific error cases like unknown routes, invalid UUIDs, or API key issues.
// The response contains metadata like request ID and documentation URL for debugging.
func HandleError(ctx *fiber.Ctx, status int, cause string, errors []ErrorObject) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if strings.Contains(cause, "Cannot GET") {
		errors = append(errors, ErrorObject{
			Name:            "Unknown route",
			Reason:          "This is not a known route.",
			DeveloperReason: cause,
			Code:            404,
		})
	} else if strings.Contains(cause, "ErrHandler: json:") || strings.Contains(cause, "ERROR: ") {
		errors = append(errors, ErrorObject{
			Name:            "Internal API Error",
			Reason:          "This type is different to the type provided by our database. Report this to Seraph with the developer reason and path located here: " + ctx.OriginalURL() + ".",
			DeveloperReason: cause,
			Code:            500,
		})
	} else if strings.Contains(cause, InvalidUuid) {
		errors = append(errors, ErrorObject{
			Name:            InvalidUuid,
			Reason:          "Please enter a valid UUID v4 dashed or undashed",
			DeveloperReason: "The UUID provided was not valid. Ensure you are using a v4 UUID.",
			Code:            400,
		})
	} else if strings.Contains(cause, InvalidApiKey) {
		errors = append(errors, ErrorObject{
			Name:            InvalidApiKey,
			Reason:          "The API Key has been locked or is invalid.",
			DeveloperReason: "Check your API Key is valid and is in the header.",
			Code:            401,
		})
	} else {
		errors = append(errors, ErrorObject{
			Name:            "Bad Request",
			Reason:          "Please retry the request or look at the docs!",
			DeveloperReason: cause,
			Code:            400,
		})
	}

	requestId := utils.CopyString(ctx.GetRespHeader("x-seraph-request"))
	errors = append(errors, ErrorObject{
		Name:            "Request-ID",
		Reason:          requestId,
		DeveloperReason: "The request ID to report to Seraph, This will not be stored if the request is a public request",
		Code:            200,
	})

	errorResponse := ErrorResponse{
		Success:       false,
		Code:          status,
		Cause:         cause,
		Extra:         errors,
		Documentation: ctx.BaseURL(),
		MsTime:        time.Now().UnixMilli(),
	}

	return ctx.Status(status).JSON(errorResponse)
}

// ParseBody parses the HTTP request body into a generic type T and returns the result or an error if parsing fails.
func ParseBody[T any](ctx *fiber.Ctx) (T, error) {
	var bodyPost T
	if err := ctx.BodyParser(&bodyPost); err != nil {
		return bodyPost, err
	}
	return bodyPost, nil
}

// ToTimePointer converts a time.Time value to a pointer to that value.
func ToTimePointer(t time.Time) *time.Time {
	return &t
}
