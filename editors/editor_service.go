package editors

import (
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors/idea"
	"github.com/MehrunesSky/gecrets/editors/vim"
)

type EditorService interface {
	Open(secrets []common.SecretI)
	Update(secrets []common.SecretI) []common.SecretI
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
