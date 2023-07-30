# bubble清单

本项目来源知了传课：https://www.bilibili.com/video/BV1gJ411p7xC/?spm_id_from=333.999.0.0&vd_source=84fc27804252448ba51ef3b6abfd5d36

笔者对代码有优化更新。

一个基于gin+gorm开发的练手小项目，通过该项目可初识go web开发该有的姿势。



## 使用指南

### 下载

```
git clone https://github.com/uestc-wxy/bubble.git
```

### 配置MySQL

1. 在你的数据库中执行以下命令，创建本项目所用的数据库：

```
CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
```

1. 在`bubble/conf/config.ini`文件中按如下提示配置数据库连接信息。

```
port = 8080
release = false

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
db = bubble
```

### 编译

```
go build
```

### 执行

```bash
./bubble
```

启动之后，使用浏览器打开`http://127.0.0.1:8080/`即可。

![image](https://github.com/uestc-wxy/bubble/assets/120303802/153e14ac-08dd-49e3-9d9d-210701c7c9a2)
