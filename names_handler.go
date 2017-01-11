package main

/*
 Just a sample implementation of a simple REST service.
 Arguments to name generation service is translated from JSON to a Context struct.
 JSON format:
 {
	"Pattern": ["C", S", "S", "S", "V"],
	"Prefix": ["go", "lo"],
	"Suffix": ["la", "na"],
	"FilterConsonants": "qxlvwyc",
	"FilterVowels": "ao",
	"Dedup": true,
	"NameCount": 10
}
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RespNames struct {
	Status string
	Names  []string
}

/*
	Marshalls the result to JSON
	TODO Implement handling of marshaling error
*/

func NamesHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("NamesHandler")

	respNames := RespNames{}
	var ctx Context
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ctx)
	if err != nil {
		respNames.Status = "error"
	} else {
		fmt.Println("ctx", ctx)
		respNames = buildRespNames(&ctx)
		fmt.Println("respNames", respNames)
	}

	b, err := json.Marshal(&respNames)
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	// TODO handle errors

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out.Bytes())
}

/*
	Calls the names generator and returns result
*/

func buildRespNames(ctx *Context) RespNames {

	fmt.Println("buildRespNames")

	names := GenerateNames(ctx)
	r := RespNames{"ok", names}

	return r
}
