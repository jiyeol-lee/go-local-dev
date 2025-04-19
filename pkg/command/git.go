package command

import (
	"encoding/json"
	"os/exec"
)

type BranchSyncStatus struct {
	Behind int `json:"behind"`
	Ahead  int `json:"ahead"`
}

func GetCurrentBranch(dir string) (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	branch := string(output)
	branch = branch[:len(branch)-1] // Remove the trailing newline character
	return branch, nil
}

func GetBranchSyncStatus(dir string) (*BranchSyncStatus, error) {
	cmd := exec.Command(
		"sh",
		"-c",
		`git rev-list --left-right --count origin/$(git rev-parse --abbrev-ref HEAD)...HEAD | awk '{printf("{ \"behind\": %s, \"ahead\": %s }\n", $1, $2)}'`,
	)
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var status BranchSyncStatus
	err = json.Unmarshal(output, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
