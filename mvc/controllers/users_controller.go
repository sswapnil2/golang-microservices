package controllers

import (
	"encoding/json"
	"github.com/sswapnil2/golang-microservices/mvc/services"
	"github.com/sswapnil2/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}
	// return the user to client
	jsonVal, _ := json.Marshal(user)

	res.Write(jsonVal)
}
