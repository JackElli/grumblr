package user

import (
	"encoding/json"
	"fmt"
	"grumblrapi/main/responder"
	"grumblrapi/main/user"
	"grumblrapi/main/userstore"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	NEW_USER = "/user"
	GET_USER = "/user/{userId}"
)

type UserMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	UserStore userstore.UserStorer
}

func NewNewUserMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, userStore userstore.UserStorer) *UserMgr {
	e := &UserMgr{
		Logger:    logger,
		Router:    router,
		Responder: responder,
		UserStore: userStore,
	}
	e.Register()
	return e
}

// NewUser inserts a user into the database
func (mgr *UserMgr) NewUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		//TODO add validation here
		var userDetails struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		json.NewDecoder(req.Body).Decode(&userDetails)

		user := user.NewUser(userDetails.Username, userDetails.Password)
		err := mgr.UserStore.Insert(user.Id, user)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		mgr.Logger.Info(fmt.Sprintf("Succesfully added user with user id %s", user.Id))
		mgr.Responder.Respond(w, http.StatusOK, "Succesfully added user")
	}
}

// NewUser inserts a user into the database
func (mgr *UserMgr) GetUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		user, err := mgr.UserStore.Get(userId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusNotFound, err)
		}

		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

func (mgr *UserMgr) Register() {
	mgr.Router.HandleFunc(NEW_USER, mgr.NewUser()).Methods("POST")
	mgr.Router.HandleFunc(GET_USER, mgr.GetUser()).Methods("GET")
}
