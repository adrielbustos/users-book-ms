package accesstoken

import (
	"testing"
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
