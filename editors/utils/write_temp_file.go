package utils

import (
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"log"
	"os"
	"strings"
)

func WriteTempFile(model common.SecretI, secrets []common.SecretI) string {
	s := strings.Builder{}

	s.WriteString(fmt.Sprintf("#" + model.ToJson() + "\n"))
	for _, p := range secrets {
		s.WriteString(fmt.Sprintf("%s\n", p.ToJson()))
	}

	f, err := os.CreateTemp("", "")

	if err != nil {
		log.Fatalln(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	defer func(f *os.File) {
		err := f.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	_, err = f.WriteString(s.String())

	return f.Name()
}
