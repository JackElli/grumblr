package grumbles

import (
	"encoding/json"
	"grumblrapi/main/categorymgr"
	"grumblrapi/main/grumblemgr"
	"grumblrapi/main/responder"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

func TestGrumbles(t *testing.T) {
	type testcase struct {
		desc              string
		expectedResultLen int
		expectedStatus    int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()
	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblemgr.NewGrumbleStoreMock()

	testcases := []testcase{
		{
			desc:              "HAPPY retrieved grumbles",
			expectedResultLen: 2,
			expectedStatus:    200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			grumblesMgrMock := NewGrumblesMgr(
				rMock,
				"dev",
				loggerMock,
				responderMock,
				grumbleStoreMock,
				nil,
			)

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("GET", ROOT, nil)
			grumblesMgrMock.Router.ServeHTTP(w, r)

			var response []grumblemgr.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, len(response), testCase.expectedResultLen)
			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}

const CATEGORIES_ROOT = "/grumbles/info/categories/"

func TestCategories(t *testing.T) {
	type testcase struct {
		desc           string
		_type          string
		expectedResult []categorymgr.Category
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()
	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblemgr.NewGrumbleStoreMock()
	categoryStoreMock := categorymgr.NewCategoryStoreMock()

	testcases := []testcase{
		{
			desc:  "HAPPY retrieved categories",
			_type: "friends",
			expectedResult: []categorymgr.Category{
				{
					Id:   "testcat1",
					Type: "friends",
					People: []string{
						"jack",
					},
					Name: "Weather",
				},
			},
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			grumblesMgrMock := NewGrumblesMgr(
				rMock,
				"dev",
				loggerMock,
				responderMock,
				grumbleStoreMock,
				categoryStoreMock,
			)

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("GET", CATEGORIES_ROOT+testCase._type, nil)
			grumblesMgrMock.Router.ServeHTTP(w, r)

			var response []categorymgr.Category
			json.NewDecoder(w.Body).Decode(&response)

			assert.DeepEqual(t, response, testCase.expectedResult)
			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
