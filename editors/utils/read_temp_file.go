package utils

import (
	"bufio"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/utils"
	"log"
	"strings"
)

func ReadSecrets(readFileService utils.FileOpenerService, model common.SecretI, filepath string) []common.SecretI {
	f, err := readFileService.OpenFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	b := bufio.NewScanner(f)

	var pairs []common.SecretI
	for b.Scan() {
		if strings.HasPrefix(b.Text(), "#") {
			continue
		}
		nP := Unmarshal(model, b.Bytes())
		pairs = append(pairs, nP)
	}
	return pairs
}
