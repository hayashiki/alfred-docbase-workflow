package handler

import (
	"context"
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/docbase"
	"strconv"
)

func DoSearch(wf *aw.Workflow, args []string) (string, error) {

	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	ctx := context.Background()
	docbaseService, err := docbase.NewClient(ctx, wf)
	items, err := docbaseService.List(args[0])

	if err != nil {
		return "", fmt.Errorf("fail to get docbase posts 👀")
	}

	icon := &aw.Icon{Value: "./icon/icon01.png"}
	for _, item := range items {
		wf.
			NewItem(item.Title).
			UID(strconv.Itoa(item.ID)).
			Subtitle(item.User.Name).
			Arg(item.URL).
			Valid(true).
			Icon(icon)
	}

	wf.WarnEmpty("No Posts were found.", "Try different query.")
	wf.SendFeedback()

	return "", nil
}