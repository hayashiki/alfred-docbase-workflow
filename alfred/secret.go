package alfred

import (
	"fmt"
	aw "github.com/deanishe/awgo"
)

const (
	docbaseSecret = "docbase_secret"
)

func GetDocbaseSecret(wf *aw.Workflow) (string, error) {
	token, err := wf.Keychain.Get(docbaseSecret)
	if err != nil {
		return "", fmt.Errorf("token not found in the keychain")
	}

	return token, nil
}

func SetDocbaseSecret(wf *aw.Workflow, secret string) error {
	return wf.Keychain.Set(docbaseSecret, secret)
}
