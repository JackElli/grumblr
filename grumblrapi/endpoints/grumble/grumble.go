package grumble

import (
	"encoding/json"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT        = "/grumble"
	ADD_COMMENT = "/grumble/{grumbleId}/comment"
)

type NewGrumbleMgr struct {
	Logger       *zap.Logger
	Router       *mux.Router
	Responder    responder.Responder
	GrumbleStore grumblestore.GrumbleStorer
}

func NewNewGrumbleMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, grumbleStore grumblestore.GrumbleStorer) *NewGrumbleMgr {
	e := &NewGrumbleMgr{
		Logger:       logger,
		Router:       router,
		Responder:    responder,
		GrumbleStore: grumbleStore,
	}
	e.Register()
	return e
}

// NewGrumble inserts a grumble into the database
func (mgr *NewGrumbleMgr) NewGrumble() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		//TODO add validation here
		// Get grumble details from endpoint call
		var grumble grumblestore.Grumble
		json.NewDecoder(req.Body).Decode(&grumble)

		// Create a new grumble from template
		grumble = *grumblestore.NewGrumble(
			grumble.CreatedBy,
			grumble.Message,
			grumble.Type,
			grumble.Category,
		)

		// Insert into database
		err := mgr.GrumbleStore.Insert(grumble.Id, &grumble)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		mgr.Logger.Info("Succesfully added grumble")
		mgr.Responder.Respond(w, http.StatusOK, "Successfully added grumble")
	}
}

func (mgr *NewGrumbleMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.NewGrumble()).Methods("POST")
}
