package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Slack struct {
	TokenFile string
}

type Stars struct {
	Items []interface{}
}

func (s Slack) Get() (int, error) {
	token, err := s.getToken(s.TokenFile)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://slack.com/api/stars.list?token=%s&limit=1000", token), nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	res := &Stars{}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(res); err != nil {
		return 0, err
	}

	return len(res.Items), nil
}

func (s Slack) getToken(file string) (string, error) {
	token, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes.SplitN(token, []byte("\n"), 2)[0]), nil
}
