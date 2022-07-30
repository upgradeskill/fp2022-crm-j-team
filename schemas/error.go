package schemas

type DatabaseError struct {
	Type string
	Code int
}

type ErrorResponse struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
