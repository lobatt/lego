package fileutils

import (
	"encoding/json"
	"os"
)

//util funcs, to be re-organize to separate okg
func EnsureFilePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0777)
	}
	return nil
}

func SerializeToJson(data interface{}) []byte {
	s, e := json.Marshal(data)
	if e != nil {
		panic(e)
	}
	return s
}
