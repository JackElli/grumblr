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

type Login struct {
	Username string
	Password string
	Bucket   string
}

// getLogin returns the username, password and bucket of different envs
// this is a way to get around using secrets
func getLogin(env string) *Login {
	if env == "prod" {
		return &Login{
			Username: "Administrator",
			Password: "Grumblr2!",
			Bucket:   "grumblr",
		}
	} else {
		return &Login{
			Username: "Administrator",
			Password: "password",
			Bucket:   "grumblr",
		}
	}
}

// SetupEndpoints sets up the means with which to talk to grumblrapi
func (e *Endpoints) SetupEndpoints(env string, r *mux.Router) error {
	loginDetails := getLogin(env)
	cb, err := couchbase.NewCouchbase(
		e.Logger,
		"db",
		loginDetails.Username,
		loginDetails.Password,
		loginDetails.Bucket,
	)
	if err != nil {
		return err
	}

	// Set the environment
	scope := cb.Bucket.Scope(env)

	// Set up storers
	grumbleStorer := grumblestore.NewGrumbleStore(e.Logger, scope)
	categoryStorer := categorystore.NewCategoryStore(e.Logger, scope)
	userStorer := userstore.NewUserStore(e.Logger, scope)

	// For responding to the user
	responder := responder.NewResponder()

	// For the endpoints that aren't restricted by auth
	public := r.PathPrefix("/").Subrouter()
	grumble.NewNewGrumbleMgr(public, e.Logger, responder, grumbleStorer, userStorer)
	user.NewNewUserMgr(public, e.Logger, responder, userStorer)
	grumbles.NewGrumblesMgr(public, env, e.Logger, responder, grumbleStorer, categoryStorer)
	global.NewGlobalMgr(public, env, e.Logger, responder, grumbleStorer)

	return nil
}
