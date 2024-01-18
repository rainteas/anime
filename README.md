## 蜜柑动漫订阅下载器

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&pause=1000&random=false&width=435&lines=%E8%9C%9C%E6%9F%91%E5%8A%A8%E6%BC%AB%E8%87%AA%E5%8A%A8%E6%9B%B4%E6%96%B0%E4%B8%8B%E8%BD%BD)](https://git.io/typing-svg)

### 1. 介绍
自动从订阅中获取最新的动漫更新，然后下载到本地。仅用于学习

### 2.依赖
依赖transmission下载器，aria2还未完成需要可以提交pr

### 3.编译

```
CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o anime main.go
```

### 4.配置
参考项目内config.ini
提前导入animation.sql 数据库文件。
rss_meta放入每个动漫的订阅地址、名称、季节
item表为记录哪些下载过了，不需要重复下载

### 4.启动

```
anime "./config.ini"
```

### 贡献
欢迎提交issue和PR

