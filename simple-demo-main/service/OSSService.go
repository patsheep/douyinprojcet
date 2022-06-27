package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

	err := videoBucket.PutObjectFromFile("video/"+name, absPath+"\\video\\"+name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = videoBucket.SetObjectACL("video/"+name, oss.ACLPublicRead)
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
	defer wg.Done()
}

//直接上传文件，已弃用
func DirectUpload(name string) {
	absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )

	err := videoBucket.PutObjectFromFile("video/"+"x"+name, absPath+"\\video\\"+name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = videoBucket.SetObjectACL("video/"+"x"+name, oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}
