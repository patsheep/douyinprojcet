package service

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dao"

	"github.com/RaymondCode/simple-demo/util"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"sync"
)


var videoBucket *oss.Bucket
var client *oss.Client
var coverBucket *oss.Bucket

func InitOss() {

	accessKey := config.CONFIG.OssConfig.Key
	accessSecret := config.CONFIG.OssConfig.Secret
	endpoint := config.CONFIG.OssConfig.Endpoint
	bucket :=config.CONFIG.OssConfig.Bucket
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	var err error
	client, err = oss.New(endpoint, accessKey, accessSecret)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("OSS连接失败")
		return
	}
	// 填写存储空间名称，例如examplebucket。
	videoBucket, err = client.Bucket(bucket)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("OSSBucket连接失败")
		return
	}

	client, err = oss.New(endpoint, accessKey, accessSecret)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	coverBucket, err = client.Bucket(bucket)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("已连接bucket" + videoBucket.BucketName)

}
func UploadFile(name string,id int64, wg *sync.WaitGroup) {

	//absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )
	//err := videoBucket.PutObjectFromFile("video/"+name, absPath+"\\video\\"+name)
	defer wg.Done()
	err := videoBucket.UploadFile("video/"+name, config.PROJECTPATH+config.VIDEO_ADDR+name, 100*1024, oss.Routines(10), oss.Checkpoint(true, ""), oss.ContentMD5(util.CountBase64Val(config.PROJECTPATH+config.VIDEO_ADDR+name))) //该方法支持断点续传，分片上传
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("上传OSS失败，文件将保留本地")
		dao.Db.Table("video").Where("id", id).Updates(map[string]interface{}{"play_url": config.PROJECTPATH+config.VIDEO_ADDR + name})
		return
	}
	err = videoBucket.SetObjectACL("video/"+name, oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = os.Remove(config.PROJECTPATH+config.VIDEO_ADDR+name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}
func UploadCover(name string,id int64, wg *sync.WaitGroup) {
	absPath, _ := os.Getwd()

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。

	//	fp :=absPath+name
	//strings.ReplaceAll(fp,"\\","\\\\")
	//err := bucket.PutObjectFromFile(name,absPath+"\\"+name )
	fmt.Println("cover/" + name + ".jpeg" + " " + absPath + "\\cover\\" + name + ".jpeg")
	err := coverBucket.PutObjectFromFile("cover/"+name+".jpeg", config.PROJECTPATH+config.COVER_ADDR+name+".jpeg")
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("封面上传失败，文件将保留本地")
		dao.Db.Table("video").Where("id", id).Updates(map[string]interface{}{"cover_url": config.PROJECTPATH+config.COVER_ADDR + name[0:len(name)-4] + ".jpeg"})
		return
	}
	err = coverBucket.SetObjectACL("cover/"+name+".jpeg", oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = os.Remove(config.PROJECTPATH+config.COVER_ADDR + name + ".jpeg")
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
