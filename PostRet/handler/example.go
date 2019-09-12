package handler

import (
	"context"

	example "MicroIhome/PostRet/proto/example"
	"github.com/astaxie/beego"
	"MicroIhome/IhomeWeb/utils"
	"time"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"MicroIhome/IhomeWeb/models"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

//加密函数
func Md5String(s string) string {
	//创建1个md5对象
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

func (e *Example) PostRet(ctx context.Context, req *example.Request, rsp *example.Response) error {

	beego.Info("PostRet  注册 /api/v1.0/users")

	rsp.Errno = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Errno)

	/*验证短信验证码*/

	//redis操作
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
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//通过手机号获取到短信验证码
	sms_code := bm.Get(req.Mobile)
	if sms_code == nil {
		beego.Info("获取数据失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//短信验证码对比
	sms_code_str, _ := redis.String(sms_code, nil)

	if sms_code_str != req.SmsCode {
		beego.Info("短信验证码错误")
		rsp.Errno = utils.RECODE_SMSERR
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*将数据存入数据库*/
	o := orm.NewOrm()
	user := models.User{Mobile: req.Mobile, Password_hash: Md5String(req.Password), Name: req.Mobile}

	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("注册数据失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	beego.Info("user_id", id)

	/*创建sessionid  （唯一的随即码）*/
	sessionid := Md5String(req.Mobile + req.Password)

	rsp.SessionId = sessionid

	/*以sessionid为key的一部分创建session*/
	//name //名字暂时使用手机号
	bm.Put(sessionid+"name", user.Mobile, time.Second*3600)
	//user_id
	bm.Put(sessionid+"user_id", id, time.Second*3600)
	//手机号
	bm.Put(sessionid+"mobile", user.Mobile, time.Second*3600)

	return nil
}