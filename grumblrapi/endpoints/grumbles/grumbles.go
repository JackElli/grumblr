package grumbles

import (
	"fmt"
	"grumblrapi/main/categorystore"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT       = "/grumbles/{category}"
	CATEGORIES = "/grumbles/info/categories/{type}"
)

type GrumblesMgr struct {
	Env            string
	Logger         *zap.Logger
	Router         *mux.Router
	Responder      responder.Responder
	GrumbleStorer  grumblestore.GrumbleStorer
	CategoryStorer categorystore.CategoryStorer
}

func NewGrumblesMgr(router *mux.Router, env string, logger *zap.Logger, responder responder.Responder, grumbleStorer grumblestore.GrumbleStorer, categoryStorer categorystore.CategoryStorer) *GrumblesMgr {
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

		var grumbles []grumblestore.Grumble
		var err error

		if category == "recents" {
			grumbles, err = mgr.GrumbleStorer.Query(fmt.Sprintf("SELECT grumbles.* from grumblr.%s.grumbles WHERE type='friends' ORDER BY dateCreated DESC LIMIT 50", mgr.Env))
			if err != nil {
				mgr.Responder.Error(w, 500, err)
			}

		} else {
			grumbles, err = mgr.GrumbleStorer.Query(fmt.Sprintf("SELECT grumbles.* from grumblr.%s.grumbles WHERE type='friends' AND category='%s' ORDER BY dateCreated DESC LIMIT 50", mgr.Env, category))
			if err != nil {
				mgr.Responder.Error(w, 500, err)
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
		}

		mgr.Logger.Info(fmt.Sprintf("Successfully retrieved categories of type %s", _type))
		mgr.Responder.Respond(w, 200, categories)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.FriendsGrumbles()).Methods("GET")
	mgr.Router.HandleFunc(CATEGORIES, mgr.Categories()).Methods("GET")
}
