package idea

import (
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors/custom"
)

type IdeaExec struct {
	custom.Custom
}

func NewIdeaExec(model common.SecretI) *IdeaExec {
	return &IdeaExec{*custom.NewCustom("idea --wait", model)}
}
