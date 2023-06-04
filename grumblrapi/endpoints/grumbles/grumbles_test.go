package grumbles

import (
	"encoding/json"
	"grumblrapi/grumble"
	"grumblrapi/grumblestore"
	"grumblrapi/responder"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gotest.tools/v3/assert"
)

var testRoute = "/grumbles"

func TestGrumbles(t *testing.T) {
	type testcase struct {
		desc              string
		expectedResultLen int
		expectedStatus    int
	}

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()

	testcases := []testcase{
		{
			desc:              "HAPPY added correct document",
			expectedResultLen: 2,
			expectedStatus:    200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			grumblesMgrMock := NewGrumblesMgr(
				rMock,
				responderMock,
				grumbleStoreMock,
			)
			grumblesMgrMock.Register()

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("GET", testRoute, nil)
			grumblesMgrMock.Router.ServeHTTP(w, r)

			var response []grumble.Grumble
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, len(response), testCase.expectedResultLen)
			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
