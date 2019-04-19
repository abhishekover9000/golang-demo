package controllers

import (
	"fmt"
	"net/http"

	"../utils"
)

type Controller struct{}

func (c Controller) ProtectedEndpoint() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Protected Endpoint invoked")
		utils.RespondWithJSON(w, "Yes")
	}

}
