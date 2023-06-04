package grumbles

import (
	"grumblrapi/grumblestore"
	"grumblrapi/responder"
	"net/http"

	"github.com/gorilla/mux"
)

var route = "/grumbles"

type GrumblesMgr struct {
	Router        *mux.Router
	Responder     responder.Responder
	GrumbleStorer grumblestore.GrumbleStorer
}

func NewGrumblesMgr(router *mux.Router, responder responder.Responder, grumbleStorer grumblestore.GrumbleStorer) *GrumblesMgr {
	return &GrumblesMgr{
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
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.FriendsGrumbles()).Methods("GET")
}
