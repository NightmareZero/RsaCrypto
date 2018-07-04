package app

import (
	"encoding/base64"
	"file"
	"fmt"
	"log"
	"rsa"

	"github.com/farmerx/gorsa"
)

// Generate ...
// 生成Key逻辑
func Generate() error {
	err := rsa.Generator(1024)
	if nil != err {
		log.Fatalln(err)
		return err
	}
	return nil
}

func EncryptFile(name string) error {
	fileContext, err := file.Read(name)
	if nil != err {
		return err
	}

	strPubKey, err := file.Read("public.pem")
	if nil != err {
		return err
	}

	var RSA = gorsa.RSASecurity{}
	if err := RSA.SetPublicKey(strPubKey); err != nil {
		return err
	}

	pubenctypt, err := RSA.PubKeyENCTYPT([]byte(fileContext))
	if err != nil {
		return err
	}

	writeFinal := base64.StdEncoding.EncodeToString(pubenctypt)

	err = file.Write(name, writeFinal)
	if err != nil {
		return err
	}

	fmt.Println("ok!!!")

	return nil
}

func DecryptFile(name string) error {

	fileContext, err := file.Read(name)
	if nil != err {
		return err
	}

	fileVal, err := base64.StdEncoding.DecodeString(fileContext)
	if nil != err {
		return err
	}

	strPriKey, err := file.Read("private.pem")
	if nil != err {
		return err
	}

	var RSA = gorsa.RSASecurity{}
	if err := RSA.SetPrivateKey(strPriKey); err != nil {
		return err
	}

	pridecrypt, err := RSA.PriKeyDECRYPT(fileVal)
	if err != nil {
		return err
	}

	err = file.Write(name, string(pridecrypt))
	if err != nil {
		return err
	}

	fmt.Println("ok!!!")

	return nil
}
