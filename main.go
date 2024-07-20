package main

import (
	"log"

	"github.com/dawidl022/gitlab-contribution-sync/config"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = initLocalRepo(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = syncContributions(conf)
	if err != nil {
		log.Fatal(err)
	}
}
