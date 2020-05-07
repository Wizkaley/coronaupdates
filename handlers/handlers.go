package handlers

import (
	"coronaupdates/models"
	"coronaupdates/utils"
	"encoding/json"
	"html/template"
	"net/http"
)

// Hello is a basic handler
func Hello(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/get.html"))
	if r.Method != http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	resp := utils.GetData()
	var cases []models.Cases
	err := json.Unmarshal(resp, &cases)
	utils.CheckError(err)
	// fmt.Print(cases)

	tmpl.Execute(w, cases)
}
