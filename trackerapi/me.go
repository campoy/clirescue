// This package provides an easy to use wrapper on the tracker API
package trackerapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL string = "https://www.pivotaltracker.com/services/v5/me"

var client http.Client

// APIToken returns the authentication token corresponding to the given
// username and password and an error if the operation fails.
func APIToken(usr, password string) (string, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("creating request")
	}
	req.SetBasicAuth(usr, password)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	var meResp struct {
		APIToken string `json:"api_token"`
		Error    string `json:"error"`
	}

	err = json.Unmarshal(body, &meResp)
	if err != nil {
		return "", fmt.Errorf("unmarshal response: %v", err)
	}
	// if there's API called failed, return the message as an error.
	if meResp.Error != "" {
		return "", fmt.Errorf(meResp.Error)
	}
	return meResp.APIToken, nil
}
