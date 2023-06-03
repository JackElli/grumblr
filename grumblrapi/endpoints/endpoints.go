package endpoints

import (
	"grumblrapi/endpoints/newgrumble"
	"grumblrapi/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Endpoints struct {
	Logger *zap.Logger
}

func NewEndpointsMgr() *Endpoints {
	return &Endpoints{
		Logger: &zap.Logger{},
	}
}

func (e *Endpoints) SetupEndpoints(r *mux.Router) {
	responder := responder.NewResponder()

	newGrumbleMgr := newgrumble.NewNewGrumbleMgr(r, responder)
	newGrumbleMgr.Register()
}
