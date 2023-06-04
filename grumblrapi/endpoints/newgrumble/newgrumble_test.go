package newgrumble

import (
	"bytes"
	"encoding/json"
	"grumblrapi/grumble"
	"grumblrapi/grumblestore"
	"grumblrapi/responder"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"gotest.tools/v3/assert"
)

var testRoute = "/grumble"

func TestNewGrumble(t *testing.T) {
	type testcase struct {
		desc           string
		expectedMsg    string
		expectedStatus int
	}

	responderMock := responder.NewResponder()
	grumbleStoreMock := grumblestore.NewGrumbleStoreMock()

	newGrumble := grumble.NewGrumble(
		"jack",
		"this is a test",
		time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC),
		"friends",
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
				responderMock,
				grumbleStoreMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			newGrumbleData, _ := json.Marshal(newGrumble)

			r, _ := http.NewRequest("POST", testRoute, bytes.NewBuffer(newGrumbleData))
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
