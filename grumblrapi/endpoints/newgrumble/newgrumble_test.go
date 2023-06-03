package newgrumble

import (
	"grumblrapi/responder"
	"net/http"
	"net/http/httptest"
	"testing"

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

	testcases := []testcase{
		{
			desc:           "HAPPY added correct document",
			expectedMsg:    "Added document successfully",
			expectedStatus: 200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newGrumbleMgrMock := NewNewGrumbleMgr(
				rMock,
				responderMock,
			)
			newGrumbleMgrMock.Register()

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("POST", testRoute, nil)
			newGrumbleMgrMock.Router.ServeHTTP(w, r)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
