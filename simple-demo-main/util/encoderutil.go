package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CountBase64Val(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	io.Copy(h, f)
	re := h.Sum(nil) //算MD5值
	fmt.Printf("%x\n", re)
	mdHex := base64.StdEncoding.EncodeToString(h.Sum(nil)[:]) //MD5先转二进制数组再转base64编码
	fmt.Println(mdHex)
	f.Close()
	return mdHex

}

func GetMd5Val(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
