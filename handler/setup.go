package handler

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"
	"log"
	"strings"
)

const (
	argsSeparator = " "
)

// DoSetup is setup alfred setting
// first arg is Client team
// second arg is Client secret token
func DoSetup(wf *aw.Workflow, args []string) (string, error) {
	log.Printf("args %+v", args)
	if len(args) != 1 {
		return "", fmt.Errorf("tt please provide some input 👀")
	}

	parts := strings.Split(args[0], argsSeparator)

	if err := alfred.SetDocBaseTeam(wf, parts[0]); err != nil {
		return "", fmt.Errorf("please provide your Client team name 👀")
	}

	if err := alfred.SetDocBaseSecret(wf, parts[1]); err != nil {
		return "", fmt.Errorf("please provide your Client token 👀")
	}

	return "Setup Completed 😉", nil
}
