package models

import (
	"log"
	"os"
)

//LogError logs all error to file
func LogError(err error) {
	f, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	logger := log.New(f, "Error: ", log.LstdFlags)
	logger.Println(err.Error())
}

//ValidResponse structures valid api response into a json object
func ValidResponse(code int, body interface{}, message string) ResponseObject {
	var response ResponseObject
	response.Code = code
	response.Message = message
	response.Body = body

	return response
}
