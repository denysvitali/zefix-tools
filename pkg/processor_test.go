package processor_test

import (
	processor "github.com/denysvitali/documents-processor/pkg"
	"testing"
)

func TestProcessor(t *testing.T) {
	p, err := processor.New(
		"https://localhost:9200",
		"admin",
		"admin",
		true,
		"documents",
	)

	if err != nil {
		t.Fatal(err)
	}

	err = p.Process()
	if err != nil {
		t.Fatal(err)
	}
}
