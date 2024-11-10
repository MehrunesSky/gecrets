package editors

import (
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors/idea"
	"github.com/MehrunesSky/gecrets/editors/vim"
)

//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name EditorService --inpackage --inpackage-suffix
type EditorService interface {
	Open(secrets common.SecretIs)
	Update(secrets common.SecretIs) common.SecretIs
}

func GetEditorByName(name string, model common.SecretI) (EditorService, error) {
	var editor EditorService
	var err error
	switch name {
	case "idea":
		editor = idea.NewIdeaExec(model)
	case "vim":
		editor = vim.NewVimExec(model)
	default:
		err = fmt.Errorf("this editor %s doesn't exist", name)
	}
	return editor, err
}
