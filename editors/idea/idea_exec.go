package idea

import (
	"github.com/MehrunesSky/gecrets/common"
	editorUtils "github.com/MehrunesSky/gecrets/editors/utils"
	"log"
	"os"
	"os/exec"
)

type IdeaExec struct {
	filepath string
	model    common.SecretI
}

func NewIdeaExec(model common.SecretI) *IdeaExec {
	return &IdeaExec{model: model}
}

func (v *IdeaExec) exec() {
	cmd := exec.Command("idea --wait", v.filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func (v *IdeaExec) Open(secrets []common.SecretI) {
	v.filepath = editorUtils.WriteTempFile(v.model, secrets)
	v.exec()
}

func (v *IdeaExec) Update(secrets []common.SecretI) []common.SecretI {
	v.filepath = editorUtils.WriteTempFile(v.model, secrets)

	v.exec()

	return editorUtils.ReadSecrets(v.model, v.filepath)

}
