package test

import (
	"fmt"
	"lovehome_beego/utils"
	"testing"
)

func TestUpload(t *testing.T) {
	hash, err := utils.UploadIPFS("liguoyu3564liguoyu3564")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
}
