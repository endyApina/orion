package controller

import (
	"encoding/json"
	"errors"
	"orion/models"
	"net/http"
)

//Registration godoc
//@Summary Handle unique User Registration
//@Description Accept JSON data of User objects and returns valid response
//@Accept json
//@Tags Authentication
//@produce json
//@Param   UserData      body models.RegistrationData true  "The User Registration data"
//@Success 200 {object} models.RegistrationData	"ok"
//@Failure 400 {object} models.ResponseObject "Check Response Message"
//@Router /subscribe [post]
func Subscribe(w http.ResponseWriter, r *http.Request) {
	var subscribe models.Subscribe

	err := decodeJSONBody(w, r, &subscribe)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			models.LogError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&mr)
		} else {
			models.LogError(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = subscribe.RegisterNewSubscriber()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ValidResponse(400, subscribe, "error"))
	}

	go subscribe.NotifyNewSubscriber()

	// userRegistration := models.HandleUserRegistration()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ValidResponse(200, subscribe, "success"))
}
