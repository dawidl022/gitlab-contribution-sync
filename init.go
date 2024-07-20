package main

import (
	"os"

	"github.com/dawidl022/gitlab-contribution-sync/config"
	"github.com/dawidl022/gitlab-contribution-sync/contributions"
	"github.com/dawidl022/gitlab-contribution-sync/git"
)

// initialise the local repository if it doesn't exist, with a contributions.json file
func initLocalRepo(config config.Config) error {
	err := os.MkdirAll(config.TargetDir, 0755)
	if err != nil {
		return err
	}

	err = git.InitRepo(config)
	if err != nil {
		return err
	}

	return contributions.Init(config.TargetDir)
}
