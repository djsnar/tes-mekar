package config

import (
	"database/sql"
	"fmt"
	"log"
	"mekar/controller"
	"mekar/repositories"
	"mekar/usecases"
	"mekar/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// func CreateRouter
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

// func RunServer
func RunServer(router *mux.Router) {
	port := utils.ReadEnv("RouterPort", "3030")
	log.Printf("Starting Web Server at %v\n", port)
	err := http.ListenAndServe("localhost: "+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

// func InitRouter
func InitRouter(r *mux.Router, db *sql.DB) {
	userRepository := repositories.InitUserRepoImpl(db)
	userUsecase := usecases.InitUserUseCase(userRepository)
	controller.UserController(r, userUsecase)

	r.HandleFunc("/home", home).Methods(http.MethodGet)
}

// func Home
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Write([]byte("Data Not Found"))
}
