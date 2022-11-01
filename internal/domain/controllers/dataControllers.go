package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/victorbej/hack-auth/internal/domain/models"
	"github.com/victorbej/hack-auth/internal/domain/utils"
)

var CreateUserTable = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	ut := &models.UserTable{}

	err := json.NewDecoder(r.Body).Decode(ut)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	ut.UserId = user
	resp := ut.Create()
	utils.Respond(w, resp)
}

var GetUserTablesFor = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.Context().Value("user").(uint)
	data := models.GetUserTables(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
