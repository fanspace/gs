package core

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetMinioConn() (*minio.Client, error) {
	minioClient, err := minio.New(Cfg.MinioSettings.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(Cfg.MinioSettings.KeyID, Cfg.MinioSettings.AccessKey, ""),
		Secure: false,
	})
	if err != nil {
		Error(err.Error())
		return nil, err
	}
	return minioClient, err
}
