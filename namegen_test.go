package main

import (
	"fmt"
	"testing"
)

func TestGenerateNames(t *testing.T) {

	ctx := Context{}
	ctx.FilterConsonants = "fhjvqwxz"
	ctx.FilterVowels = "uyoi"
	ctx.Prefix = []string{"go"}
	ctx.Suffix = []string{"lo"}
	ctx.NameCount = 10
	ctx.Pattern = []string{PREFIX, SYLLABLE, SYLLABLE, SUFFIX}
	ctx.Dedup = false

	res := GenerateNames(&ctx)

	for _, name := range res {
		fmt.Println(name)
		// tests go here, for example check if filters, pattern, prefix, suffix, dedup works
	}
}

func TestFilterList(t *testing.T) {

	list := []rune{'q', 'r', 'd', 'f', 'g', 'ä'}
	filter := "qdä"

	filtered := filterList(list, filter)

	res_str := ""
	for _, c := range filtered {
		res_str += string(c)
	}

	expected := "rfg"

	if expected != res_str {
		t.Error("Expected ", expected, " got ", res_str)
	}

}
