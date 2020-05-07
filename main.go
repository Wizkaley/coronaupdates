package main

import (
	"net/http"
	"weather/handlers"
)

func main() {
	// resp := utils.GetData()
	// var cases []models.Cases
	// err := json.Unmarshal(resp, &cases)
	// utils.CheckError(err)
	// fmt.Print(cases)

	http.HandleFunc("/", handlers.Hello)
	http.ListenAndServe(":1234", nil)
}
