package goconfluence

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// NewAPI implements API constructor
func NewAPI(location string, username string, token string) (*API, error) {
	if len(location) == 0 || len(username) == 0 || len(token) == 0 {
		return nil, errors.New("url, username or token empty")
	}

	u, err := url.ParseRequestURI(location)

	if err != nil {
		return nil, err
	}

	a := new(API)
	a.endPoint = u

	// Get the base url for confluence. It's whatever comes before `/rest`.
	base := location[:strings.Index(location, "/rest")]
	confBase, err := url.ParseRequestURI(base)
	if err != nil {
		return nil, err
	}
	a.confluenceBase = confBase
	a.token = token
	a.username = username
	a.client = &http.Client{}

	return a, nil
}

// DebugFlag is the global debugging variable
var DebugFlag = false

// SetDebug enables debug output
func SetDebug(state bool) {
	DebugFlag = state
}

// Debug outputs debug messages
func Debug(msg interface{}) {
	if DebugFlag {
		fmt.Printf("%+v\n", msg)
	}
}
