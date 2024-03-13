package requester

import (
	"fmt"
	"github.com/Waelson/go-temperatura-cep/internal/model"
	"io/ioutil"
	"net/http"
)

type HttpRequest interface {
	MakeRequest(url string) (string, int, error)
}

type httpRequest struct{}

func (h *httpRequest) MakeRequest(url string) (string, int, error) {
	resp, err := http.Get(url)
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
