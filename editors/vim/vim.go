package vim

import (
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors/custom"
)

type VimExec struct {
	custom.Custom
}

func NewVimExec(model common.SecretI) *VimExec {
	return &VimExec{*custom.NewCustom(model, "vim")}
}
