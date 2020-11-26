package implement

import (
	"accesshttpclient"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"time"
)

type HTTPClientDomain struct {
	Client *http.Client
}

type response struct {
	ID  string `json:"id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func NewHTTPDomain() HTTPClientDomain {
	return HTTPClientDomain{
		Client: &http.Client{
			Timeout: time.Second * 100000,
		},
	}
}

func (adapter HTTPClientDomain) HandlerClient(req httpclient.HTTPClient) error {

	var resp interface{}
	b, _ := json.Marshal(req.Data)
	_, err := RequestUnmarshal(adapter.Client, req.Method, req.URL, string(b), buildAuthHeader(req.APIKey), &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp webhook : ", resp)
	return nil
}

func (adapter HTTPClientDomain) RegisterHandlerClient(req httpclient.HTTPClient) (string, error) {
	var resp response
	b, _ := json.Marshal(req.Data)
	_, err := RequestUnmarshal(adapter.Client, req.Method, req.URL, string(b), buildAuthHeader(req.APIKey), &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("resp webhook : ", resp)
	return resp.ID, nil
}

func (adapter HTTPClientDomain) NewHandlerClient(req httpclient.HTTPClient) (string, error) {
	var resp httpclient.NewResponse
	b, _ := json.Marshal(req.Data)
	_, err := RequestUnmarshal(adapter.Client, req.Method, req.URL, string(b), buildAuthHeader(req.APIKey), &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp : ", resp)
	return resp.ID, nil
}

func buildAuthHeader(apiKey string) map[string]string {
	h := map[string]string{
		"Authorization": apiKey,
		"Content-Type":  "application/json",
	}
	return h

}

func RequestUnmarshal(client *http.Client, method, url, body string, header map[string]string, v interface{}) (*http.Response, error) {

	br := strings.NewReader(body)
	req, err := http.NewRequest(method, url, br)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == 403 {
		fmt.Println(res.StatusCode)
		bodyBytes, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return nil, nil
	}

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		fmt.Println("new Decoder ", err)
		return res, err
	}
	// fmt.Println(v)
	return res, nil
}
