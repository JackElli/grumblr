package global

import (
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var route = "/global"

type GlobalMgr struct {
	Logger        *zap.Logger
	Router        *mux.Router
	Responder     responder.Responder
	GrumbleStorer grumblestore.GrumbleStorer
}

func NewGlobalMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, grumbleStorer grumblestore.GrumbleStorer) *GlobalMgr {
	return &GlobalMgr{
		Logger:        logger,
		Router:        router,
		Responder:     responder,
		GrumbleStorer: grumbleStorer,
	}
}

// Grumbles returns all of the friends grumbles
func (mgr *GlobalMgr) GlobalGrumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		grumbles, err := mgr.GrumbleStorer.Query("SELECT grumbles.* from grumblr.dev.grumbles WHERE type='global' ORDER BY dateCreated DESC LIMIT 50")
		if err != nil {
			mgr.Responder.Error(w, 500, err)
		}
		mgr.Logger.Info("Successfully retrieved global grumbles")
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GlobalMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.GlobalGrumbles()).Methods("GET")
}
