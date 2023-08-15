package auth

import (
	"bytes"
	"encoding/json"
	"grumblrapi/main/responder"
	"grumblrapi/main/usermgr"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

var testRoute = "/auth"

func TestNewUser(t *testing.T) {
	type testcase struct {
		desc           string
		expectedMsg    string
		expectedStatus int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()
	responderMock := responder.NewResponder()
	userStoreMock := usermgr.NewUserStoreMock()

	testcases := []testcase{
		{
			desc:           "HAPPY added correct document",
			expectedStatus: 200,
		},
	}

	newUser := usermgr.NewUser(
		"jack",
		"this is a test",
	)

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			authMgrMock := NewAuthMgr(
				rMock,
				"dev",
				loggerMock,
				responderMock,
				userStoreMock,
				nil,
			)

			w := httptest.NewRecorder()
			newUserData, _ := json.Marshal(newUser)

			r, _ := http.NewRequest("POST", NEW_USER, bytes.NewBuffer(newUserData))
			authMgrMock.Router.ServeHTTP(w, r)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
