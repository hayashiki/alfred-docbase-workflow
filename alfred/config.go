package alfred

import aw "github.com/deanishe/awgo"

const (
	docbaseTeam = "docbase_team"
)

func GetDocbaseTeam(wf *aw.Workflow) string {
	return wf.Config.Get(docbaseTeam)
}

func SetDocbaseTeam(wf *aw.Workflow, id string) error {
	return wf.Config.Set(docbaseTeam, id, false).Do()
}
