package handler

import (
	"context"

	example "MicroIhome/PostAvatar/proto/example"
	"strconv"
	"MicroIhome/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"MicroIhome/IhomeWeb/utils"
	"path"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {

	beego.Info("上传头像  PostAvatar /api/v1.0/user/avatar")

	/*初始化返回值*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	size := len(req.Avatar)

	//图片数据验证
	if req.Filesize != int64(size) {
		beego.Info("传输数据丢失")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*获取文件的后缀名*/
	//func Ext(path string) string

	//返回.jpg
	ext := path.Ext(req.Fileext)

	/*调用fdfs函数上传到图片服务器*/
	//路径..jpg
	fileid, err := utils.UploadByBuffer(req.Avatar, ext[1:])
	if err != nil {
		beego.Info("上传失败", err)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*得到fileid*/
	beego.Info(fileid)

	/*获取sessionid*/
	sessionid := req.SessionId

	/*连接redis*/
	//配置缓存参数
	redis_conf := map[string]string{
		"key": utils.G_server_name,
		//127.0.0.1:6379
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	beego.Info(redis_conf)

	//将map进行转化成为json
	redis_conf_js, _ := json.Marshal(redis_conf)

	//创建redis句柄
	bm, err := cache.NewCache("redis", string(redis_conf_js))
	if err != nil {
		beego.Info("redis连接失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*拼接key获取当前用户的user_id*/
	session_userid := sessionid + "user_id"
	user_id := bm.Get(session_userid)

	user_id_str, _ := redis.String(user_id, nil)

	id, _ := strconv.Atoi(user_id_str)

	/*将图片的存储地址（fileid）更新到user表中*/

	//创建user表对象
	user := models.User{Id: id, Avatar_url: fileid}
	//连接数据库
	o := orm.NewOrm()

	_, err = o.Update(&user, "avatar_url")
	if err != nil {

		beego.Info("数据更新失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}

	/*回传fielid*/
	rsp.AvatarUrl = fileid

	return nil
}
