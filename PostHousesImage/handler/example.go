package handler

import (
	"context"

	example "MicroIhome/PostHousesImage/proto/example"
	"MicroIhome/IhomeWeb/utils"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"MicroIhome/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
)

type Example struct{}

func (e *Example) PostHousesImage(ctx context.Context, req *example.Request, rsp *example.Response) error {

	//打印被调用的函数
	beego.Info("发送房屋图片PostHousesImage  /api/v1.0/houses/:id/images")
	//初始化返回正确的返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	/*获取文件的后缀名*/ //dsnlkjfajadskfksda.sadsdasd.sdasd.jpg
	beego.Info("后缀名", path.Ext(req.Filename))
	//.jpg
	fileext := path.Ext(req.Filename)

	/*将获取到的图片数据成为二进制信息存入fastdfs*/
	GroupName, RemoteFileId, err := models.UploadByBuffer(req.Image, fileext[1:])
	if err != nil {
		beego.Info("Postupavatar  models.UploadByBuffer err", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	beego.Info(GroupName, RemoteFileId)

	/*从请求url中得到我们的house_id*/

	houseid, _ := strconv.Atoi(req.Id)

	//创建house 对象
	house := models.House{Id: houseid}
	//创建数据库句柄
	o := orm.NewOrm()
	err = o.Read(&house)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*判断index_image_url 是否为空 */
	if house.Index_image_url == "" {
		/*空就把这张图片设置为主图片*/
		house.Index_image_url = RemoteFileId
	}

	/*将该图片添加到 house的全部图片当中*/
	houseimage := models.HouseImage{House: &house, Url: RemoteFileId}

	house.Images = append(house.Images, &houseimage)
	//将图片对象插入表单之中
	_, err = o.Insert(&houseimage)
	if err != nil {

		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//对house表进行更新
	_, err = o.Update(&house)
	if err != nil {

		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*返回正确的数据回显给前端*/

	rsp.Url = RemoteFileId

	return nil
}
