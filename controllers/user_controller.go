package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/psinthorn/F2Go/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("User must be a number"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
