package current_user

import (
	"deltaboard-server/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type IdentityCheckResponse struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

func IdentityCheck(RealName, CardNo string) (bool, error) {

	data := url.Values{"realName": {RealName}, "cardNo": {CardNo}}
	req, err := http.NewRequest("POST", config.GetConfig().Identity.Url, strings.NewReader(data.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "APPCODE "+config.GetConfig().Identity.AppCode)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(string(content))
	}

	result := IdentityCheckResponse{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return false, err
	}

	if result.ErrorCode == 0 { // 身份信息一致
		return true, nil
	}

	return false, nil
}
