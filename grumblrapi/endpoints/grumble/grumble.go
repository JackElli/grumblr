package grumble

import (
	"encoding/json"
	"fmt"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT        = "/grumble"
	GET_GRUMBLE = "/grumble/{grumbleId}"
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
		mgr.Responder.Respond(w, http.StatusOK, grumble)
	}
}

// GetGrumble gets the grumble based on the id passed
func (mgr *NewGrumbleMgr) GetGrumble() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		grumbleId := mux.Vars(req)["grumbleId"]

		// Get from database
		grumble, err := mgr.GrumbleStore.Get(grumbleId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		mgr.Logger.Info(fmt.Sprintf("Succesfully retrieved grumble for %s", grumbleId))
		mgr.Responder.Respond(w, http.StatusOK, grumble)
	}
}

// Comment adds a comment to the grumble passed
func (mgr *NewGrumbleMgr) Comment() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		grumbleId := mux.Vars(req)["grumbleId"]
		//TODO add validation here
		// Get grumble details from endpoint call
		var comment grumblestore.Comment
		json.NewDecoder(req.Body).Decode(&comment)

		// Create a new comment from template
		comment = *grumblestore.NewComment(
			comment.CreatedBy,
			comment.Message,
		)

		// Get grumble to update and add comment
		grumble, err := mgr.GrumbleStore.Get(grumbleId)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		// Prepend comment
		grumble.Comments = append([]grumblestore.Comment{comment}, grumble.Comments...)

		// Update database
		err = mgr.GrumbleStore.Update(grumbleId, grumble)
		if err != nil {
			mgr.Responder.Error(w, http.StatusInternalServerError, err)
			return
		}

		mgr.Logger.Info("Succesfully added comment")
		mgr.Responder.Respond(w, http.StatusOK, "Successfully added comment")
	}
}

func (mgr *NewGrumbleMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.NewGrumble()).Methods("POST")
	mgr.Router.HandleFunc(GET_GRUMBLE, mgr.GetGrumble()).Methods("GET")
	mgr.Router.HandleFunc(ADD_COMMENT, mgr.Comment()).Methods("POST")
}
