package accesstoken

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("not null")
	}
	if at.AccessToken != "" {
		t.Error("new Access token should not have define the new id")
	}
	if at.UserID != 0 {
		t.Error("new access token should not have an assoicated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token is created for 3h")
	}
}
