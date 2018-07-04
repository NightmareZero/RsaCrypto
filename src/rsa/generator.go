package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// Generator ...
// RSA 生成 公钥 私钥
func Generator(bits int) error {
	privateKey, err := makePrivateKey(bits)
	if nil != err {
		return err
	}
	return makePublicKey(privateKey)
}

// 生成私钥
func makePrivateKey(bits int) (*rsa.PrivateKey, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	blocker := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// try create file
	file, err := os.Create("private.pem")
	if err != nil {
		return nil, err
	}

	err = pem.Encode(file, &blocker)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// 生成配对的公钥
func makePublicKey(privateKey *rsa.PrivateKey) error {
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	blocker := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, &blocker)
	if err != nil {
		return err
	}
	return nil
}
