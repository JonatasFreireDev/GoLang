package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.PostPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.PutPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
