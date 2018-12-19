package utils

import (
	"bytes"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

var sh *shell.Shell

func init() {
	sh = shell.NewShell("localhost:5001")
}

// 上传
func UploadIPFS(str string) (string, error) {
	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CatIPFS(hash string) (string, error) {

	read, err := sh.Cat(hash)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}
