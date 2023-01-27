package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
