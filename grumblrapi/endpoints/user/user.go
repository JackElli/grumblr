package user

import (
	"encoding/json"
	"grumblrapi/main/responder"
	"grumblrapi/main/usermgr"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT       = "/user"
	GET_USER   = ROOT + "/{userId}"
	ADD_FRIEND = ROOT + "/{userId}/friend"
)

type UserMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	UserStore usermgr.UserStorer
}

func NewNewUserMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, userStore usermgr.UserStorer) *UserMgr {
	e := &UserMgr{
		Logger:    logger,
		Router:    router,
		Responder: responder,
		UserStore: userStore,
	}
	e.Register()
	return e
}

// GetUser returns the user from the db
func (mgr *UserMgr) GetUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		user, err := mgr.UserStore.Get(userId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusNotFound, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

// AddFriend adds a friendsId to the userIds friends list
func (mgr *UserMgr) AddFriend() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		// Decond data from request
		var Friend struct {
			FriendId string `json:"friendId"`
		}
		json.NewDecoder(req.Body).Decode(&Friend)

		// Get the friends info
		friend, err := mgr.UserStore.Get(Friend.FriendId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusNotFound, err)
			return
		}

		// Get the user who's wanting to add the friend
		user, err := mgr.UserStore.Get(userId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusNotFound, err)
			return
		}
		user.Friends = append(user.Friends, *friend)

		// Update the users information with the new friends list
		err = mgr.UserStore.Update(userId, user)
		if err != nil {
			mgr.Responder.Error(w, http.StatusNotFound, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

func (mgr *UserMgr) Register() {
	mgr.Router.HandleFunc(GET_USER, mgr.GetUser()).Methods("GET")
	mgr.Router.HandleFunc(ADD_FRIEND, mgr.AddFriend()).Methods("POST")
}
