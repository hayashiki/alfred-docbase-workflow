package docbase

import (
	"context"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"
	"github.com/hayashiki/docbase-go"
	"net/http"
)

const (
	offset = 1
	limit = 20
)

// DocBase describe an instance of API client.
type DocBase struct {
	Client *docbase.Client
}

func (c *DocBase) List(query string) ([]*docbase.Post, error) {

	opts := &docbase.PostListOptions{
		Q:       query,
		Page:    offset,
		PerPage: limit,
	}

	posts, resp, err := c.Client.Posts.List(opts)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return posts, err

}

// NewClient return a instance of DocBase client.
func NewClient(ctx context.Context, wf *aw.Workflow) (*DocBase, error) {
	team := alfred.GetDocBaseTeam(wf)
	token, err := alfred.GetDocBaseSecret(wf)
	if err != nil {
		return nil, err
	}

	return &DocBase{
		Client: docbase.NewClient(nil, team, token),
	}, nil
}
