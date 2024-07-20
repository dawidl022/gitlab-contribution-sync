# Gitlab Contribution Sync

A simple tool to sync your GitLab contribution count to your GitHub profile.
Only public data on your GitLab profile is used to determine the contribution
counts, so it is safe to sync private contributions.

Caveat: Due to the nature of the public GitLab API, only contributions from the
last 12 months can be synchronised.

## Usage

### Prerequisites

1. Create a blank private repo in your GitHub account.
2. Ensure you have [`git`](https://git-scm.com/) installed locally.
3. Ensure you are able to push to the private repo you created in step 1 (e.g.
   [via
   SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh)).
4. If you wish to sync private contribution counts from your GitLab profile too,
   ensure you have enabled the "Include private contributions on your profile"
   option in your GitLab profile settings.
   
### Installation

```
go install github.com/dawidl022/gitlab-contribution-sync@latest
```

### CLI

```bash
gitlab-contribution-sync <gitlab-username> <local-target-directory> <remote-github-repo>
```

where:

- `<gitlab-username>` is a placeholder for your GitLab username.
- `<local-target-directory>` is a placeholder for a path on your local
  filesystem, where the tool will create and update a local repository. This
  directory will be created upon first usage.
- `<remote-github-repo>` is a placeholder for the remote (push) address of the
  GitHub repo you want to sync your contributions to.

e.g.

```bash
gitlab-contribution-sync "dawidl022" "$HOME/Development/my-gitlab-contributions" "git@github.com:dawidl022/my-gitlab-contributions.git"
```

To re-sync your changes, simply re-run the command with the same arguments. You
can also set up a cron job to periodically run the tool.

## How it works

The tool uses the same API endpoint GitLab uses to render the "calendar" on a
user's profile: `https://gitlab.com/users/<username>/calendar.json`. For each
contribution count in the calendar, a commit is created in the private GitHub
repository with the correct date (but not time) of the GitLab contribution.

## Future development

- [go-git](https://github.com/go-git/go-git) could be used to move the file and
git operations into memory, in order to improve performance.

- descriptive error messages when something fails
