package requester

import (
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpRequest interface {
	MakeRequest(url string) (string, int, error)
	Normalize(str string) (string, error)
}

type httpRequest struct{}

func (h *httpRequest) Normalize(str string) (string, error) {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	parsedURL.RawQuery = url.QueryEscape(parsedURL.Query().Encode())
	return parsedURL.String(), nil
}

func (h *httpRequest) MakeRequest(urlStr string) (string, int, error) {

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, model.InternalError
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return string(body), http.StatusOK, nil
}

func NewHttpRequest() HttpRequest {
	return &httpRequest{}
}
