package configuration

import (
	"testing"
)

type TestConfig struct {
	LocalFilePath  string
	OutputFilePath string
}

func TestLoadFromJson(t *testing.T) {
	input := `{
	"LocalFilePath": "/tmp/data/uploads",
	"OutputFilePath": "/tmp/data/artworks",
	"FSTDir" : "/tmp/fast-style-transfer"
} `

	config := &TestConfig{}
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
