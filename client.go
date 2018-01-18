// ref. http://docs.repro.io/ja/dev/user-profile-api/index.html#user-profiles-payload
package repro

import (
	"bytes"
	"net/http"
)

const (
	updateUserProfileUrl = "https://api.repro.io/v2/user_profiles"
	tokenHeaderKey       = "X-Repro-Token"
)

// Global http client
var repro reproClient

func init() {
	client := &http.Client{}
	repro = reproClient{client, ""}
}

type reproClient struct {
	*http.Client
	token string
}

func (r *reproClient) SetToken(token string) {
	r.token = token
}

func Setup(token string) {
	repro.SetToken(token)
}

func SendUserProfile(body []byte) error {
	req, err := http.NewRequest(http.MethodPut, updateUserProfileUrl, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add(tokenHeaderKey, repro.token)
	resp, err := repro.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
