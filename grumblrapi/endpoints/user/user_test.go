package user

import (
	"bytes"
	"encoding/json"
	"grumblrapi/main/responder"

	"grumblrapi/main/userstore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/v3/assert"
)

func TestAddFriend(t *testing.T) {
	type testcase struct {
		desc                    string
		expectedNumberOfFriends int
		userId                  string
		friendId                string
		expectedStatus          int
	}

	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()
	responderMock := responder.NewResponder()
	userStoreMock := userstore.NewUserStoreMock()

	testcases := []testcase{
		{
			desc:                    "HAPPY added friend",
			expectedNumberOfFriends: 1,
			userId:                  "test1",
			friendId:                "test2",
			expectedStatus:          200,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newUserMgrMock := NewNewUserMgr(
				rMock,
				loggerMock,
				responderMock,
				userStoreMock,
			)

			var friendURL = "/user/" + testCase.userId + "/friend"
			type Friend struct {
				FriendId string `json:"friendId"`
			}

			newFriend := Friend{
				FriendId: testCase.friendId,
			}

			w := httptest.NewRecorder()
			newFriendData, _ := json.Marshal(newFriend)

			// Add the friend
			r, _ := http.NewRequest("POST", friendURL, bytes.NewBuffer(newFriendData))
			newUserMgrMock.Router.ServeHTTP(w, r)

			// Check if user friend has been added
			var user userstore.User
			json.NewDecoder(w.Body).Decode(&user)

			assert.Equal(t, len(user.Friends), testCase.expectedNumberOfFriends)
			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
		})
	}
}
