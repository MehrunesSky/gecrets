package utils

import (
	"bufio"
	"github.com/MehrunesSky/gecrets/common"
	"log"
	"os"
	"strings"
)

func ReadSecrets(model common.SecretI, filepath string) []common.SecretI {
	f, err := os.Open(filepath)
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
