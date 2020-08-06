package app

import (
	"fmt"
	"github.com/sswapnil2/golang-microservices/mvc/controllers"
	"net/http"
)

const PORT string = ":8000"

func StartApp() {

	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}

	fmt.Println("Started application at 127.0.0.1: ", PORT)
}