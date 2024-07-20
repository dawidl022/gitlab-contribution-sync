package contributions

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"sort"
)

type Date string

type Contributions struct {
	Date  Date
	Count int
}

// sort contributions chronologically
func SortContributions(contributions map[Date]int) []Contributions {
	chronologicalContributions := make([]Contributions, 0, len(contributions))
	for date, count := range contributions {
		chronologicalContributions = append(chronologicalContributions, Contributions{date, count})
	}
	sort.Slice(chronologicalContributions, func(i, j int) bool {
		return chronologicalContributions[i].Date < chronologicalContributions[j].Date
	})
	return chronologicalContributions
}

const Filename = "contributions.json"

func Init(targetDir string) error {
	_, err := os.Stat(contributionsPath(targetDir))
	if errors.Is(err, os.ErrNotExist) {
		return os.WriteFile(contributionsPath(targetDir), []byte("{}"), 0644)
	}
	return err
}

func contributionsPath(targetDir string) string {
	return path.Join(targetDir, Filename)
}

// read contributions.json file and parse it into a map
func ReadSyncedContributions(targetDir string) (map[Date]int, error) {
	fileBytes, err := os.ReadFile(contributionsPath(targetDir))
	if err != nil {
		return nil, err
	}

	var contributions map[Date]int
	err = json.Unmarshal(fileBytes, &contributions)

	return contributions, err
}

// write the contributions map to contributions.json file in chronological order
func WriteSyncedContributions(contributions map[Date]int, targetDir string) error {
	jsonContributions, err := json.MarshalIndent(contributions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(contributionsPath(targetDir), jsonContributions, 0644)
}
