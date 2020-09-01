package handler

import (
	"fmt"
	"github.com/deanishe/awgo"
)

func DoInstall(wf *aw.Workflow, _ []string) (string, error) {
	fmt.Print("Downloading install...")

	if err := wf.InstallUpdate(); err != nil {
		return "", fmt.Errorf("fail to install err: %w", err)
	}

	return "", nil
}
