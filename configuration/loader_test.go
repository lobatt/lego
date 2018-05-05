package configuration

import (
	"testing"
)

func TestLoadFromJson(t *testing.T) {
	input := `{
	"LocalFilePath": "/tmp/data/uploads",
	"OutputFilePath": "/tmp/data/artworks",
	"FSTDir" : "/tmp/fast-style-transfer"
} `

	config := &TransformerConfig{}
	if e := LoadFromJson([]byte(input), config); e != nil {
		t.Log(e)
		t.Fail()
	}

	if config.LocalFilePath != "/tmp/data/uploads" {
		t.Fail()
	}

	if config.OutputFilePath != "/tmp/data/artworks" {
		t.Fail()
	}
}

