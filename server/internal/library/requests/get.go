package requests

import (
	"net/http"
	"time"
)

func Get(url string, response interface{}, headers http.Header, timeout int) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if headers != nil {
		req.Header = headers
	}

	if timeout != 0 {
		client.Timeout = time.Second * time.Duration(timeout)
	}

	return request(req, response)
}
