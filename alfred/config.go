package alfred

import aw "github.com/deanishe/awgo"

const (
	DocBaseTeam = "docbase_team"
)

func GetDocBaseTeam(wf *aw.Workflow) string {
	return wf.Config.Get(DocBaseTeam)
}

func SetDocBaseTeam(wf *aw.Workflow, id string) error {
	return wf.Config.Set(DocBaseTeam, id, false).Do()
}

func RemoveDocBaseTeam(wf *aw.Workflow) error {
	return wf.Config.Unset(DocBaseTeam).Do()
}
