package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBuildResponseNames(t *testing.T) {

	// Marshal
	respNames := buildRespNames()
	b, err := json.Marshal(&respNames)
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		panic(err)
	}

	// Unmarshal
	resp := RespNames{}
	err2 := json.Unmarshal(out.Bytes(), &resp)

	if err2 != nil {
		panic(err2)
	}

	for i, v := range resp.Names {
		if v != respNames.Names[i] {
			t.Error("Expected ", v, " got ", respNames.Names[i])
			break
		}
	}
}
