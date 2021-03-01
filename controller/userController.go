package controller

import (
	"encoding/json"
	"log"
	"mekar/middleware"
	"mekar/model"
	"mekar/responses"
	"mekar/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func UserController(r *mux.Router, user usecases.UserUsecase) {
	UserHandler := UserHandler{user}

	// Global Interceptor (Middleware)
	r.Use(middleware.ActivityLogMiddleware)

	us := r.PathPrefix("").Subrouter()
	us.HandleFunc("/users", UserHandler.UsersList).Methods(http.MethodGet)
	us.HandleFunc("/user/{NoKTP}", UserHandler.GetUserByID).Methods(http.MethodGet)
	us.HandleFunc("/user", UserHandler.CreateNewUser).Methods(http.MethodPost)
	us.HandleFunc("/user/{NoKTP}", UserHandler.UpdateUser).Methods(http.MethodPut)
	us.HandleFunc("/user/{NoKTP}", UserHandler.DeleteUser).Methods(http.MethodDelete)
}

func (s UserHandler) UsersList(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserUsecase.GetAllUser()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		var userResponse responses.MyRespon = responses.MyRespon{"Getting User Data Successfully!", users}
		byteOfUsers, err := json.Marshal(userResponse)
		if err != nil {
			w.Write([]byte("Something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfUsers)
	}
}

func (s UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	NoKTP := param["NoKTP"]
	User, err := s.UserUsecase.GetUserByID(NoKTP)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ServiceResponse.Message = "User Code Not Found!"
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Write([]byte(ServiceResponse.Message))
	// 	}
	// } else {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Not Found"))
	} else {
		var userResponse responses.MyRespon = responses.MyRespon{"Well Done!", User}
		byteOfUsers, err := json.Marshal(userResponse)
		if err != nil {
			w.Write([]byte("Something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfUsers)
	}
}

func (s UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var User model.User
	_ = json.NewDecoder(r.Body).Decode(&User)
	err := s.UserUsecase.CreateNewUser(User)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Can't Create New User Data!"))
	} else {
		log.Println("Create New User Data Successfully!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Create New User Data Successfully!"))
	}
}

func (s UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user *model.User
	NoKTP := (vars["NoKTP"])
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := s.UserUsecase.UpdateUser(NoKTP, user)
	if err != nil {
		w.Write([]byte("Oops something went wrong!"))
	} else {
		log.Println("Update User Data Successfully!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Update User Data Successfully!"))
	}
}

func (s UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	err := s.UserUsecase.DeleteUser(param["NoKTP"])
	if err != nil {
		w.Write([]byte("Can't Delete User Data!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delete User Data Successfully!"))
	}
}

// Convert integer to string
// key := "id"
// id := utils.GetPathVar(key, r)
// idStudent, _ := strconv.Atoi(id)
