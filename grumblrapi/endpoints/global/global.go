package global

import (
	"fmt"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT = "/global/{category}"
)

type GlobalMgr struct {
	Logger        *zap.Logger
	Router        *mux.Router
	Responder     responder.Responder
	GrumbleStorer grumblestore.GrumbleStorer
}

func NewGlobalMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, grumbleStorer grumblestore.GrumbleStorer) *GlobalMgr {
	e := &GlobalMgr{
		Logger:        logger,
		Router:        router,
		Responder:     responder,
		GrumbleStorer: grumbleStorer,
	}
	e.Register()
	return e
}

// Grumbles returns all of the friends grumbles
func (mgr *GlobalMgr) GlobalGrumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		category := mux.Vars(req)["category"]

		var grumbles []grumblestore.Grumble
		var err error

		if category == "recents" {
			grumbles, err = mgr.GrumbleStorer.Query("SELECT grumbles.* from grumblr.dev.grumbles WHERE type='global' ORDER BY dateCreated DESC LIMIT 50")
			if err != nil {
				mgr.Responder.Error(w, 500, err)
			}
		} else {
			grumbles, err = mgr.GrumbleStorer.Query(fmt.Sprintf("SELECT grumbles.* from grumblr.dev.grumbles WHERE type='global' AND category='%s' ORDER BY dateCreated DESC LIMIT 50", category))
			if err != nil {
				mgr.Responder.Error(w, 500, err)
			}
		}

		mgr.Logger.Info("Successfully retrieved grumbles")
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GlobalMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.GlobalGrumbles()).Methods("GET")
}
