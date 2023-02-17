package service

import (
	"context"
	"errors"
	"fmt"
	"gingate/commons"
	"gingate/core"
	log "gingate/core"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func Upload2Minio(file *multipart.FileHeader, resname, ext, bucketname, objtype string) (string, error) {
	mc, err := core.GetMinioConn()
	if err != nil {
		return "", err
	}
	ct, err := core.ConvertExt2ContentType(ext)
	if err != nil {
		return "", err
	}
	err = checkFileSize(file)
	if err != nil {
		return "", err
	}
	err = checkFileExt(file, objtype)
	if err != nil {
		return "", err
	}
	fileContent, _ := file.Open()
	defer fileContent.Close()
	info, err := mc.PutObject(context.Background(), bucketname, resname, fileContent, file.Size, minio.PutObjectOptions{ContentType: ct})
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	return info.Key, nil
}

func checkFileSize(file *multipart.FileHeader) error {
	var limitedSize int = 1024 * 1024
	fileLimit := core.VOptions.GetInt("UploadProps.MAX_UPLOAD_FILE_SIZE") * limitedSize
	// not allow unlimited
	if fileLimit == 0 {
		fileLimit = core.DEFAULT_FILE_SIZE_LIMIT
	}
	if file.Size > int64(fileLimit) {
		return errors.New(fmt.Sprintf("%s:%dM", commons.CUS_ERR_4022, core.VOptions.GetInt("UploadProps.MAX_UPLOAD_FILE_SIZE")))
	}
	return nil
}

func checkFileExt(file *multipart.FileHeader, objtype string) error {
	multiExt := commons.IsMultiExt(file.Filename)
	if multiExt {
		return errors.New(commons.CUS_ERR_4102)
	}
	switch objtype {
	// 图片类型的
	case "avatar", "cover", "seal", "image", "images", "icon", "pic":
		isok := commons.IsImage(file.Filename)
		if !isok {
			return errors.New(commons.CUS_ERR_4021)
		}
	case "attachment":
		isok := commons.IsValidFile(file.Filename)
		if !isok {
			return errors.New(commons.CUS_ERR_4021)
		}
	default:
		return errors.New(commons.CUS_ERR_4100)
	}
	return nil
}
