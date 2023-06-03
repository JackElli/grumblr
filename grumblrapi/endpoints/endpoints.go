package endpoints

import (
	"grumblrapi/couchbase"
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/newgrumble"
	"grumblrapi/grumblestore"
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

func (e *Endpoints) SetupEndpoints(r *mux.Router) error {
	cb, err := couchbase.NewCouchbase(e.Logger, "db", "Administrator", "password", "grumblr")
	if err != nil {
		return err
	}
	scope := cb.Bucket.Scope("dev")
	col := cb.Collection("grumbles")

	grumbleStorer := grumblestore.NewGrumbleStore(e.Logger, scope, col)

	responder := responder.NewResponder()

	public := r.PathPrefix("/").Subrouter()

	newGrumbleMgr := newgrumble.NewNewGrumbleMgr(public, responder)
	newGrumbleMgr.Register()
	grumblesMgr := grumbles.NewGrumblesMgr(public, responder, grumbleStorer)
	grumblesMgr.Register()

	return nil
}
