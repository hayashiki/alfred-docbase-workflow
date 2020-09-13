package docbase

import (
	"context"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/alfred-docbase-workflow/alfred"
	"github.com/hayashiki/docbase-go"
)

// DocBase describe an instance of API client.
type DocBase struct {
	Client *docbase.Client
}

func (c *DocBase) List(query string) ([]docbase.Post, error) {
	// TODO: Flexible Query
	opts := &docbase.PostListOptions{
		Q:       "desc%3Ascore%20" + query,
		Page:    1,
		PerPage: 20,
	}

	posts, _, err := c.Client.Posts.List(opts)

	return posts.Posts, err

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
