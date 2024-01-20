package api

import (
	"html/template"
	"net/http"

	"github.com/musarafik/newsy/services"
	"github.com/musarafik/newsy/structs"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	headlines, err := services.GetHeadlines()
	data := structs.HomeData{
		PageTitle: "Headlines",
		Headlines: headlines,
	}
	if err != nil {
		http.Error(w, "Error getting headlines", http.StatusInternalServerError)
		return
	}

	template, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
