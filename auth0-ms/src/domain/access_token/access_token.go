package accesstoken

import "time"

const (
	exprirationTime = 24
)

type Accesstoken struct {
	Accesstoken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() Accesstoken {
	return Accesstoken{
		Expires: time.Now().UTC().Add(exprirationTime * time.Hour).Unix(),
	}
}

func (at Accesstoken) IsExpired() bool {
	return false
}

// WEB 123
// ANDROID 234
