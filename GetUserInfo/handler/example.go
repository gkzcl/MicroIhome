package handler

import (
	"context"

	example "MicroIhome/GetUserInfo/proto/example"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"MicroIhome/IhomeWeb/utils"
	"encoding/json"
	"strconv"
	"MicroIhome/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

func (e *Example) GetUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {

	beego.Info("获取用户信息 GetUserInfo /api/v1.0/user ")

	/*初始化错误码*/
	rsp.Errno = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Errno)

	/*获取sessionid*/
	sessionid := req.Sessionid

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
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
	}

	/*拼接key*/
	sessionuserid := sessionid + "user_id"

	/*通过key获取到user_id*/
	user_id := bm.Get(sessionuserid)
	beego.Info(user_id)
	beego.Info(sessionuserid)
	//beego.Info(reflect.TypeOf(user_id),user_id)
	//beego.Info(reflect.TypeOf(userid_str),userid_str)
	//user_id ,_:= strconv.Atoi(string(value_id.([]uint8)[0]))
	id ,_:= strconv.Atoi(string(user_id.([]uint8)[0]))
	beego.Info(id)

	/*通过user_id获取到用户表信息*/
	//创建1个user对象
	user := models.User{Id: id}
	//创建orm句柄
	o := orm.NewOrm()
	err = o.Read(&user)
	if err != nil {
		beego.Info("数据获取失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Errno)
	}

	beego.Info(user)

	/*将信息返回*/
	rsp.UserId = strconv.Itoa(user.Id)
	rsp.Name = user.Name
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.Mobile = user.Mobile
	rsp.AvatarUrl = user.Avatar_url

	return nil
}
