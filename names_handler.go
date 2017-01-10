package main

/*
 Just a sample implementation.
 All arguments are hard coded at this stage.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RespNames struct {
	Names []string
}

/*
	Marshalls the result to JSON
	TODO Implement handling of parameters used to control the generation on names
*/

func NamesHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("NamesHandler, name_count:" + r.URL.Query().Get("name_count"))

	respNames := buildRespNames()
	b, err := json.Marshal(&respNames)

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out.Bytes())
}

/*
	Calls the names generator and returns result
*/

func buildRespNames() RespNames {

	fmt.Println("buildResponse")

	filter_consonants := "fhjvqwxz"
	filter_vowels := "uyoi"
	prefix := []string{"go", "dat", "it"}
	suffix := []string{"lo", "mo"}
	name_count := 10
	pattern := []string{PREFIX, SYLLABLE, SYLLABLE, SUFFIX}
	dedup := true

	names := GenerateNames(
		pattern,
		prefix,
		suffix,
		filter_consonants,
		filter_vowels,
		dedup,
		name_count)

	r := RespNames{names}
	return r
}
