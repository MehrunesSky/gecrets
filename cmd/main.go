package main

import (
	"fmt"
	gsecrets "github.com/MehrunesSky/gsecrets"
)

func main() {
	v := gsecrets.NewVault("mehr")
	var pairs []gsecrets.Pair
	for _, value := range v.GetSecretIds() {
		secretValue := v.GetSecretValue(value)
		fmt.Println(value, secretValue)
		pairs = append(pairs, gsecrets.Pair{value, secretValue})
	}

	fmt.Println(pairs)
	nPairs := gsecrets.NewVimExec().UpdateWithVim(pairs)

	fmt.Println(nPairs)

	for _, p := range nPairs {
		v.SetSecretValue(p.A, p.B)
	}

}
