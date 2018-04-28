package tbot

import (
	"fmt"

	"github.com/fatih/structs"
)

// All Parameter we need in the tbot.config.json
type conf struct {
	twitterConsumerKey       string
	twitterConsumerSecret    string
	twitterAccessToken       string
	twitterAccessTokenSecret string
}

// StartConfig will ask to the user every parameter tbot need.
func StartConfig() {
	var c conf
	s := structs.New(c)

	for _, f := range s.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())
	}
}

/**
* PART : FILES
**/

func createFile() {

}

func writeInTheFile() {

}

func deleteFile() {

}
