MicroIhome:

ENVIRONMENT REQUIRED:

#ubuntu


#go
#protoBuf

protoc --go_out=./ *.proto #不加grpc插件

protoc --go_out=plugins=grpc:./ *.proto #添加grpc插件

#grpc

#consul

consul agent -dev

http://localhost:8500

consul agent -server -bootstrap-expect 2 -data-dir /tmp/consul -node=n1 - bind=192.168.150.20 -ui -config-dir /etc/consul.d -rejoin -join 192.168.150.20 - client 0.0.0.0

consul agent -server -bootstrap-expect 2 -data-dir /tmp/consul -node=n2 - bind=192.168.150.21 -ui -rejoin -join 192.168.150.20

consul agent -data-dir /tmp/consul -node=n3 -bind=192.168.150.23 -config-dir /etc/consul.d -rejoin -join 192.168.150.20

#micro

micro new --type "web" MicroIhome/IhomeWeb

micro new --type "srv" MicroIhome/GetArea

...

#httprouter

go get -u -v github.com/julienschmidt/httprouter

#beego

$ go get -u -v github.com/astaxie/beego

$ go get -u -v github.com/beego/bee

##配置beego执行文件环境变量：$GOPATH/bin

	$ vim .bashrc
  
	//在最后一行插入
  
	export PATH="$GOPATH/bin:$PATH"
  
	//然后保存退出
  
	$ source .bashrc
  
#mysql

#mysql-driver

go get -u -v github.com/go-sql-driver/mysql

#database

mysql -uroot -proot

drop database go1micro;

create database if not exists go1micro default charset utf8 collate utf8_general_ci;

use go1micro;

source go1micro.sql;

source ./conf/data.sql;

#redis+fdfs+nginx

/home/roger/go/src/MicroIhome/IhomeWeb/server.sh

#go-redis-api

go get -v -u github.com/gomodule/redigo/redis

go get -v -u github.com/garyburd/redigo

#beego-cache

go get -u -v github.com/astaxie/beego/cache

#图片验证码库

go get -u -v github.com/afocus/captcha
