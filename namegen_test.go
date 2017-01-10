package main

import (
	"fmt"
	"testing"
)

func TestGenerateNames(t *testing.T) {

	filter_consonants := "fhjvqwxz"
	filter_vowels := "uyoi"
	prefix := []string{"go", "dat", "it"}
	suffix := []string{"lo", "mo"}
	name_count := 10
	pattern := []string{PREFIX, SYLLABLE, SYLLABLE, SUFFIX}
	dedup := false

	res := GenerateNames(pattern, prefix, suffix, filter_consonants, filter_vowels, dedup, name_count)

	for _, name := range res {
		fmt.Println(name)
		// tests go here, for example check if filters, pattern, dedup works
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
