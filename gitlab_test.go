package main

import (
	"fmt"
	"testing"
)

func TestFetchGitlabContributions(t *testing.T) {
	contributions, err := fetchGitlabContributions("dawidl022")
	if err != nil {
		panic(err)
	}
	fmt.Println(contributions)
}
