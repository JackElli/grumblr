package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"grumblrapi/main/jwtmgr"
	"grumblrapi/main/responder"
	"grumblrapi/main/usermgr"
	"net/http"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT     = "/auth"
	NEW_USER = "/user"

	ONE_DAY = 1 * 24 * 60 * time.Minute
)

type AuthMgr struct {
	Env       string
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	UserStore usermgr.UserStorer
	JWTMgr    jwtmgr.JWTManager
}

func NewAuthMgr(router *mux.Router, env string, logger *zap.Logger, responder responder.Responder, userStore usermgr.UserStorer, jwtMgr jwtmgr.JWTManager) *AuthMgr {
	e := &AuthMgr{
		Env:       env,
		Logger:    logger,
		Router:    router,
		Responder: responder,
		UserStore: userStore,
		JWTMgr:    jwtMgr,
	}
	e.Register()
	return e
}

// Auth endpoint returns a JWT set for an expiration if the user exists
// it also sets a cookie for the client of this JWT
func (mgr *AuthMgr) Auth() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		// Check if we have a cookie, if it's a valid cookie
		// we want to return the user info back
		c, err := req.Cookie("token")
		if err == nil {
			hasValidCookie, user := mgr.checkCookie(w, c)
			if !hasValidCookie {
				mgr.Responder.Error(w, 401, errors.New("unauthorised"))
				return
			}
			mgr.Responder.Respond(w, 200, user)
			return
		}

		var authVal struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		err = json.NewDecoder(req.Body).Decode(&authVal)
		if err != nil {
			mgr.Responder.Error(w, 401, err)
			return
		}

		auth, err := authUser(mgr, authVal.Username, authVal.Password)
		if !auth {
			mgr.Responder.Error(w, 401, err)
			return
		}

		user, err := mgr.UserStore.GetByUsername(authVal.Username)
		if err != nil {
			mgr.Responder.Error(w, 401, err)
			return
		}

		expirationDate := time.Now().Add(ONE_DAY)
		jwt, err := mgr.JWTMgr.CreateJWT(user, expirationDate)
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   jwt,
			Expires: expirationDate,
		})

		mgr.Responder.Respond(w, 200, user)
	}
}

// checkCookie, checks to see whether the auth endpoint hit already
// contains a cookie and therfore we don't need to create a new
// one
func (mgr *AuthMgr) checkCookie(w http.ResponseWriter, c *http.Cookie) (bool, *usermgr.User) {
	jwtStr := c.Value
	// If the jwt passed in the cookie is invalid, we want to
	// remove the cookie
	jwtClaims, jwtOk := mgr.JWTMgr.ParseJWT(jwtStr)
	if !jwtOk {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Expires: time.Now(),
		})
		return false, nil
	}

	claims := *jwtClaims
	userId := claims["user"].(string)
	user, err := mgr.UserStore.Get(userId)

	// If the user no longer exists in the db, we want to remove
	// the cookie
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Expires: time.Now(),
		})
		return false, nil
	}

	return true, user
}

// authUser returns true if user is authed, false if not
func authUser(mgr *AuthMgr, username string, password string) (bool, error) {
	user, err := mgr.UserStore.GetByUsername(username)
	if err != nil {
		mgr.Logger.Warn(err.Error())
		return false, err
	}

	if user.Password != password {
		err := errors.New("Incorrect username or password")
		mgr.Logger.Warn(err.Error())
		return false, err
	}
	return true, nil
}

// NewUser inserts a user into the database
func (mgr *AuthMgr) NewUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		//TODO add validation here
		var userDetails struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		json.NewDecoder(req.Body).Decode(&userDetails)

		user := usermgr.NewUser(userDetails.Username, userDetails.Password)
		err := mgr.UserStore.Insert(user.Id, user, &gocb.InsertOptions{
			// This will change when we move away from demo
			Expiry: 24 * time.Hour,
		})
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		mgr.Logger.Info(fmt.Sprintf("Succesfully added user with user id %s", user.Id))
		mgr.Responder.Respond(w, http.StatusOK, "Succesfully added user")
	}
}

func (mgr *AuthMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.Auth()).Methods("POST")
	mgr.Router.HandleFunc(NEW_USER, mgr.NewUser()).Methods("POST")
}
