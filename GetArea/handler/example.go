package handler

import (
	"context"
	example "MicroIhome/GetArea/proto/example"
	"MicroIhome/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"MicroIhome/IhomeWeb/models"
	"encoding/json"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"

)

type Example struct{}

func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {

	rsp.Error = utils.RECODE_OK
	rsp.ErrorMsg = utils.RecodeText(rsp.Error)

	// 初始化redis
	redis_conf := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}

	redis_conf_js, _ := json.Marshal(redis_conf)
	// 创建redis缓存
	bm, err := cache.NewCache("redis", string(redis_conf_js))
	if err != nil {
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrorMsg = utils.RecodeText(rsp.Error)
	}
	// redis缓存获取区域信息
	area_value := bm.Get("area_info00")

	// 区域信息存在,从redis缓存中获取区域信息.
	if area_value != nil {

		area_map := []map[string]interface{}{}
		//将获取到的数据进行json的解码操作
		json.Unmarshal(area_value.([]byte), &area_map)
		//将查询到的数据按照proto的格式发送给web服务.
		for _, value := range area_map {
			tmp := example.Response_Areas{Aid: int32(value["aid"].(float64)), Aname: value["aname"].(string)}
			rsp.Data = append(rsp.Data, &tmp)
		}
		return nil
	}
	// 区域信息不存在,从mysql数据库中查询区域信息.
	o := orm.NewOrm()
	qs := o.QueryTable("area")
	var area []models.Area
	num, err := qs.All(&area)
	if err != nil {
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrorMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	if num == 0 {
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrorMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	// 将mysql中查询获取区域的数据转换成json格式
	area_json, _ := json.Marshal(area)

	// 区域信息存入redis缓存
	err = bm.Put("area_info", area_json, time.Second*3600)
	if err != nil {
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrorMsg = utils.RecodeText(rsp.Error)
	}
	//将查询到的数据按照proto的格式发送给web服务.
	for _, value := range area {
		tmp := example.Response_Areas{Aid: int32(value.Id), Aname: value.Name}
		rsp.Data = append(rsp.Data, &tmp)
	}
	return nil
}