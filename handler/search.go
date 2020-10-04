package handler

import (
	"context"
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/docbase"
	"strconv"
)

// DoSearch is search DocBase posts by keyword
func DoSearch(wf *aw.Workflow, arg string) (string, error) {
	if arg == "" {
		return "", fmt.Errorf("please provide team or token input 😂")
	}

	ctx := context.Background()
	docBaseService, err := docbase.NewClient(ctx, wf)
	items, err := docBaseService.List(arg)

	if err != nil {
		return "", fmt.Errorf("fail to get docbase posts 😂")
	}

	for _, item := range items {
		wf.
			NewItem(item.Title).
			UID(strconv.Itoa(item.ID)).
			Subtitle(item.User.Name).
			Arg(item.URL).
			Valid(true).
			Icon(&aw.Icon{Value: "./icon/icon01.png"})
	}

	wf.WarnEmpty("No Posts were found.", "Try different query.")
	wf.SendFeedback()

	return "", nil
}
