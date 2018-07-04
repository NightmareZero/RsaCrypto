package test

import (
	"rsa"
	"testing"
)

func TestGenerator(t *testing.T) {
	//rsa 密钥文件产生
	rsa.Generator(1024)
}
