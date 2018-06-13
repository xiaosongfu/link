# link
收集所有你想要收藏的网页 ^_^

---

> GoReporter (https://github.com/360EntSecGroup-Skylar/goreporter)
```
goreporter -p ./ -r ./goreport -f html
```

> 交叉编译命令
```
GOOS=linux GOARCH=amd64 go build

GOOS：386、amd64、arm
GOARCH：darwin、freebsd、linux、windows
```

> docker
```shell
docker build -t link:1.1.0 -f Dockerfile2 .

docker run -p 12054:1205 -d link:1.1.0
```