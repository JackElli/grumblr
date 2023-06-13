package grumble

import (
	"encoding/json"
	"grumblrapi/main/grumble"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var route = "/grumble"

type NewGrumbleMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	DataStore grumblestore.GrumbleStorer
}

func NewNewGrumbleMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, dataStore grumblestore.GrumbleStorer) *NewGrumbleMgr {
	return &NewGrumbleMgr{
		Logger:    logger,
		Router:    router,
		Responder: responder,
		DataStore: dataStore,
	}
}

// NewGrumble inserts a grumble into the database
func (mgr *NewGrumbleMgr) NewGrumble() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		//TODO add validation here
		var grumble grumble.Grumble
		json.NewDecoder(req.Body).Decode(&grumble)

		grumbleId := uuid.New().String()
		grumble.Id = grumbleId

		err := mgr.DataStore.Insert(grumbleId, &grumble)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}
		mgr.Logger.Info("Succesfully added grumble")
		mgr.Responder.Respond(w, http.StatusOK, "Successfully added grumble")
	}
}

func (mgr *NewGrumbleMgr) Register() {
	mgr.Router.HandleFunc(route, mgr.NewGrumble()).Methods("POST")
}
