package handler

import (
	"context"

	example "MicroIhome/PostUserAuth/proto/example"
	"github.com/astaxie/beego"
	"MicroIhome/IhomeWeb/utils"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"reflect"
	"MicroIhome/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"time"

	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"strconv"
)

type Example struct{}

func (e *Example) PostUserAuth(ctx context.Context, req *example.Request, rsp *example.Response) error {

	//打印被调用的函数
	beego.Info(" 实名认证 Postuserauth  api/v1.0/user/auth ")
	//创建返回空间
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*从session中获取我们的user_id*/
	//构建连接缓存的数据
	redis_config_map := map[string]string{
		"key": utils.G_server_name,
		//"conn":"127.0.0.1:6379",
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	beego.Info(redis_config_map)
	redis_config, _ := json.Marshal(redis_config_map)

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

	value_id := bm.Get(sessioniduserid)
	beego.Info(value_id, reflect.TypeOf(value_id))
	id, _ := strconv.Atoi(string(value_id.([]uint8)[0]))

	//创建user对象
	user := models.User{Id: id,
		Real_name: req.RealName,
		Id_card: req.IdCard,
	}
	beego.Info(user)
	/*更新user表中的 姓名和 身份号*/
	o := orm.NewOrm()
	//更新表
	_, err = o.Update(&user, "real_name", "id_card")
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*更新我们的session中的user_id*/
	bm.Put(sessioniduserid, string(user.Id), time.Second*600)

	return nil
}
