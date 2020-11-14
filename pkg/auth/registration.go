package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Redict/imgstore/pkg/server"
)

// UIRegistrationForm is used to process registration form data from user
type UIRegistrationForm struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatpassword"`
}

// RegistrationHandler function
func RegistrationHandler(rw http.ResponseWriter, r *http.Request) {
	if !server.CheckPOST(rw, r) {
		return
	}

	var s UIRegistrationForm

	err := server.DecodeJSON(rw, r, &s)
	if err != nil {
		var mr *server.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(rw, mr.Msg, mr.Status)
		} else {
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		log.Println(err.Error())
		return
	}
	fmt.Fprintf(rw, "UIRegistrationForm: %+v\n", s)
}
