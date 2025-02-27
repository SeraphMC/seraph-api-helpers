package validation

// ErrorResponse represents the structure of an API error response, including success, error code, cause, additional error details, documentation URL, and response time in milliseconds.
type ErrorResponse struct {
	Success       bool          `json:"success"`
	Code          int           `json:"code"`
	Cause         string        `json:"cause"`
	Extra         []ErrorObject `json:"extra"`
	Documentation string        `json:"documentation"`
	MsTime        int64         `json:"msTime"`
}

// ErrorObject represents a structured error with a name, user-facing reason,
// developer-facing reason, and an associated error code.
type ErrorObject struct {
	Name            string `json:"name"`
	Reason          string `json:"reason"`
	DeveloperReason string `json:"developer_reason"`
	Code            int    `json:"code"`
}
