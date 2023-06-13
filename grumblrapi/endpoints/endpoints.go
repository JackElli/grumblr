package endpoints

import (
	"grumblrapi/endpoints/global"
	"grumblrapi/endpoints/grumble"
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/user"
	"grumblrapi/main/couchbase"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"grumblrapi/main/userstore"

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

	grumbleStorer := grumblestore.NewGrumbleStore(e.Logger, scope)
	userStorer := userstore.NewUserStore(e.Logger, scope)
	responder := responder.NewResponder()

	public := r.PathPrefix("/").Subrouter()

	newGrumbleMgr := grumble.NewNewGrumbleMgr(public, e.Logger, responder, grumbleStorer)
	newGrumbleMgr.Register()
	newUserMgr := user.NewNewUserMgr(public, e.Logger, responder, userStorer)
	newUserMgr.Register()
	grumblesMgr := grumbles.NewGrumblesMgr(public, e.Logger, responder, grumbleStorer)
	grumblesMgr.Register()
	globalMgr := global.NewGlobalMgr(public, e.Logger, responder, grumbleStorer)
	globalMgr.Register()

	return nil
}
