package requester

import (
	"io/ioutil"
	"net/http"
)

type HttpRequest interface {
	MakeRequest(url string) (string, error)
}

type httpRequest struct{}

func (h *httpRequest) MakeRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func NewHttpRequest() HttpRequest {
	return &httpRequest{}
}
