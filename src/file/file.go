package file

import (
	"io/ioutil"
)

// Read ...
func Read(name string) (string, error) {
	bt, err := ioutil.ReadFile(name)
	if nil != err {
		return "", err
	}
	return string(bt), err
}

// Write ...
func Write(name string, context string) error {
	bt := []byte(context)
	return ioutil.WriteFile(name, bt, 0731)
}
