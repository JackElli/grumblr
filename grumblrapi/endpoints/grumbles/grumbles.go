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

func (mgr *GrumblesMgr) Grumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		grumbles, err := mgr.GrumbleStorer.GetAll()
		if err != nil {
			mgr.Responder.Error(w, 500, err)
		}
		mgr.Responder.Respond(w, 200, grumbles)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.Grumbles()).Methods("GET")
}
