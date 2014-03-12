package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/campoy/cliRescue/cmdutil"
	"github.com/campoy/clirescue/trackerapi"

	"github.com/codegangsta/cli"
)

var tokenFile = cachePath()

// cachePath returns the path for the cache file, relative to the user's home
// directory. It panics if the user's information is not available.
func cachePath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, ".tracker")
}

// getToken prints the authentication token for the user. It tries to find it
// in the cache file first, if not found it requests user and password to
// obtained from the tracker API.
func getToken(c *cli.Context) {
	// try to find token in file first.
	content, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		// only log an error if the file existed and couldn't read it
		if !os.IsNotExist(err) {
			log.Printf("load token: %v\n", content)
		}
	} else {
		log.Printf("token: %s", content)
		return
	}

	usr, err := cmdutil.ReadLine("Username")
	if err != nil {
		log.Fatalf("read username: %v", err)
	}
	pwd, err := cmdutil.ReadSilentLine("Password")
	if err != nil {
		log.Fatalf("read password: %v", err)
	}

	// get the tracker API auth token or panic if it fails.
	token, err := trackerapi.APIToken(usr, pwd)
	if err != nil {
		log.Fatalf("authentication: %v", err)
	}
	log.Printf("token: %s\n", token)

	err = ioutil.WriteFile(tokenFile, []byte(token), os.ModePerm)
	if err != nil {
		log.Fatalf("save token: %v", err)
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "clirescue"
	app.Usage = "CLI tool to talk to the Pivotal Tracker's API"

	app.Commands = []cli.Command{
		{
			Name:   "me",
			Usage:  "prints out Tracker's representation of your account",
			Action: getToken,
		},
	}

	app.Run(os.Args)
}
