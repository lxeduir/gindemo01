package public

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var accessKey = "fM089N-8DRYM2nQb4g9DD-uRgjzktZ1WABbTAlGg"
var secretKey = "muJ7MCqZWWC_h-5dc-HH2pxQrUvmo12KycXjEymT"
var bucket = "img-blog-1"

func Upload(file string, filename string, params map[string]string) (storage.PutRet, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: params,
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, "gxa/"+filename, file, &putExtra)
	return ret, err
}
