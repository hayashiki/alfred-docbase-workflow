package handler

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"
	"strings"
)

const (
	argsSeparator = " "
)

// DoSetup is setup DocBase setting
// First split arg is team name
// Second split arg is token
func DoSetup(wf *aw.Workflow, arg string) (string, error) {
	if arg == "" {
		return "", fmt.Errorf("please provide team or token input 😂")
	}

	parts := strings.Split(arg, argsSeparator)

	if err := alfred.SetDocBaseTeam(wf, parts[0]); err != nil {
		return "", fmt.Errorf("please provide your team name 😂")
	}

	if err := alfred.SetDocBaseSecret(wf, parts[1]); err != nil {
		return "", fmt.Errorf("please provide your token 😂")
	}

	return "Setup Completed 😉", nil
}
