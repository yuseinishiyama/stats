package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

type RequestToken struct {
	Code string `json:"code"`
}

type Authorization struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

var Origin = "https://getpocket.com"

func GenerateToken(consumerKeyFile string, token string) error {
	consumerKey, err := getConsumerKey(consumerKeyFile)
	if err != nil {
		return err
	}
	return genAccessToken(consumerKey, token)
}

func ObtainRequestToken(consumerKey string, redirectURL string) (*RequestToken, error) {
	res := &RequestToken{}
	err := PostJSON(
		"/v3/oauth/request",
		map[string]string{
			"consumer_key": consumerKey,
			"redirect_uri": redirectURL,
		},
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ObtainAccessToken(consumerKey string, requestToken *RequestToken) (*Authorization, error) {
	res := &Authorization{}
	err := PostJSON(
		"/v3/oauth/authorize",
		map[string]string{
			"consumer_key": consumerKey,
			"code":         requestToken.Code,
		},
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GenerateAuthorizationURL(requestToken *RequestToken, redirectURL string) string {
	values := url.Values{"request_token": {requestToken.Code}, "redirect_uri": {redirectURL}}
	return fmt.Sprintf("%s/auth/authorize?%s", Origin, values.Encode())
}

func PostJSON(action string, data, res interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", Origin+action, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return doJSON(req, res)
}

func doJSON(req *http.Request, res interface{}) error {
	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}

func getConsumerKey(file string) (string, error) {
	consumerKey, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes.SplitN(consumerKey, []byte("\n"), 2)[0]), nil
}

func genAccessToken(consumerKey string, authFile string) error {
	accessToken := &Authorization{}
	accessToken, err := obtainAccessToken(consumerKey)
	if err != nil {
		return err
	}

	return saveJSONToFile(authFile, accessToken)
}

func obtainAccessToken(consumerKey string) (*Authorization, error) {
	ch := make(chan struct{})
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/favicon.ico" {
				http.Error(w, "Not Found", 404)
				return
			}

			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintln(w, "Authorized.")
			ch <- struct{}{}
		}))
	defer ts.Close()

	redirectURL := ts.URL

	requestToken, err := ObtainRequestToken(consumerKey, redirectURL)
	if err != nil {
		return nil, err
	}

	url := GenerateAuthorizationURL(requestToken, redirectURL)
	fmt.Println(url)

	<-ch

	return ObtainAccessToken(consumerKey, requestToken)
}

func saveJSONToFile(path string, v interface{}) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}

	defer w.Close()

	return json.NewEncoder(w).Encode(v)
}

func loadJSONFromFile(path string, v interface{}) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}

	defer r.Close()

	return json.NewDecoder(r).Decode(v)
}
