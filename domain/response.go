package domain

type CommonResponse struct {
	Status  int         `json:"status"`          // Status Code (ex. 200, 400, 500, etc...)
	Error   string      `json:"error,omitempty"` // Error Code (ex. "Invalid Request", "Parsing Error", etc...')
	Message string      `json:"message"`         // Response Message
	Data    interface{} `json:"data,omitempty"`  // Response Data
	Path    string      `json:"path,omitempty"`  // Request Path
}
