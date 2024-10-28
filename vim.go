package gsecrets

import (
	"bufio"
	"fmt"
	"github.com/MehrunesSky/gsecrets/utils"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Pair struct {
	A, B string
}

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

func (v *VimExec) UpdateWithVim(pair []Pair) []Pair {
	s := strings.Builder{}

	for _, p := range pair {
		s.WriteString(fmt.Sprintf("%s=%s", p.A, p.B))
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

func (v *VimExec) read(oPair []Pair) []Pair {
	f, err := os.Open(v.filepath)
	if err != nil {
		log.Fatalln(err)
	}
	b := bufio.NewScanner(f)

	var re = regexp.MustCompile(`(?m)^\s*([\S]*)\s*=\s*([\S]*)\s*$`)
	var pairs []Pair
	for b.Scan() {
		for _, match := range re.FindAllStringSubmatch(b.Text(), -1) {
			nP := Pair{
				A: match[1],
				B: match[2],
			}
			if utils.NotContains(oPair, nP) {
				pairs = append(pairs, nP)
			}
		}
	}
	return pairs
}
