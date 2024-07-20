package config

import "os"

type Config struct {
	GitlabUsername string
	TargetDir      string
	GithubRepo     string
}

func ReadConfig() (Config, error) {
	return Config{
		GitlabUsername: os.Args[1],
		TargetDir:      os.Args[2],
		GithubRepo:     os.Args[3],
	}, nil
}
