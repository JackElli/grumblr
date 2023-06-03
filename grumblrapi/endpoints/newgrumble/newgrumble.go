package newgrumble

import (
	"grumblrapi/responder"
	"net/http"

	"github.com/gorilla/mux"
)

var route = "/grumble"

type NewGrumbleMgr struct {
	Router    *mux.Router
	Responder responder.Responder
}

func NewNewGrumbleMgr(router *mux.Router, responder responder.Responder) *NewGrumbleMgr {
	return &NewGrumbleMgr{
		Router:    router,
		Responder: responder,
	}
}

func (mgr *NewGrumbleMgr) NewGrumble() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		mgr.Responder.Respond(w, 200, "hello")
	}
}

func (mgr *NewGrumbleMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.NewGrumble()).Methods("POST")
}
