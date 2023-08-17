package grumbles

import (
	"fmt"
	"grumblrapi/main/categorymgr"
	"grumblrapi/main/grumblemgr"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/couchbase/gocb/v2"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT             = "/grumbles"
	GRUMBLES         = ROOT + "/{category}"
	CATEGORIES       = ROOT + "/info/categories/{type}"
	GRUMBLES_BY_USER = ROOT + "/user/{userId}"
)

type GrumblesMgr struct {
	Env            string
	Logger         *zap.Logger
	Router         *mux.Router
	Responder      responder.Responder
	GrumbleStorer  grumblemgr.GrumbleStorer
	CategoryStorer categorymgr.CategoryStorer
}

func NewGrumblesMgr(router *mux.Router, env string, logger *zap.Logger, responder responder.Responder, grumbleStorer grumblemgr.GrumbleStorer, categoryStorer categorymgr.CategoryStorer) *GrumblesMgr {
	e := &GrumblesMgr{
		Env:            env,
		Logger:         logger,
		Router:         router,
		Responder:      responder,
		GrumbleStorer:  grumbleStorer,
		CategoryStorer: categoryStorer,
	}
	e.Register()
	return e
}

// FriendsGrumbles returns all of the friends grumbles
func (mgr *GrumblesMgr) FriendsGrumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		category := mux.Vars(req)["category"]

		var grumbles []grumblemgr.Grumble
		var err error

		if category == "recents" {
			grumbles, err = mgr.GrumbleStorer.Query(
				fmt.Sprintf("SELECT grumbles.* from grumblr.%s.grumbles WHERE type='friends' ORDER BY dateCreated DESC LIMIT 50", mgr.Env),
				nil,
			)
			if err != nil {
				mgr.Responder.Error(w, 500, err)
				return
			}

		} else {
			grumbles, err = mgr.GrumbleStorer.Query(
				fmt.Sprintf("SELECT grumbles.* from grumblr.%s.grumbles WHERE type='friends' AND category='%s' ORDER BY dateCreated DESC LIMIT 50", mgr.Env, category),
				nil,
			)
			if err != nil {
				mgr.Responder.Error(w, 500, err)
				return
			}
		}

		mgr.Logger.Info("Successfully retrieved grumbles")
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

// Categories returns all of the categories in this friend group
func (mgr *GrumblesMgr) Categories() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		_type := mux.Vars(req)["type"]

		categories, err := mgr.CategoryStorer.Query(
			fmt.Sprintf("SELECT categories.* from grumblr.%s.categories WHERE type='%s' LIMIT 7", mgr.Env, _type),
		)
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		mgr.Logger.Info(fmt.Sprintf("Successfully retrieved categories of type %s", _type))
		mgr.Responder.Respond(w, 200, categories)
	}
}

// GetGrumblesByUser returns all of the grumbles created by the userId passed
func (mgr *GrumblesMgr) GetGrumblesByUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		nps := map[string]interface{}{
			"userId": userId,
		}

		grumbles, err := mgr.GrumbleStorer.Query(fmt.Sprintf(
			"SELECT grumbles.* from grumblr.%s.grumbles WHERE createdBy=$userId ORDER BY dateCreated DESC LIMIT 50", mgr.Env),
			&gocb.QueryOptions{
				NamedParameters: nps,
			},
		)
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		// TODO does this need to be locked down to the user??

		mgr.Logger.Info(fmt.Sprintf("Successfully retrieved grumbles for user %s", userId))
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(GRUMBLES, mgr.FriendsGrumbles()).Methods("GET")
	mgr.Router.HandleFunc(GRUMBLES_BY_USER, mgr.GetGrumblesByUser()).Methods("GET")
	mgr.Router.HandleFunc(CATEGORIES, mgr.Categories()).Methods("GET")
}
