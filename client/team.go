package client

import "github.com/nextrevision/go-runscope"

// GetDefaultTeam returns the default team for the account
func (rc *RunscopeClient) GetDefaultTeam() (runscope.Team, error) {
	account, err := rc.Runscope.GetAccount()
	if err != nil {
		return runscope.Team{}, err
	}

	return account.Teams[0], nil
}
