package alfred

import (
	"fmt"
	"github.com/deanishe/awgo"
)

const (
	docBaseSecret = "docbase_secret"
)

func GetDocBaseSecret(wf *aw.Workflow) (string, error) {
	token, err := wf.Keychain.Get(docBaseSecret)
	if err != nil {
		return "", fmt.Errorf("token not found in the keychain")
	}

	return token, nil
}

func SetDocBaseSecret(wf *aw.Workflow, secret string) error {
	return wf.Keychain.Set(docBaseSecret, secret)
}

func RemoveToken(wf *aw.Workflow) error {
	return wf.Keychain.Delete(docBaseSecret)
}
