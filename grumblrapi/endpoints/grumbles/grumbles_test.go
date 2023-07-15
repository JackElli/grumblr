package grumbles

import (
	"encoding/json"
	"grumblrapi/main/grumblestore"
	"grumblrapi/main/responder"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

var testRoute = "/grumbles"

func TestGrumbles(t *testing.T) {
	type testcase struct {
		desc              string
		expectedResultLen int
		expectedStatus    int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()
	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()

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
				loggerMock,
				responderMock,
				grumbleStoreMock,
				nil, // add test
			)
			grumblesMgrMock.Register()

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("GET", testRoute, nil)
			grumblesMgrMock.Router.ServeHTTP(w, r)

			var response []grumblestore.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, len(response), testCase.expectedResultLen)
			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
