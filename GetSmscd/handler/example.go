package handler

import (
	"context"

	example "MicroIhome/GetSmscd/proto/example"

	"github.com/astaxie/beego/cache"
	"github.com/garyburd/redigo/redis"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"

	"time"
	"github.com/astaxie/beego"
	"MicroIhome/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"MicroIhome/IhomeWeb/models"
	"encoding/json"
	"math/rand"
)

type Example struct{}

func (e *Example) GetSmscd(ctx context.Context, req *example.Request, rsp *example.Response) error {

	beego.Info("获取短信验证码 GetSmscd /api/v1.0/smscode/:mobile ")

	//初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)

	/*验证手机号是否存在*/
	//创建数据库orm句柄
	o := orm.NewOrm()
	//使用手机号作为查询条件
	user := models.User{Mobile: req.Mobile}

	err := o.Read(&user)
	//如果不报错就说明查找到了
	//查找到就说明手机号存在
	if err == nil {
		beego.Info("用户已存在")
		rsp.Error = utils.RECODE_MOBILEERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*验证图片验证码是否正确*/
	//连接redis
	//配置缓存参数
	redis_conf := map[string]string{
		"key": utils.G_server_name,
		//127.0.0.1:6379
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}

	//将map进行转化成为json
	redis_conf_js, _ := json.Marshal(redis_conf)

	//创建redis句柄
	bm, err := cache.NewCache("redis", string(redis_conf_js))
	if err != nil {
		beego.Info("redis连接失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//通过uuid查找图片验证码的值进行对比
	uuid := bm.Get(req.Uuid)
	if uuid == nil {
		beego.Info("redis获取失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//reflect.TypeOf(value)会返回当前数据的变量类型
	//beego.Info(reflect.TypeOf(value), value)
	//格式转换
	uuid_str, _ := redis.String(uuid, nil)

	if uuid_str != req.Imagestr {
		beego.Info("数据不匹配 图片验证码值错误")
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*调用 短信接口发送短信*/
	//创建随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(9999) + 1001
	beego.Info("验证码", size)
	////发送短信的配置信息
	//messageconfig := make(map[string]string)
	////预先创建好的appid
	//messageconfig["appid"] = "29672"
	////预先获得的app的ket
	//messageconfig["appkey"] = "89d90165cbea8cae80137d7584179bdb"
	////加密方式默认
	//messageconfig["signtype"] = "md5"
	//
	////messagexsend
	////创建短信发送的句柄
	//messagexsend := submail.CreateMessageXSend()
	////短信发送的手机号
	//submail.MessageXSendAddTo(messagexsend, req.Mobile)
	////短信发送的模板
	//submail.MessageXSendSetProject(messagexsend, "NQ1J94")
	////验证码
	//submail.MessageXSendAddVar(messagexsend, "code", strconv.Itoa(size))
	////发送短信的请求
	//fmt.Println("MessageXSend ", submail.MessageXSendRun(submail.MessageXSendBuildRequest(messagexsend), messageconfig))

	/*将短信验证码存入缓存数据库*/
	err = bm.Put(req.Mobile, size, time.Second*300)
	if err != nil {
		beego.Info("redis创建失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}

	return nil
}