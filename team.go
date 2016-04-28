package main

import "github.com/nextrevision/go-runscope"

func getDefaultTeam() (*runscope.Team, error) {
	account, err := client.GetAccount()
	if err != nil {
		return &runscope.Team{}, err
	}

	return &account.Teams[0], nil
}
