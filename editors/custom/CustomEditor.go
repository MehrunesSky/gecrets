package custom

import (
	"github.com/MehrunesSky/gecrets/common"
	editorUtils "github.com/MehrunesSky/gecrets/editors/utils"
	"github.com/MehrunesSky/gecrets/utils"
	"log"
)

type Custom struct {
	cmd               string
	executor          utils.Executor
	model             common.SecretI
	fileService       utils.FileService
	fileOpenerService utils.FileOpenerService
}

func NewCustom(cmd string, model common.SecretI) *Custom {
	return &Custom{
		cmd:               cmd,
		model:             model,
		executor:          utils.OsExecutor{},
		fileService:       utils.OsFileService{},
		fileOpenerService: utils.OsFileService{},
	}
}

func (v *Custom) exec(filepath string) {
	err := v.executor.Execute(v.cmd, filepath)
	if err != nil {
		log.Fatalln(err)
	}
}

func (v *Custom) Open(secrets []common.SecretI) {
	filepath := editorUtils.WriteTempFile(v.fileService, v.model, secrets)
	v.exec(filepath)
}

func (v *Custom) Update(secrets []common.SecretI) []common.SecretI {
	filepath := editorUtils.WriteTempFile(v.fileService, v.model, secrets)

	v.exec(filepath)

	return editorUtils.ReadSecrets(v.fileOpenerService, v.model, filepath)

}
