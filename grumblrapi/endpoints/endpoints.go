package endpoints

import (
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/newgrumble"
	"grumblrapi/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Endpoints struct {
	Logger *zap.Logger
}

func NewEndpointsMgr(logger *zap.Logger) *Endpoints {
	return &Endpoints{
		Logger: logger,
	}
}

func (e *Endpoints) SetupEndpoints(r *mux.Router) {
	responder := responder.NewResponder()

	public := r.PathPrefix("/").Subrouter()

	newGrumbleMgr := newgrumble.NewNewGrumbleMgr(public, responder)
	newGrumbleMgr.Register()
	grumblesMgr := grumbles.NewGrumblesMgr(public, responder)
	grumblesMgr.Register()
}
