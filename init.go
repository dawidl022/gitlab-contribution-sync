package main

import (
	"os"

	"github.com/dawidl022/gitlab-contribution-sync/config"
	"github.com/dawidl022/gitlab-contribution-sync/contributions"
	"github.com/dawidl022/gitlab-contribution-sync/git"
)

// TODO check if target directory exists, has a contributions.json file and
// has git initialised, if not, mkdir -p and git init, and create empty {}
// contributions.json

// initialise the local repository if it doesn't exist
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
