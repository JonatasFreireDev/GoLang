package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var p models.Personalidade

	if err != nil {
		fmt.Fprint(w, "Ocorreu um Erro!")
	}

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

	// for _, personalidade := range models.Personalidades {
	// 	if personalidade.Id == id {
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }
}

func PostPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var p models.Personalidade

	if err != nil {
		fmt.Fprint(w, "Ocorreu um Erro!")
	}

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func PutPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var p models.Personalidade

	if err != nil {
		fmt.Fprint(w, "Ocorreu um Erro!")
	}

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)

}
