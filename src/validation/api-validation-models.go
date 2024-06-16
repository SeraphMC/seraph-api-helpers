package validation

type ErrorResponse struct {
	Success       bool          `json:"success"`
	Code          int           `json:"code"`
	Cause         string        `json:"cause"`
	Extra         []ErrorObject `json:"extra"`
	Documentation string        `json:"documentation"`
	MsTime        int64         `json:"msTime"`
}

type ErrorObject struct {
	Name            string `json:"name"`
	Reason          string `json:"reason"`
	DeveloperReason string `json:"developer_reason"`
	Code            int    `json:"code"`
}
