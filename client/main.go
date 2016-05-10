package client

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/nextrevision/go-runscope"
)

// RunscopeClient is
type RunscopeClient struct {
	Log      *log.Logger
	Runscope *runscope.Client
}

// NewRunscopeClient creates and returns a new RunscopeClient
func NewRunscopeClient(token string, debug bool, verbose bool) (*RunscopeClient, error) {
	rc := &RunscopeClient{
		Log: log.New(),
	}

	if debug {
		rc.Log.Level = log.DebugLevel
	} else if verbose {
		rc.Log.Level = log.InfoLevel
	} else {
		rc.Log.Level = log.WarnLevel
	}

	if token == "" {
		return rc, errors.New("must specify token")
	}

	rc.Runscope = runscope.NewClient(&runscope.Options{
		Token: token,
	})

	return rc, nil
}
