package main

func main() {
	// TODO read 3 CLI args (gitlab username, target directory, github repo
	// address)

	// TODO check if target directory exists, has a contributions.json file and
	// has git initialised, if not, mkdir -p and git init, and create empty {}
	// contributions.json

	// TODO read contributions.json file and parse it into a map

	// TODO get all the contributions from the gitlab user via API, and iterate
	// over dates in chronological order

	// TODO for each date, iterate from the count in the map to the current
	// count. Each time the count is incremented in the map, and the map is
	// written to the contributions.json file in chronological order.

	// TODO each time the count is incremented, a commit is made with the
	// message "Contribution <count> on <date>", and is timestamped with midday
	// local time on that date (naive attempt at mitigating timezone issues).

	// TODO the readme should mention that the binary should be installed,
	// and a cron job can be set up to run the binary daily.
}
