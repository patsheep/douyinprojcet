package service

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"sync"
)

var accessKey, accessSecret string
var videoBucket *oss.Bucket
var client *oss.Client
var coverBucket *oss.Bucket

func OSSkeyinit() {

	var keyandpassword = dao.GetOSSKEy()
	accessKey = keyandpassword[0]
	accessSecret = keyandpassword[1]
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	var err error
	client, err = oss.New("https://oss-cn-shanghai.aliyuncs.com", accessKey, accessSecret)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	videoBucket, err = client.Bucket("pathcystore")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	client, err = oss.New("https://oss-cn-shanghai.aliyuncs.com", accessKey, accessSecret)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	coverBucket, err = client.Bucket("pathcystore")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("已连接bucket" + videoBucket.BucketName)

}
func UploadFile(name string, wg *sync.WaitGroup) {

	absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )
	//err := videoBucket.PutObjectFromFile("video/"+name, absPath+"\\video\\"+name)
	err := videoBucket.UploadFile("video/"+name, absPath+"\\video\\"+name, 100*1024, oss.Routines(10), oss.Checkpoint(true, ""), oss.ContentMD5(util.CountBase64Val(absPath+"\\video\\"+name))) //该方法支持断点续传，分片上传
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = videoBucket.SetObjectACL("video/"+name, oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = os.Remove(absPath + "\\video\\" + name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer wg.Done()
}
func UploadCover(name string, wg *sync.WaitGroup) {
	absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )
	fmt.Println("cover/" + name + ".jpeg" + " " + absPath + "\\cover\\" + name + ".jpeg")
	err := coverBucket.PutObjectFromFile("cover/"+name+".jpeg", absPath+"\\cover\\"+name+".jpeg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = coverBucket.SetObjectACL("cover/"+name+".jpeg", oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = os.Remove(absPath + "\\cover\\" + name + ".jpeg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer wg.Done()
}

//直接上传文件，已弃用
func DirectUpload(name string) {
	absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )

	/*	objectValue := "0123456789"

		mh := md5.Sum([]byte(objectValue))
		md5B64 := base64.StdEncoding.EncodeToString(mh[:])
		fmt.Println(md5B64)*/
	h := md5.New()
	f, err := os.Open(absPath + "\\video\\" + name)
	if err != nil {
		return
	}

	io.Copy(h, f)
	re := h.Sum(nil) //算MD5值
	fmt.Printf("%x\n", re)
	mdHex := base64.StdEncoding.EncodeToString(h.Sum(nil)[:]) //MD5先转二进制数组再转base64编码
	fmt.Println(mdHex)
	options := []oss.Option{
		oss.ContentMD5(mdHex),
		//B076A6BD67245533F9D1ACCF1112248C

	}
	err = videoBucket.PutObjectFromFile("video/"+"x"+name, absPath+"\\video\\"+name, options...)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	/*	err = videoBucket.SetObjectACL("video/"+"x"+name, oss.ACLPublicRead)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}*/

}
