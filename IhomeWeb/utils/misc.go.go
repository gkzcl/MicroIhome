package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/weilaihui/fdfs_client"
	"fmt"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求

func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

func Md5String(s string) string {
	//创建1个md5对象
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

//上传二进制文件到fdfs中的操作
func UploadByBuffer(filebuffer []byte,fileExt string)(fileid string, err error){
	fd_cilent,err :=fdfs_client.NewFdfsClient("/home/roger/go/src/MicroIhome/IhomeWeb/conf/client.conf")
	if err!=nil{
		fmt.Println("创建句柄失败",err)
		fileid=""
		return
	}

	fd_rsq,err:=fd_cilent.UploadByBuffer(filebuffer,fileExt)
	if err!=nil{
		fmt.Println("上传失败",err)
		fileid=""
		return
	}

	fmt.Println(fd_rsq.GroupName)
	fmt.Println(fd_rsq.RemoteFileId)

	fileid = fd_rsq.RemoteFileId

	return fileid,nil
}