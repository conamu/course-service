package api

import (
	"bytes"
	"course-service/course"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type authTokenRequest struct {
	Token string `json:"token"`
}

func AuthenticateToken(token string) (string, error) {
	payload := authTokenRequest{Token: token}
	data, err := json.MarshalIndent(&payload, "", " ")
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "https://auth.karl-bock.academy/validate", buffer)
	if err != nil {
		return "", err
	}
	req.Header = map[string][]string{
		"X-KBU-Auth": {"abcdefghijklmnopqrstuvwxyz"},
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("auth service did not respond OK")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	valRes := &course.ValidationResponse{}
	err = json.Unmarshal(body, valRes)
	if err != nil {
		return "", err
	}
	return valRes.Role, nil
}
