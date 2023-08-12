package grumble

import (
	"bytes"
	"encoding/json"

	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"grumblrapi/main/userstore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

func TestNewGrumble(t *testing.T) {
	const ROOT_TEST = "/grumble"

	type testcase struct {
		desc           string
		expectedMsg    string
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()
	userStoreMock := userstore.NewUserStoreMock()

	newGrumble := grumblestore.NewGrumble(
		"jack",
		grumblestore.Text,
		"this is a test",
		"friends",
		"Test",
	)

	testcases := []testcase{
		{
			desc:           "HAPPY added correct document",
			expectedMsg:    "Successfully added grumble",
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newGrumbleMgrMock := NewNewGrumbleMgr(
				rMock,
				loggerMock,
				responderMock,
				grumbleStoreMock,
				userStoreMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			newGrumbleData, _ := json.Marshal(newGrumble)

			r, _ := http.NewRequest("POST", ROOT_TEST, bytes.NewBuffer(newGrumbleData))
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}

func TestComment(t *testing.T) {

	type testcase struct {
		desc           string
		expectedResult grumblestore.Grumble
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()
	userStoreMock := userstore.NewUserStoreMock()

	newGrumble := grumblestore.NewGrumble(
		"jack",
		grumblestore.Text,
		"this is a test",
		"friends",
		"Test",
	)

	var COMMENT_TEST = "/grumble/" + newGrumble.Id + "/comment"

	grumbleMock := newGrumble
	grumbleMock.Comments = []grumblestore.Comment{
		{
			Message: "hello this is a test comment",
		},
	}

	var payload = grumblestore.NewComment(
		"jack",
		"hello this is a test comment",
	)

	testcases := []testcase{
		{
			desc:           "HAPPY added comment",
			expectedResult: *grumbleMock,
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newGrumbleMgrMock := NewNewGrumbleMgr(
				rMock,
				loggerMock,
				responderMock,
				grumbleStoreMock,
				userStoreMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			newGrumbleData, _ := json.Marshal(payload)
			r, _ := http.NewRequest("POST", COMMENT_TEST, bytes.NewBuffer(newGrumbleData))
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			var response grumblestore.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.Equal(t, response.Comments[0].Message, testCase.expectedResult.Comments[0].Message)
		})
	}
}

func TestAgree(t *testing.T) {

	type testcase struct {
		desc           string
		expectedResult grumblestore.Grumble
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()
	userStoreMock := userstore.NewUserStoreMock()

	newGrumble := grumblestore.NewGrumble(
		"jack",
		grumblestore.Text,
		"this is a test",
		"friends",
		"Test",
	)

	var AGREE_TEST = "/grumble/" + newGrumble.Id + "/agree"

	grumbleMock := newGrumble
	grumbleMock.Agrees["test1"] = true

	type body struct {
		UserId string `json:"userId"`
	}

	var payload = body{
		UserId: "test1",
	}

	testcases := []testcase{
		{
			desc:           "HAPPY agree",
			expectedResult: *grumbleMock,
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newGrumbleMgrMock := NewNewGrumbleMgr(
				rMock,
				loggerMock,
				responderMock,
				grumbleStoreMock,
				userStoreMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			agreeData, _ := json.Marshal(payload)
			r, _ := http.NewRequest("POST", AGREE_TEST, bytes.NewBuffer(agreeData))
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			var response grumblestore.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.DeepEqual(t, response.Agrees, testCase.expectedResult.Agrees)
		})
	}
}

func TestDisagree(t *testing.T) {

	type testcase struct {
		desc           string
		expectedResult grumblestore.Grumble
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()
	userStoreMock := userstore.NewUserStoreMock()

	newGrumble := grumblestore.NewGrumble(
		"jack",
		grumblestore.Text,
		"this is a test",
		"friends",
		"Test",
	)

	var DISAGREE_TEST = "/grumble/" + newGrumble.Id + "/disagree"

	grumbleMock := newGrumble
	grumbleMock.Disagrees["test1"] = true

	type body struct {
		UserId string `json:"userId"`
	}

	var payload = body{
		UserId: "test1",
	}

	testcases := []testcase{
		{
			desc:           "HAPPY disagree",
			expectedResult: *grumbleMock,
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newGrumbleMgrMock := NewNewGrumbleMgr(
				rMock,
				loggerMock,
				responderMock,
				grumbleStoreMock,
				userStoreMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			agreeData, _ := json.Marshal(payload)
			r, _ := http.NewRequest("POST", DISAGREE_TEST, bytes.NewBuffer(agreeData))
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			var response grumblestore.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.DeepEqual(t, response.Agrees, testCase.expectedResult.Agrees)
		})
	}
}
