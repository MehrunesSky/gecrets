package vim

import (
	"bufio"
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/utils"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type VimExec struct {
	filepath string
}

func NewVimExec() *VimExec {
	return new(VimExec)
}

func (v *VimExec) writeToTemp(s string) {
	f, err := os.OpenFile(v.filepath, os.O_APPEND|os.O_WRONLY, 0644)
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

	_, err = f.WriteString(s)

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

func (v *VimExec) OpenWithVim(pair []common.Secret) {
	s := strings.Builder{}

	for _, p := range pair {
		s.WriteString(fmt.Sprintf("%s=%s\n", p.Key, p.Value))
	}

	f, err := os.CreateTemp("", "")

	if err != nil {
		log.Fatalln(err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatalln(err)
		}
	}(f.Name())

	v.filepath = f.Name()
	v.writeToTemp(s.String())
	v.exec()
	v.writeToTemp(s.String())
}

func (v *VimExec) UpdateWithVim(pair []common.Secret) []common.Secret {
	s := strings.Builder{}

	for _, p := range pair {
		s.WriteString(fmt.Sprintf("%s=%s\n", p.Key, p.Value))
	}

	f, err := os.CreateTemp("", "")

	if err != nil {
		log.Fatalln(err)
	}

	v.filepath = f.Name()

	v.writeToTemp(s.String())
	v.exec()
	return v.read(pair)

}

func (v *VimExec) read(oPair []common.Secret) []common.Secret {
	f, err := os.Open(v.filepath)
	if err != nil {
		log.Fatalln(err)
	}
	b := bufio.NewScanner(f)

	var re = regexp.MustCompile(`(?m)^\s*([\S]*)\s*=\s*([\S]*)\s*$`)
	var pairs []common.Secret
	for b.Scan() {
		for _, match := range re.FindAllStringSubmatch(b.Text(), -1) {
			nP := common.Secret{
				Key:   match[1],
				Value: match[2],
			}
			if utils.NotContains(oPair, nP) {
				pairs = append(pairs, nP)
			}
		}
	}
	return pairs
}
