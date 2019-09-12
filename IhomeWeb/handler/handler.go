package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"github.com/micro/go-micro/client"
	example "github.com/micro/examples/template/srv/proto/example"
	GETAREA "MicroIhome/GetArea/proto/example"
	GETINDEX "MicroIhome/GetIndex/proto/example"
	GETIMAGECD "MicroIhome/GetImageCd/proto/example"
	GETSESSION "MicroIhome/GetSession/proto/example"
	GETSMSCD "MicroIhome/GetSmscd/proto/example"
	POSTRET "MicroIhome/PostRet/proto/example"
	POSTLOGIN "MicroIhome/PostLogin/proto/example"
	DELETESESSION "MicroIhome/DeleteSession/proto/example"
	GETUSERINFO "MicroIhome/GetUserInfo/proto/example"
	POSTAVATAR "MicroIhome/PostAvatar/proto/example"
	POSTUSERAUTH "MicroIhome/PostUserAuth/proto/example"
	GETUSERHOUSES "MicroIhome/GetUserHouses/proto/example"
	POSTHOUSES "MicroIhome/PostHouses/proto/example"
	POSTHOUSESIMAGE "MicroIhome/PostHousesImage/proto/example"
	GETHOUSEINFO "MicroIhome/GetHouseInfo/proto/example"
	GETHOUSES "MicroIhome/GetHouses/proto/example"
	POSTORDERS "MicroIhome/PostOrders/proto/example"
	GETUSERORDER "MicroIhome/GetUserOrder/proto/example"
	PUTORDERS "MicroIhome/PutOrders/proto/example"
	PUTCOMMENT "MicroIhome/PutComment/proto/example"
	PUTUSERINFO "MicroIhome/PutUserInfo/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"MicroIhome/IhomeWeb/models"
	"MicroIhome/IhomeWeb/utils"
	"image"
	"image/png"
	"github.com/astaxie/beego"
	"regexp"
	"github.com/afocus/captcha"
	"fmt"
	"io/ioutil"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 获取区域信息
func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 创建服务，获取句柄
	server := grpc.NewService()
	// 初始化服务
	server.Init()
	// 调用服务，返回句柄
	getAreaService := GETAREA.NewExampleService("go.micro.srv.GetArea", server.Client())
	// 调用服务，返回数据
	rsp, err := getAreaService.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	area_list := []models.Area{}

	for _, value := range rsp.Data {
		tmp := models.Area{Id: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tmp)
	}

	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.ErrorMsg,
		"data":   area_list,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 获取首页轮播图信息
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("获取首页轮播 url：api/v1.0/houses/index")
	server :=grpc.NewService()
	server.Init()

	exampleClient := GETINDEX.NewExampleService("go.micro.srv.GetIndex", server.Client())


	rsp, err := exampleClient.GetIndex(context.TODO(),&GETINDEX.Request{})
	if err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 502)
		return
	}
	data := []interface{}{}
	json.Unmarshal(rsp.Max,&data)

	//创建返回数据map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":data,
	}
	w.Header().Set("Content-Type", "application/json")

	// 将返回数据map发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

}

// 获取验证码图片信息
func GetImageCd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	server := grpc.NewService()
	server.Init()

	// 调用服务
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", server.Client())

	//获取uuid
	uuid := ps.ByName("uuid")

	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid: uuid,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//接收图片信息的 图片格式
	var img image.RGBA

	img.Stride = int(rsp.Stride)
	img.Pix = []uint8(rsp.Pix)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	var image captcha.Image

	image.RGBA = &img

	//将图片发送给浏览器
	png.Encode(w, image)

}

