package endpoints

import (
	"grumblrapi/endpoints/global"
	"grumblrapi/endpoints/grumble"
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/user"
	"grumblrapi/main/categorystore"
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

	// Set up storers
	grumbleStorer := grumblestore.NewGrumbleStore(e.Logger, scope)
	categoryStorer := categorystore.NewCategoryStore(e.Logger, scope)
	userStorer := userstore.NewUserStore(e.Logger, scope)
	responder := responder.NewResponder()

	// For the endpoints that aren't restricted by auth
	public := r.PathPrefix("/").Subrouter()
	grumble.NewNewGrumbleMgr(public, e.Logger, responder, grumbleStorer)
	user.NewNewUserMgr(public, e.Logger, responder, userStorer)
	grumbles.NewGrumblesMgr(public, e.Logger, responder, grumbleStorer, categoryStorer)
	global.NewGlobalMgr(public, e.Logger, responder, grumbleStorer)

	return nil
}
