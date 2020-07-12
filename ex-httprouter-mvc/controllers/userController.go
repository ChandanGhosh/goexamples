package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chandanghosh/goexercises/ex-httprouter-mvc/models"
	"github.com/julienschmidt/httprouter"
	uid "github.com/satori/go.uuid"
)

var users = map[string]models.User{}

// UserController ...
type UserController struct {
}

// NewUserController ...
func NewUserController() *UserController {
	return &UserController{}
}

// GetUsers ...
func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err.Error())
	}
}

//GetUser ...
func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if usr, ok := users[id]; ok {
		_ = json.NewEncoder(w).Encode(&usr)
		return
	}
}

// UpdateUser ...
func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	users[user.Id] = user
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	}
}

// DeleteUser ...
func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if _, ok := users[id]; ok {
		delete(users, id)
	}
	fmt.Fprintln(w, "User deleted")
}

// CreateUser ...
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usr := models.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	}
	sid := uid.NewV4()
	usr.Id = sid.String()
	users[usr.Id] = usr
	mj, _ := json.Marshal(usr)
	_, _ = fmt.Fprintf(w, "%s\n", mj)
}
