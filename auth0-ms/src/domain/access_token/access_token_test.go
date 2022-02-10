package access_token

import (
	"testing"
)

const (
	exprirationTime = 24
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access")
	}
	if at.AccessToken != "" {
		t.Error("new access should not hace defined")
	}
	if at.UserId != 0 {
		t.Error("new access should not have an associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("empty access token expired in 3 hous")
	}
}
