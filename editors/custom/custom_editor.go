package custom

import (
	"bufio"
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	editorUtils "github.com/MehrunesSky/gecrets/editors/utils"
	"github.com/MehrunesSky/gecrets/utils"
	"log"
	"strings"
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
	filepath := v.Write(secrets)
	v.exec(filepath)
}

func (v *Custom) Update(secrets []common.SecretI) []common.SecretI {
	filepath := v.Write(secrets)

	v.exec(filepath)

	return v.ReadSecrets(filepath)
}

func (v *Custom) ReadSecrets(filepath string) common.SecretIs {
	f, err := v.fileOpenerService.OpenFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	b := bufio.NewScanner(f)

	var pairs common.SecretIs
	for b.Scan() {
		if strings.HasPrefix(b.Text(), "#") {
			continue
		}
		nP := editorUtils.Unmarshal(v.model, b.Bytes())
		pairs = append(pairs, nP)
	}
	return pairs
}

func (v *Custom) Write(secrets common.SecretIs) string {
	s := strings.Builder{}

	s.WriteString(fmt.Sprintf("#" + v.model.ToJson() + "\n"))
	for _, p := range secrets {
		s.WriteString(fmt.Sprintf("%s\n", p.ToJson()))
	}

	f, err := v.fileService.CreateTempFile()

	if err != nil {
		log.Fatalln(err)
	}

	defer func(f utils.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	defer func(f utils.File) {
		err := f.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	_, err = f.WriteString(s.String())

	return f.Name()
}
