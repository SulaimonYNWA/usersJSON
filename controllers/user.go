package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	user_model "github.com/dovudwkt/playground/mvc_server/models/user"
	reply "github.com/dovudwkt/playground/mvc_server/pkg/http_reply"
)

type UserController struct {
	UserModel user_model.Model
}

func (h UserController) AddUser(w http.ResponseWriter, req *http.Request) {
	re := reply.NewHTTPReplier(reply.Options{
		ResponseWriter: w,
		ContentType:    "application/json",
		Accept:         []string{"application/json"},
	})
	setupResponse(&w, req)

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		re.Error(err.Error(), http.StatusBadRequest, false)
		return
	}

	var user user_model.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		re.Error("json unmarshall: "+err.Error(), http.StatusBadRequest, true)
		return
	}

	id, err := h.UserModel.AddUser(user)
	if err != nil {
		err := errors.New("add user controller: " + err.Error())
		re.Error(err.Error(), http.StatusInternalServerError, true)
		return
	}

	re.Reply(id, http.StatusCreated, false)
}

// GetAll - handler to get all users
func (h UserController) GetAll(w http.ResponseWriter, req *http.Request) {
	re := reply.NewHTTPReplier(reply.Options{
		ResponseWriter: w,
		ContentType:    "application/json",
		Accept:         []string{"application/json"},
	})
	setupResponse(&w, req)

	users, err := h.UserModel.GetAll()
	if err != nil {
		err := errors.New("get users controller: " + err.Error())
		re.Error(err.Error(), http.StatusInternalServerError, true)
		return
	}

	re.Reply(users, http.StatusOK, false)
}

// GetByID - handler to get a user by id. If user not found, 404 StatusNotFound returned.
func (h UserController) GetByID(w http.ResponseWriter, req *http.Request) {
	re := reply.NewHTTPReplier(reply.Options{
		ResponseWriter: w,
		ContentType:    "application/json",
		Accept:         []string{"application/json"},
	})
	setupResponse(&w, req)

	// get id url parameter
	userID, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)
	if err != nil {
		re.Error(err.Error(), http.StatusInternalServerError, true)
		return
	}

	// get user by id
	user, err := h.UserModel.Get(userID)
	if err == sql.ErrNoRows { // user by id not exist
		re.Reply(http.StatusText(http.StatusNotFound), http.StatusNotFound, false)
		return
	}
	if err != nil {
		err := errors.New("get user by id controller: " + err.Error())
		re.Error(err.Error(), http.StatusInternalServerError, true)
		return
	}

	re.Reply(*user, http.StatusOK, false)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
