// ref. http://docs.repro.io/ja/dev/user-profile-api/index.html#user-profiles-payload
package repro

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func SendUserProfile(body []byte) (ReproResponse, error) {
	req, err := http.NewRequest(http.MethodPost, updateUserProfileUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(tokenHeaderKey, repro.token)
	resp, err := repro.Do(req)
	defer resp.Body.Close()

	rr := NewReproResponse(resp.StatusCode, resp.Header)
	if err != nil {
		return rr, err
	}
	if !rr.IsOK() {
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return rr, err
		}
		output := ReproError{}
		err = json.Unmarshal(rbody, &output)
		if err != nil {
			return rr, err
		}
		return rr, &output
	}

	return rr, nil
}
