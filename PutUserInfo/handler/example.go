package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"reflect"
	"github.com/astaxie/beego/orm"
	"time"
	"MicroIhome/IhomeWeb/utils"
	"MicroIhome/IhomeWeb/models"
	example "MicroIhome/PutUserInfo/proto/example"
	"strconv"
)

type Example struct{}

func (e *Example) PutUserInfo(ctx context.Context, req *example.Request, rsp *example.Response) error {

	//打印被调用的函数
	beego.Info("---------------- PUT  /api/v1.0/user/name PutUersinfo() ------------------")

	//创建返回空间
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*得到用户发送过来的name*/
	beego.Info(rsp.Username)

	/*从从sessionid获取当前的userid*/
	//连接redis
	redis_config_map := map[string]string{
		"key": utils.G_server_name,
		//"conn":"127.0.0.1:6379",
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	beego.Info(redis_config_map)
	redis_config, _ := json.Marshal(redis_config_map)
	beego.Info(string(redis_config))

	//连接redis数据库 创建句柄
	bm, err := cache.NewCache("redis", string(redis_config))
	if err != nil {
		beego.Info("缓存创建失败", err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//拼接key
	sessioniduserid := req.Sessionid + "user_id"
	//获取userid
	value_id := bm.Get(sessioniduserid)
	beego.Info(value_id, reflect.TypeOf(value_id))

	id ,_:= strconv.Atoi(string(value_id.([]uint8)[0]))
	beego.Info(id, reflect.TypeOf(id))

	//创建表对象
	user := models.User{Id: id, Name: req.Username}
	/*更新对应user_id的name字段的内容*/
	//创建数据库句柄
	o := orm.NewOrm()
	//更新
	_, err = o.Update(&user, "name")
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)

		return nil
	}

	/*更新session user_id*/
	sessionidname := req.Sessionid + "name"
	bm.Put(sessioniduserid, string(user.Id), time.Second*600)
	/*更新session name*/
	bm.Put(sessionidname, string(user.Name), time.Second*600)

	/*成功返回数据*/
	rsp.Username = user.Name
	return nil
}
