package vim

import (
	"github.com/MehrunesSky/gecrets/common"
	editorUtils "github.com/MehrunesSky/gecrets/editors/utils"
	"log"
	"os"
	"os/exec"
)

type VimExec struct {
	filepath string
	model    common.SecretI
}

func NewVimExec(model common.SecretI) *VimExec {
	return &VimExec{model: model}
}

func (v *VimExec) exec() {
	cmd := exec.Command("vim", v.filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func (v *VimExec) Open(secrets []common.SecretI) {
	v.filepath = editorUtils.WriteTempFile(v.model, secrets)
	v.exec()
}

func (v *VimExec) Update(secrets []common.SecretI) []common.SecretI {
	v.filepath = editorUtils.WriteTempFile(v.model, secrets)

	v.exec()

	return editorUtils.ReadSecrets(v.model, v.filepath)

}
