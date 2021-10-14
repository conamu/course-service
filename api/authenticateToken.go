package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type authTokenRequest struct {
	Token string `json:"token"`
}

func authenticateToken(token string) error {
	payload := authTokenRequest{Token: token}
	data, err := json.MarshalIndent(&payload, "", " ")
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "https://auth.karl-bock.academy/validate", buffer)
	if err != nil {
		return err
	}
	req.Header = map[string][]string{
		"X-KBU-Auth": {"abcdefghijklmnopqrstuvwxyz"},
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("auth service did not respond OK")
	}
	return nil
}
