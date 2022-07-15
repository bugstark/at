package service

import (
	"context"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	Accesskey string
	Secretkey string
	Bucket    string
}

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
