package alfred

import (
	"fmt"
	"github.com/deanishe/awgo"
)

const (
	docbaseSecret = "docbase_secret"
)

func GetDocBaseSecret(wf *aw.Workflow) (string, error) {
	token, err := wf.Keychain.Get(docbaseSecret)
	if err != nil {
		return "", fmt.Errorf("token not found in the keychain")
	}

	return token, nil
}

func SetDocBaseSecret(wf *aw.Workflow, secret string) error {
	return wf.Keychain.Set(docbaseSecret, secret)
}

func RemoveToken(wf *aw.Workflow) error {
	return wf.Keychain.Delete(docbaseSecret)
}
