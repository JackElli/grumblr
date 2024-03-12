package endpoints

import (
	"grumblrapi/endpoints/auth"
	"grumblrapi/endpoints/global"
	"grumblrapi/endpoints/grumble"
	"grumblrapi/endpoints/grumbles"
	"grumblrapi/endpoints/middleware"
	"grumblrapi/endpoints/user"
	"grumblrapi/main/categorymgr"
	"grumblrapi/main/couchbase"
	"grumblrapi/main/grumblemgr"
	"grumblrapi/main/jwtmgr"
	"grumblrapi/main/responder"
	"grumblrapi/main/usermgr"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

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
	return &Login{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Bucket:   os.Getenv("BUCKET"),
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
	responder := responder.NewResponder()

	// Set up managers
	userMgr := usermgr.NewUserStore(env, e.Logger, scope)
	jwtMgr := jwtmgr.NewJWTManager(SECRET_KEY)
	middlewareMgr := middleware.NewMiddlewareMgr(jwtMgr)
	grumbleMgr := grumblemgr.NewGrumbleStore(e.Logger, scope, userMgr)
	categoryMgr := categorymgr.NewCategoryStore(e.Logger, scope)

	// The public endpoint for auth
	public := r.PathPrefix("/").Subrouter()
	auth.NewAuthMgr(public, env, e.Logger, responder, userMgr, jwtMgr)

	// For the endpoints that are restricted by auth
	restricted := r.PathPrefix("/").Subrouter()
	restricted.Use(middlewareMgr.Middleware)
	grumble.NewNewGrumbleMgr(restricted, e.Logger, responder, grumbleMgr, userMgr)
	user.NewNewUserMgr(restricted, e.Logger, responder, userMgr)
	grumbles.NewGrumblesMgr(restricted, env, e.Logger, responder, grumbleMgr, categoryMgr)
	global.NewGlobalMgr(restricted, env, e.Logger, responder, grumbleMgr)

	return nil
}
