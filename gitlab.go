package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dawidl022/gitlab-contribution-sync/contributions"
)

// get all the contributions from the gitlab user via API
func fetchGitlabContributions(username string) (map[contributions.Date]int, error) {
	var contribs map[contributions.Date]int

	resp, err := http.Get(fmt.Sprintf("https://gitlab.com/users/%s/calendar.json", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&contribs)
	return contribs, err
}