// 获取短信验证码信息
func GetSmscd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	beego.Info("获取短信验证码 GetSmscd /api/v1.0/smscode/:mobile ")
	//通过传入参数URL下Query获取前端的在url里的带参
	//beego.Info(r.URL.Query())
	//map[text:[9346] id:[474494b0-18eb-4eb7-9e68-a5ecf3c8317b]]
	//获取参数
	test := r.URL.Query()["text"][0]
	id := r.URL.Query()["id"][0]
	mobile := ps.ByName("mobile")

	//通过正则进行手机号的判断
	//创建正则条件
	mobile_reg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	//通过条件判断字符串是否匹配规则 返回正确或失败
	bl := mobile_reg.MatchString(mobile)
	//如果手机号不匹配那就直接返回错误不调用服务
	if bl == false {
		response := map[string]interface{}{
			"error":  utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// 发送数据
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	//创建并初始化服务
	server := grpc.NewService()
	server.Init()

	// 调用服务
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", server.Client())
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile:   mobile,
		Imagestr: test,
		Uuid:     id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{

		"errno":  rsp.Error,
		"errmsg": rsp.ErrMsg,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")

	// 发送数据
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

// 注册用户信息
func PostRet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("PostRet  注册 /api/v1.0/users")

	//服务创建
	server := grpc.NewService()
	server.Init()

	//接收post发送过来的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if request["mobile"].(string) == "" || request["password"].(string) == "" || request["sms_code"].(string) == "" {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		return

	}

	// 调用请求
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", server.Client())
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
		SmsCode:  request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//读取cookie   统一cookie   userlogin
	//func (r *Request) Cookie(name string) (*Cookie, error)

	cookie, err := r.Cookie("userlogin")
	if err != nil || "" == cookie.Value {
		//创建1个cookie对象
		cookie := http.Cookie{Name: "userlogin", Value: rsp.SessionId, Path: "/", MaxAge: 3600}
		//对浏览器的cookie进行设置
		http.SetCookie(w, &cookie)
	}

	//准备回传数据
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	//发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 获取session信息
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取session信息 GetSession /api/v1.0/session")

	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		// 直接返回说名用户未登陆
		response := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		// 将数据回发给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETSESSION.NewExampleService("go.micro.srv.GetSession", server.Client())
	rsp, err := exampleClient.GetSession(context.TODO(), &GETSESSION.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := make(map[string]string)
	data["name"] = rsp.UserName

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 登陆
func PostLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("登陆  PostLogin /api/v1.0/sessions")

	// 接收前端发送过来的json数据进行解码
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if request["mobile"].(string) == "" || request["password"].(string) == "" {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		return
	}

	//创建服务
	server := grpc.NewService()
	server.Init()

	// 调用服务
	exampleClient := POSTLOGIN.NewExampleService("go.micro.srv.PostLogin", server.Client())
	rsp, err := exampleClient.PostLogin(context.TODO(), &POSTLOGIN.Request{
		Mobile:   request["mobile"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//设置cookie
	//Cookie读取
	cookie, err := r.Cookie("userlogin")

	if err != nil || cookie.Value == "" {
		cookie := http.Cookie{Name: "userlogin", Value: rsp.Sessionid, Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 退出登陆
func DeleteSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("DeleteSession  退出登陆 /api/v1.0/session")

	//创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := DELETESESSION.NewExampleService("go.micro.srv.DeleteSession", server.Client())

	//获取cookie
	cookie, err := r.Cookie("userlogin")

	if err != nil || cookie.Value == "" {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	rsp, err := exampleClient.DeleteSession(context.TODO(), &DELETESESSION.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//删除sessionid
	cookie, err = r.Cookie("userlogin")
	if cookie.Value != "" || err == nil {
		cookie := http.Cookie{Name: "userlogin", Path: "/", MaxAge: -1, Value: ""}
		http.SetCookie(w, &cookie)
	}
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("获取用户信息 GetUserInfo /api/v1.0/user ")

	client := grpc.NewService()
	client.Init()

	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", client.Client())
	//获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//远程调用函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := make(map[string]interface{})
	data["name"] = rsp.Name
	data["user_id"] = rsp.UserId
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 上传头像
func PostAvatar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("上传头像  PostAvatar /api/v1.0/user/avatar")

	//获取到前端发送的文件信息
	//func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	File, FileHeader, err := r.FormFile("avatar")
	if err != nil {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	beego.Info("文件大小", FileHeader.Size)
	beego.Info("文件名", FileHeader.Filename)

	//创建一个文件大小的切片
	filebuf := make([]byte, FileHeader.Size)

	//将file的数据读到filebuf
	_, err = File.Read(filebuf)
	if err != nil {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//获取cookie
	cookie, err := r.Cookie("userlogin")
	beego.Info(cookie)
	if err != nil || cookie.Value == "" {
		beego.Info("hello")
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}
	//连接服务
	client := grpc.NewService()
	client.Init()

	// call the backend service
	exampleClient := POSTAVATAR.NewExampleService("go.micro.srv.PostAvatar", client.Client())
	rsp, err := exampleClient.PostAvatar(context.TODO(), &POSTAVATAR.Request{
		SessionId: cookie.Value,
		Fileext:   FileHeader.Filename,
		Filesize:  FileHeader.Size,
		Avatar:    filebuf,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]string)
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)
	beego.Info(utils.AddDomain2Url(rsp.AvatarUrl))

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 用户信息检查
func GetUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("用户信息检查 GetUserAuth /api/v1.0/user ")

	client := grpc.NewService()
	client.Init()

	// call the backend service
	exampleClient := GETUSERINFO.NewExampleService("go.micro.srv.GetUserInfo", client.Client())
	//获取cookie
	cookie, err := r.Cookie("userlogin")
	if err != nil || cookie.Value == "" {
		//准备回传数据
		response := map[string]interface{}{
			"errno":  utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type", "application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//远程调用函数
	rsp, err := exampleClient.GetUserInfo(context.TODO(), &GETUSERINFO.Request{
		Sessionid: cookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string]interface{})
	data["name"] = rsp.Name
	data["user_id"] = rsp.UserId
	data["mobile"] = rsp.Mobile
	data["real_name"] = rsp.RealName
	data["id_card"] = rsp.IdCard
	data["avatar_url"] = utils.AddDomain2Url(rsp.AvatarUrl)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.ErrMsg,
		"data":   data,
	}
	//设置返回数据的格式
	w.Header().Set("Content-Type", "application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// 用户实名认证
func PostUserAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info(" 实名认证 Postuserauth  api/v1.0/user/auth ")

	service := grpc.NewService()
	service.Init()

	//获取前端发送的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := POSTUSERAUTH.NewExampleService("go.micro.srv.PostUserAuth", service.Client())

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	beego.Info(userlogin.Value)
	beego.Info(request["real_name"].(string))
	beego.Info(request["id_card"].(string))
	rsp, err := exampleClient.PostUserAuth(context.TODO(), &POSTUSERAUTH.Request{
		Sessionid: userlogin.Value,
		RealName:  request["real_name"].(string),
		IdCard:    request["id_card"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

// 获取当前用户所发布的房源 GetUserHouses
func GetUserHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("获取当前用户所发布的房源 GetUserHouses /api/v1.0/user/houses")

	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETUSERHOUSES.NewExampleService("go.micro.srv.GetUserHouses", server.Client())

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.GetUserHouses(context.TODO(), &GETUSERHOUSES.Request{
		Sessionid: userlogin.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	house_list := []models.House{}
	json.Unmarshal(rsp.Mix, &house_list)

	var houses []interface{}
	for _, houseinfo := range house_list {
		fmt.Printf("house.user = %+v\n", houseinfo.Id)
		fmt.Printf("house.area = %+v\n", houseinfo.Area)
		houses = append(houses, houseinfo.To_house_info())
	}

	data_map := make(map[string]interface{})
	data_map["houses"] = houses

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

//发布房源信息
func PostHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("PostHouses 发布房源信息 /api/v1.0/houses ")
	//获取前端post请求发送的内容
	body, _ := ioutil.ReadAll(r.Body)

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}
		//设置回传格式
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	//创建连接

	service := grpc.NewService()
	service.Init()
	exampleClient := POSTHOUSES.NewExampleService("go.micro.srv.PostHouses", service.Client())

	rsp, err := exampleClient.PostHouses(context.TODO(), &POSTHOUSES.Request{
		Sessionid: userlogin.Value,
		Max:       body,
	})
	if err != nil {
		http.Error(w, err.Error(), 502)

		beego.Info(err)
		//beego.Debug(err)
		return
	}

	/*得到插入房源信息表的 id*/
	houseid_map := make(map[string]interface{})
	houseid_map["house_id"] = int(rsp.House_Id)

	resp := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   houseid_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 503)
		beego.Info(err)
		return
	}
}

//发送房屋图片
func PostHousesImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("发送房屋图片PostHousesImage  /api/v1.0/houses/:id/images")

	//创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := POSTHOUSESIMAGE.NewExampleService("go.micro.srv.PostHousesImage", server.Client())
	//获取houserid
	houseid := ps.ByName("id")
	//获取sessionid
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	file, hander, err := r.FormFile("house_image")
	if err != nil {
		beego.Info("Postupavatar   c.GetFile(avatar) err", err)

		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	beego.Info(file, hander)
	beego.Info("文件大小", hander.Size)
	beego.Info("文件名", hander.Filename)
	//二进制的空间用来存储文件
	filebuffer := make([]byte, hander.Size)
	//将文件读取到filebuffer里
	_, err = file.Read(filebuffer)
	if err != nil {
		beego.Info("Postupavatar   file.Read(filebuffer) err", err)
		resp := map[string]interface{}{
			"errno":  utils.RECODE_IOERR,
			"errmsg": utils.RecodeText(utils.RECODE_IOERR),
		}
		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.PostHousesImage(context.TODO(), &POSTHOUSESIMAGE.Request{
		Sessionid: userlogin.Value,
		Id:        houseid,
		Image:     filebuffer,
		Filesize:  hander.Size,
		Filename:  hander.Filename,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//准备返回值
	data := make(map[string]interface{})
	data["url"] = utils.AddDomain2Url(rsp.Url)
	// 返回数据map
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")

	// 回发数据
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

//获取房源详细信息
func GetHouseInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("获取房源详细信息 GetHouseInfo  api/v1.0/houses/:id ")

	//创建服务
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETHOUSEINFO.NewExampleService("go.micro.srv.GetHouseInfo", server.Client())

	id := ps.ByName("id")

	//获取sessionid
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.GetHouseInfo(context.TODO(), &GETHOUSEINFO.Request{
		Sessionid: userlogin.Value,
		Id:        id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	house := models.House{}
	json.Unmarshal(rsp.Housedata, &house)

	data_map := make(map[string]interface{})
	data_map["user_id"] = int(rsp.Userid)
	data_map["house"] = house.To_one_house_desc()

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	return
}

//搜索房屋
func GetHouses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("GetHouses")

	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := GETHOUSES.NewExampleService("go.micro.srv.GetHouses", server.Client())

	//aid=5&sd=2017-11-12&ed=2017-11-30&sk=new&p=1
	aid := r.URL.Query()["aid"][0] //aid=5   地区编号
	sd := r.URL.Query()["sd"][0]   //sd=2017-11-1   开始世界
	ed := r.URL.Query()["ed"][0]   //ed=2017-11-3   结束世界
	sk := r.URL.Query()["sk"][0]   //sk=new    第三栏条件
	p := r.URL.Query()["p"][0]     //tp=1   页数

	rsp, err := exampleClient.GetHouses(context.TODO(), &GETHOUSES.Request{
		Aid: aid,
		Sd:  sd,
		Ed:  ed,
		Sk:  sk,
		P:   p,
	})
	beego.Info(rsp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	houses_l := []interface{}{}
	json.Unmarshal(rsp.Houses, &houses_l)

	data := map[string]interface{}{}
	data["current_page"] = rsp.CurrentPage
	data["houses"] = houses_l
	data["total_page"] = rsp.TotalPage

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

//发布订单
func PostOrders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("PostOrders  发布订单 /api/v1.0/orders")

	//将post代过来的数据转化以下
	body, _ := ioutil.ReadAll(r.Body)

	userlogin, err := r.Cookie("userlogin")
	if err != nil || userlogin.Value == "" {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	service := grpc.NewService()
	service.Init()

	//调用服务
	exampleClient := POSTORDERS.NewExampleService("go.micro.srv.PostOrders", service.Client())
	rsp, err := exampleClient.PostOrders(context.TODO(), &POSTORDERS.Request{
		Sessionid: userlogin.Value,
		Body:      body,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*得到插入房源信息表的 id*/
	houseid_map := make(map[string]interface{})
	houseid_map["order_id"] = int(rsp.OrderId)

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   houseid_map,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

//获取订单
func GetUserOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info("/api/v1.0/user/orders   GetUserOrder 获取订单 ")
	server := grpc.NewService()
	server.Init()
	// call the backend service
	exampleClient := GETUSERORDER.NewExampleService("go.micro.srv.GetUserOrder", server.Client())

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}
	//获取role
	role := r.URL.Query()["role"][0] //role

	rsp, err := exampleClient.GetUserOrder(context.TODO(), &GETUSERORDER.Request{
		Sessionid: userlogin.Value,
		Role:      role,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	order_list := []interface{}{}
	json.Unmarshal(rsp.Orders, &order_list)

	data := map[string]interface{}{}
	data["orders"] = order_list

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}

	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

func PutOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// decode the incoming request as json
	//接收请求携带的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 502)
			beego.Info(err)
			return
		}
		return
	}
	server := grpc.NewService()
	server.Init()

	// call the backend service
	exampleClient := PUTORDERS.NewExampleService("go.micro.srv.PutOrders", server.Client())

	rsp, err := exampleClient.PutOrders(context.TODO(), &PUTORDERS.Request{
		Sessionid: userlogin.Value,
		Action:    request["action"].(string),
		Orderid:   ps.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 504)
		return
	}
}

//用户评价订单
func PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("PutComment  用户评价 /api/v1.0/orders/:id/comment")
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	service := grpc.NewService()
	service.Init()
	// call the backend service
	exampleClient := PUTCOMMENT.NewExampleService("go.micro.srv.PutComment", service.Client())

	//获取cookie
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.PutComment(context.TODO(), &PUTCOMMENT.Request{
		Sessionid: userlogin.Value,
		Comment:   request["comment"].(string),
		OrderId:   ps.ByName("id"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	w.Header().Set("Content-Type", "application/json")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}

//更新用户名
func PutUserInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	beego.Info(" 更新用户名 Putuserinfo /api/v1.0/user/name")
	//创建服务
	service := grpc.NewService()
	service.Init()
	// 接收前端发送内容
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 调用服务
	exampleClient := PUTUSERINFO.NewExampleService("go.micro.srv.PutUserInfo", service.Client())

	//获取用户登陆信息
	userlogin, err := r.Cookie("userlogin")
	if err != nil {
		resp := map[string]interface{}{
			"errno":  utils.RECODE_SESSIONERR,
			"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
		}

		w.Header().Set("Content-Type", "application/json")
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), 503)
			beego.Info(err)
			return
		}
		return
	}

	rsp, err := exampleClient.PutUserInfo(context.TODO(), &PUTUSERINFO.Request{
		Sessionid: userlogin.Value,
		Username:  request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收回发数据
	data := make(map[string]interface{})
	data["name"] = rsp.Username

	response := map[string]interface{}{
		"errno":  rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":   data,
	}
	w.Header().Set("Content-Type", "application/json")

	// 返回前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
}
