package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func Post(url string, jsonData interface{}, response interface{}, headers http.Header, timeout int) error {

	_d, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(_d))
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
