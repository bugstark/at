package service

import (
	"context"
	"fmt"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	Accesskey string
	Secretkey string
	Bucket    string
}

// 七牛云上传
var UPLOAD *Qiniu

func InitQiNiu(ak, sk, bucket string) {
	UPLOAD = &Qiniu{Accesskey: ak, Secretkey: sk, Bucket: bucket}
}

func (q *Qiniu) GetUpToken() string {
	policy := storage.PutPolicy{
		Scope:      q.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","status":"ok"}`,
		Expires:    3600,
	}
	return policy.UploadToken(auth.New(q.Accesskey, q.Secretkey))
}

func (q *Qiniu) UpBase64(base, key string) error {
	cfg := storage.Config{
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	uploader := storage.NewBase64Uploader(&cfg)
	return uploader.Put(context.Background(), nil, q.GetUpToken(), key, []byte(base), nil)
}

func (q *Qiniu) UploadFile(file, key string) (err error) {
	cfg := storage.Config{
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutFile(context.Background(), &ret, q.GetUpToken(), key, file, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	return
}

func (q *Qiniu) GenFileLink(file, baseurl string) string {
	mac := auth.New(q.Accesskey, q.Secretkey)
	domain := baseurl
	key := file
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	return storage.MakePrivateURL(mac, domain, key, deadline)
}

func (q *Qiniu) FileStat(key string) (file *storage.FileInfo, err error) {
	bucketManager := storage.NewBucketManager(auth.New(q.Accesskey, q.Secretkey), &storage.Config{})
	fileInfo, err := bucketManager.Stat(q.Bucket, key)
	if err != nil {
		return
	}
	return &fileInfo, nil
}
