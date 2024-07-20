package main

import (
	"github.com/dawidl022/gitlab-contribution-sync/config"
	"github.com/dawidl022/gitlab-contribution-sync/contributions"
	"github.com/dawidl022/gitlab-contribution-sync/git"
)

func syncContributions(config config.Config) error {
	syncedContributions, err := contributions.ReadSyncedContributions(config.TargetDir)
	if err != nil {
		return err
	}

	contribs, err := fetchGitlabContributions(config.GitlabUsername)
	if err != nil {
		return err
	}
	chronologicalContributions := contributions.SortContributions(contribs)

	for _, contrib := range chronologicalContributions {
		for syncedContributions[contrib.Date] < contrib.Count {
			syncedContributions[contrib.Date]++

			err := writeAndCommitContributions(syncedContributions, contrib.Date, config)
			if err != nil {
				return err
			}
		}
	}
	return git.PushContributions(config.TargetDir)
}

func writeAndCommitContributions(contribs map[contributions.Date]int, date contributions.Date, config config.Config) error {
	err := contributions.WriteSyncedContributions(contribs, config.TargetDir)
	if err != nil {
		return err
	}
	return git.CommitContributions(config.TargetDir, date, contribs[date])
}
