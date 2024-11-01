package custom

import (
	"github.com/MehrunesSky/gecrets/common"
	editorUtils "github.com/MehrunesSky/gecrets/editors/utils"
	"log"
	"os"
	"os/exec"
)

type Custom struct {
	cmd   string
	model common.SecretI
}

func NewCustom(cmd string, model common.SecretI) *Custom {
	return &Custom{cmd: cmd, model: model}
}

func (v *Custom) exec(filepath string) {
	cmd := exec.Command(v.cmd, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func (v *Custom) Open(secrets []common.SecretI) {
	filepath := editorUtils.WriteTempFile(v.model, secrets)
	v.exec(filepath)
}

func (v *Custom) Update(secrets []common.SecretI) []common.SecretI {
	filepath := editorUtils.WriteTempFile(v.model, secrets)

	v.exec(filepath)

	return editorUtils.ReadSecrets(v.model, filepath)

}
