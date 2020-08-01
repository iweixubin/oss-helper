package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"testing"
)

func Test_MinIo_MakeBucket(t *testing.T) {
	client := getMinIoClient()
	err := client.MakeBucket(context.Background(), "img", minio.MakeBucketOptions{
		Region:        "", // 数据中心的物理位置信息。
		ObjectLocking: false,
	})
	if err != nil {
		t.Error(err)
	}

	exist, err := client.BucketExists(context.Background(), "img")
	if err != nil {
		t.Error(err)
	}

	if !exist {

	}

	//client.lo
	//
	//ioutil.ReadFile("")
	//
	//bytes.NewReader()
	//io.Reader
	//ioutil.ReadFile()
	//
	//client.PutObject()

}

func Test_MinIo_PutObject(t *testing.T) {
	filePath := `F:\Docs\oss-helper\file_sniffer\samples\80px-JPEG_example_JPG_RIP_100.jpg`
	filename := genFilename("jpg")
	objectName := "x02/" + filename
	contentType := "image/jpeg"

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
		return
	}

	r := bytes.NewReader(fileBytes)

	info, err := getMinIoClient().PutObject(context.Background(), "img", objectName, r, int64(len(fileBytes)), minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%+v", info)
	}
}

func Test_MinIo_FPutObject(t *testing.T) {
	filePath := `F:\Docs\oss-helper\file_sniffer\samples\PNG_transparency_demonstration_1.png`
	filename := genFilename("png")
	objectName := "x01/" + filename

	contentType := "image/png"

	info, err := getMinIoClient().FPutObject(context.Background(), "img", objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%+v", info)
	}

}

func Test_genFilename(t *testing.T) {
	println(genFilename(".png"))
}
