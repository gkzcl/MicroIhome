package main

import (
	"github.com/micro/go-log"
	"net/http"

	"github.com/micro/go-web"
	"github.com/julienschmidt/httprouter"
	_ "MicroIhome/IhomeWeb/models"
	"MicroIhome/IhomeWeb/handler"
)

func main() {

	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":10086"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))

	// 获取地域信息
	rou.GET("/api/v1.0/areas", handler.GetArea)
	// 获取验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)
	// 获取短信验证码信息
	rou.GET("/api/v1.0/smscode/:mobile", handler.GetSmscd)
	// 注册用户
	rou.POST("/api/v1.0/users", handler.PostRet)
	// 获取session信息
	rou.GET("/api/v1.0/session", handler.GetSession)
	//登陆
	rou.POST("/api/v1.0/sessions", handler.PostLogin)
	//退出登陆
	rou.DELETE("/api/v1.0/session", handler.DeleteSession)
	// 获取用户信息
	rou.GET("/api/v1.0/user", handler.GetUserInfo)
	// 上传头像
	rou.POST("/api/v1.0/user/avatar", handler.PostAvatar)
	//请求更新用户名 PUT
	rou.PUT("/api/v1.0/user/name",handler.PutUserInfo)
	// 用户认证检查
	rou.GET("/api/v1.0/user/auth", handler.GetUserAuth)
	// 用户实名认证
	rou.POST("/api/v1.0/user/auth", handler.PostUserAuth)
	//请求当前用户已发布房源信息  GET
	rou.GET("/api/v1.0/user/houses", handler.GetUserHouses)
	//发布房源信息 POST
	rou.POST("/api/v1.0/houses", handler.PostHouses)
	//上传房源图片信息  POST
	rou.POST("/api/v1.0/houses/:id/images", handler.PostHousesImage)
	// 请求房源详细信息 GET
	rou.GET("/api/v1.0/houses/:id", handler.GetHouseInfo)
	// 获取首页轮播图
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	// 搜索  api/v1.0/houses?aid=5&sd=2017-11-12&ed=2017-11-30&sk=new&p=1
	rou.GET("/api/v1.0/houses", handler.GetHouses)
	//post 发布订单 api/v1.0/orders
	rou.POST("/api/v1.0/orders", handler.PostOrders)
	//get 查看房东/租客订单信息请求
	rou.GET("/api/v1.0/user/orders", handler.GetUserOrder)
	//put房东同意/拒绝订单
	//api/v1.0/orders/:id/status
	rou.PUT("/api/v1.0/orders/:id/status", handler.PutOrders)
	//PUT 用户评价订单信请求
	//api/v1.0/orders/:id/comment
	//api/v1.0/orders/1/comment
	rou.PUT("/api/v1.0/orders/:id/comment", handler.PutComment)


	service.Handle("/", rou)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
