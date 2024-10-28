package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func PromptYesNo(label string) bool {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" (Y/N) ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	result := strings.ToLower(strings.TrimSpace(s))
	return result == "y"
}
