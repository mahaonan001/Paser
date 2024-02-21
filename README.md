

# PaSer问卷管理系统




 
## 目录

- [PaSer问卷管理系统](#paser问卷管理系统)
  - [目录](#目录)
          - [开发前的配置要求](#开发前的配置要求)
          - [**安装步骤**](#安装步骤)
    - [文件目录说明](#文件目录说明)
    - [文件说明](#文件说明)
    - [路由](#路由)
    - [部署](#部署)
    - [使用到的框架](#使用到的框架)
    - [贡献者](#贡献者)
      - [如何参与开源项目](#如何参与开源项目)
    - [版本控制](#版本控制)
    - [作者](#作者)
    - [鸣谢](#鸣谢)




###### 开发前的配置要求
1. 需要配备go1.12及以上版本
2. 一款可以编写go语言的编译器或者任意命令行工具
3. 需要安装mysql
###### **安装步骤**

```sh
git clone https://github.com/mahaonan001/Paser.git
```
1. 找到PaSer目录下conifg中的application.yml并将其中的username，password改为该机器上的mysqlusername和password
2. cd PaSer
3. go mod tidy
4. Mac or Linux
- go build -o PaSer main.go
- ./PaSer
5. Windows
- go build -o PaSer.exe   
- .\PaSer.exe


### 文件目录说明


```
PaSer 
├── /common/
│  ├── db_admin.go
│  ├── db_asp.go
│  ├── db_uploader.go
│  ├── init_db.go
│  └── jwt.go
├── /config/
│  └── application.yml
├── /controller/
├── db_asp.go
│  ├── aboutPaper.go
│  ├── AskPaper.go
│  ├── controller.go
│  └── controller_User.go
├── /dto/
│  └── dto.go
├── /mail/
│  └── mail.go
├── /middleware/
│  └── maAuthMiddeWareil.go
├── /model/
│  ├── paperType.go
│  └── user.go
├── /response/
│  └── response.go
├── /router/
│  └── router.go
├── /util/
│  └── utils.go
├── go.mod
├── main.go
├── paper.json
├── paper_up.json
├── Paser.md
└── README.md

```

### 文件说明
1.    paper.json为管理增添问卷模板
2.    paper_up.json为用户上传问卷模板

### 路由
[此处是apiforx的路由链接](https://github.com/mahaonan001/Paser/blob/master/Paser.md)
### 部署

暂无

### 使用到的框架

- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm.io)
- [Viper](https://github.com/spf13/viper)

### 贡献者

暂无
#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

mahaonan001
1649801526@qq.com

 qq:1649801526    

 *您也可以在贡献者名单中参看所有参与该项目的开发者。*


### 鸣谢


- [shaojintian](https://github.com/shaojintian/Best_README_template)大佬的README.md模板




