package grpc_service

import (
	"context"
	"fmt"
	"oss-helper/grpc_service/contract"
)

type StorageService struct {
}

func (ss *StorageService) Upload(ctx context.Context, req *contract.UploadRequest) (resp *contract.UploadResponse, err error) {
	fmt.Printf("%+v", req)
	c := &contract.UploadResponse{}
	c.BucketName = req.Expired + "_go"
	c.ObjectName = req.FileName + "_go"

	// 返回 nil 会报错
	return c, nil //errors.New("华丽丽的错误~") //返回错误，那么 C# 是抛异常
}
