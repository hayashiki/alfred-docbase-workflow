package handler

import (
	"fmt"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"

	aw "github.com/deanishe/awgo"
)

func DoCleanup(wf *aw.Workflow, _ []string) (string, error) {
	if _, err := alfred.GetDocBaseSecret(wf); err != nil {
		return "", fmt.Errorf("already cleanuped 🤣 (%w)", err)
	}

	if err := alfred.RemoveToken(wf); err != nil {
		return "", fmt.Errorf("already cleaned up 🤣 (%w)", err)
	}

	if err := alfred.RemoveDocBaseTeam(wf); err != nil {
		return "", fmt.Errorf("please provide your Client token 👀")
	}

	return "DocBase cleanup successfully 💪", nil
}
