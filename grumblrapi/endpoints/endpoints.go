package endpoints

import (
	"grumblrapi/endpoints/global"
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/newgrumble"
	"grumblrapi/main/couchbase"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"

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

func (e *Endpoints) SetupEndpoints(r *mux.Router) error {
	cb, err := couchbase.NewCouchbase(e.Logger, "db", "Administrator", "password", "grumblr")
	if err != nil {
		return err
	}
	scope := cb.Bucket.Scope("dev")
	col := scope.Collection("grumbles")

	grumbleStorer := grumblestore.NewGrumbleStore(e.Logger, scope, col)
	responder := responder.NewResponder()

	public := r.PathPrefix("/").Subrouter()

	newGrumbleMgr := newgrumble.NewNewGrumbleMgr(public, e.Logger, responder, grumbleStorer)
	newGrumbleMgr.Register()
	grumblesMgr := grumbles.NewGrumblesMgr(public, e.Logger, responder, grumbleStorer)
	grumblesMgr.Register()
	globalMgr := global.NewGlobalMgr(public, e.Logger, responder, grumbleStorer)
	globalMgr.Register()

	return nil
}
