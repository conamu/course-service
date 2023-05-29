package api

import (
	"bytes"
	"course-service/course"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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
	log.Println(err)
	log.Println(string(data))
	req, err := http.NewRequest("POST", "http://auth-service:8080/validate", buffer)
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
	log.Println(resp.StatusCode)
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
