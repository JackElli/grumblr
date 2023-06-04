package grumbles

import (
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var route = "/grumbles"

type GrumblesMgr struct {
	Logger        *zap.Logger
	Router        *mux.Router
	Responder     responder.Responder
	GrumbleStorer grumblestore.GrumbleStorer
}

func NewGrumblesMgr(logger *zap.Logger, router *mux.Router, responder responder.Responder, grumbleStorer grumblestore.GrumbleStorer) *GrumblesMgr {
	return &GrumblesMgr{
		Logger:        logger,
		Router:        router,
		Responder:     responder,
		GrumbleStorer: grumbleStorer,
	}
}

// Grumbles returns all of the friends grumbles
func (mgr *GrumblesMgr) FriendsGrumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		grumbles, err := mgr.GrumbleStorer.Query("SELECT grumbles.* from grumblr.dev.grumbles WHERE type='friends' ORDER BY dateCreated DESC LIMIT 50")
		if err != nil {
			mgr.Responder.Error(w, 500, err)
		}
		mgr.Logger.Info("Successfully retrieved grumbles")
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.FriendsGrumbles()).Methods("GET")
}
