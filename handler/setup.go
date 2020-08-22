package handler

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"
)

// DoSetup is setup alfred setting
// first arg is Docbase team
// second arg is Docbase secret token
func DoSetup(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	if err := alfred.SetDocbaseTeam(wf, args[0]); err != nil {
		return "", fmt.Errorf("please provide your docbase team name 👀")
	}

	if err := alfred.SetDocbaseSecret(wf, args[1]); err != nil {
		return "", fmt.Errorf("please provide your docbase token 👀")
	}

	return "Setup Completed 😉", nil
}
