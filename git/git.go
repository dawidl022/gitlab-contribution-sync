package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dawidl022/gitlab-contribution-sync/config"
	"github.com/dawidl022/gitlab-contribution-sync/contributions"
)

// TODO each time the count is incremented, a commit is made with the
// message "Contribution <count> on <Date>", and is timestamped with midday
// local time on that Date (naive attempt at mitigating timezone issues).

func CommitContributions(targetDir string, date contributions.Date, count int) error {
	err := gitAdd(targetDir)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf("\"Contribution #%d on %s\"", count, date))
	cmd.Env = append(os.Environ())
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_AUTHOR_DATE=%sT12:00:00:00", date))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_COMMITTER_DATE=%sT12:00:00:00", date))
	cmd.Dir = targetDir

	return cmd.Run()
}

func gitAdd(targetDir string) error {
	cmd := exec.Command("git", "add", contributions.Filename)
	cmd.Dir = targetDir

	return cmd.Run()
}

func PushContributions(targetDir string) error {
	cmd := exec.Command("git", "push", "-u", "origin", "main")
	cmd.Dir = targetDir

	return cmd.Run()
}

func InitRepo(config config.Config) error {
	if !repoExists(config.TargetDir) {
		err := initGitRepo(config.TargetDir)
		if err != nil {
			return err
		}
		return addRemote(config)
	}
	return nil
}

func repoExists(targetDir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = targetDir

	err := cmd.Run()
	return err == nil
}

func initGitRepo(targetDir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = targetDir

	return cmd.Run()
}

func addRemote(config config.Config) error {
	cmd := exec.Command("git", "remote", "add", "origin", config.GithubRepo)
	cmd.Dir = config.TargetDir

	return cmd.Run()
}
