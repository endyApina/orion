package models

type Subscribe struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

//ResponseObject holds response object data
type ResponseObject struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}
