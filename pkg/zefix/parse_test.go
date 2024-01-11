package zefix_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/denysvitali/sparql-client"
)

func TestParse(t *testing.T) {
	f, err := os.Open("../../zefix.json")
	if err != nil {
		t.Fatalf("unable to open file: %v", err)
	}

	var result sparql.Result
	dec := json.NewDecoder(f)
	err = dec.Decode(&result)
	if err != nil {
		t.Fatalf("unable to decode file: %v", err)
	}

	for _, v := range result.Results.Bindings {
		if v["name"].Value == "KPT Assicurazioni SA" {
			t.Logf("Found: %v", v)
		}
	}
}
