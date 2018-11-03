## 预约借用系统
本系统主要是实现场地的预约或者借用，借用者可在前台进行预约申请，管理员可在后台审核预约，同时可以增删改场地和对前台的预约表单进行自定义的新增或者删除，灵活性较强，供有同样需求朋友参考使用。
![image](https://github.com/cnlh/booking/blob/master/index.png)
![image](https://github.com/cnlh/booking/blob/master/admin.png)
![image](https://github.com/cnlh/booking/blob/master/admin_1.png)
## 说明
- golang编写，无需任何扩展，编译好后直接可以在各个平台运行。
- 编写的目的是为了一个简单的功能，也是学习go语言过程中的一次入门练手，有很多不好的地方。
- 由于时间仓促，只实现了各个基本功能


## 配置说明

conf/app
```
appname = bookingSystem
httpport = 8080
runmode = dev

dbHost=127.0.0.1
dbPort=3306
dbUser=root
dbPassword=Aa11882200.
dbName=booking

sessionon = true
EnableAdmin = true

```
> EnableAdmin 是统计功能

## 安装使用
1. 安装beego

```
go get github.com/astaxie/beego

```

2. 本项目源码
```
go get github.com/cnlh/booking
```
3. 编译

```
go build
```

4. 初始化数据库

```
./bookingSystem orm syncdb
```

5. 运行
```
./bookingSystem &
```
## 特别感谢

- beego
- layui
- x-admin 后台模板
- bootstrap 前台




