package main

import (
	"math/rand"
	"time"
)

/*
	Constants used in pattern definitions.
*/

const VOWEL string = "V"
const CONSONANT string = "C"
const SYLLABLE string = "S"
const PREFIX string = "PRE"
const SUFFIX string = "SUF"

/*
	Context holds the config and settings.
*/

type Context struct {
	vowels     []rune
	consonants []rune
	syllables  []string
	prefix     []string
	suffix     []string
	dedup      bool
}

/*
	Generates random company or product names with some predefined rules that can be specified.

	pattern:			The sequence of prefix, syllables, single chars and suffix.
	prefix:				Array of strings to use as prefix
	suffix:				Array of strings to use as suffix
	filter_consonants:	A string containing consonants to exclude in final name
	filter_vowels:		A string containing vowels to exclude in final name
	dedup:				True if double letters as aa, ee, uu should be removed
	name_count:			Number of names to generate
*/

func GenerateNames(
	pattern []string,
	prefix []string,
	suffix []string,
	filter_consonants string,
	filter_vowels string,
	dedup bool,
	name_count int) []string {

	ctx := Context{}
	config(&ctx, filter_consonants, filter_vowels, prefix, suffix, dedup)

	result := []string{}
	for i := 1; i <= name_count; i++ {
		result = append(result, getRandomName(pattern, &ctx))
	}
	return result
}

/*
	Apply filter on an array of runes.
*/

func filterList(list []rune, filter string) []rune {

	res := []rune{}
	for _, l := range list {
		l_match := false
		for _, f := range filter {
			if rune(f) == l {
				l_match = true
				break
			}
		}
		if !l_match {
			res = append(res, l)
		}
	}
	return res
}

/*
	Setup the configuration. Filter and exclude letters not to use etc.
	Generate syllables.
*/

func config(ctx *Context, filter_consonants string, filter_vowels string, prefix []string, suffix []string, dedup bool) {

	ctx.vowels = []rune{'a', 'e', 'i', 'o', 'u', 'y'}
	ctx.consonants = []rune{'a', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'z'}

	if filter_consonants != "" {
		c_filtered := filterList(ctx.consonants, filter_consonants)
		ctx.consonants = c_filtered
	}

	if filter_vowels != "" {
		v_filtered := filterList(ctx.vowels, filter_vowels)
		ctx.vowels = v_filtered
	}

	ctx.prefix = prefix
	ctx.suffix = suffix
	ctx.dedup = dedup

	generateSyllables(ctx)
}

/*
	Generate all combinations of vowel + consonant or vice versa.
*/

func generateSyllables(ctx *Context) {

	for _, c := range ctx.consonants {
		for _, v := range ctx.vowels {
			syl1 := string(c) + string(v)
			syl2 := string(v) + string(c)
			ctx.syllables = append(ctx.syllables, syl1)
			ctx.syllables = append(ctx.syllables, syl2)
		}
	}
}

/*
	Not used at the moment
	Use for wash of the names generated.
*/

func isVowel(c rune, ctx *Context) bool {

	for _, v := range ctx.vowels {
		if c == v {
			return true
		}
	}
	return false
}

func getRandomName(pattern []string, ctx *Context) string {

	name := ""
	for _, part_type := range pattern {
		name += string(getRandomPart(part_type, ctx))

	}

	// remove duplicate sequences like ee, aa etc.
	if ctx.dedup {
		deduped := ""
		i := 0
		for j := 0; j < len(name); j++ {
			if i < len(name)-1 {
				if name[i] != name[i+1] {
					deduped += string(name[i])
					i++
				} else {
					i++
				}
			}
		}
		name = deduped
	}

	return name
}

func getRandomPart(part_type string, ctx *Context) string {

	s := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(s)

	switch string(part_type) {
	case VOWEL:
		ix := r.Intn(len(ctx.vowels))
		return string(ctx.vowels[ix])
	case CONSONANT:
		ix := r.Intn(len(ctx.consonants))
		return string(ctx.consonants[ix])
	case PREFIX:
		ix := r.Intn(len(ctx.prefix))
		return string(ctx.prefix[ix])
	case SUFFIX:
		ix := r.Intn(len(ctx.suffix))
		return string(ctx.suffix[ix])
	default:
		ix := r.Intn(len(ctx.syllables))
		return string(ctx.syllables[ix])
	}
}
