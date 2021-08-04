package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 30,
	}
}

func request(req *http.Request, response interface{}) error {

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("request error with code: %v", resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	return nil
}
