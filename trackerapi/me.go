package trackerapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	u "os/user"
	"path/filepath"

	"github.com/campoy/clirescue/cmdutil"
)

const URL string = "https://www.pivotaltracker.com/services/v5/me"

var FileLocation string = fromHome("/.tracker")

func Me() error {
	u, p, err := getCredentials()
	if err != nil {
		return err
	}

	token, err := getAPIToken(u, p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FileLocation, []byte(token), 0644)
}

func getAPIToken(usr, password string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(usr, password)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// TODO here? really?
	fmt.Printf("\n****\nAPI response: \n%s\n", string(body))

	var meResp struct {
		APIToken string `json:"api_token"`
	}

	err = json.Unmarshal(body, &meResp)
	return meResp.APIToken, err
}

func getCredentials() (usr, pwd string, err error) {
	fmt.Print("Username: ")
	usr, err = cmdutil.ReadLine()
	if err != nil {
		return
	}

	cmdutil.Silence()
	defer cmdutil.Unsilence()

	fmt.Print("Password: ")
	pwd, err = cmdutil.ReadLine()

	return
}

func fromHome(path string) string {
	usr, err := u.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path)
}
