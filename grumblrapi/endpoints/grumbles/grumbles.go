package grumbles

import (
	"grumblrapi/grumble"
	"grumblrapi/responder"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var route = "/grumbles"

type GrumblesMgr struct {
	Router    *mux.Router
	Responder responder.Responder
}

func NewGrumblesMgr(router *mux.Router, responder responder.Responder) *GrumblesMgr {
	return &GrumblesMgr{
		Router:    router,
		Responder: responder,
	}
}

func (mgr *GrumblesMgr) Grumbles() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		mockGrumbles := []grumble.Grumble{
			{
				CreatedBy: "user:1",
				Message:   "this is the very first grumble",
				Date:      time.Now(),
			},
		}
		mgr.Responder.Respond(w, 200, mockGrumbles)
	}
}

func (mgr *GrumblesMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.Grumbles()).Methods("GET")
}
