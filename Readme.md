## 微信机器人-wxrobot
## Usage：

### 	download：

​		

```
cd ${GOPATH}/src
git clone https://github.com/godaner/wxrobot.git
```

### textreply.cfg：

```
[default]
hello: hello , i am godaner !

explain:
	*.it is a demo reply config
	*.don't change node [default].
    *.if anyone send "hello" to you , your wxrobot will response "hello , i am godaner !" to him auto.
```



### 	run demo：

​		

```
cd ${GOPATH}/src/wxrobot
go run main.go -textReplyPath ${GOPATH}/src/wxrobot/textreply.cfg
if you wanna run wxrobot in background:
	nohup go run main.go -textReplyPath ${GOPATH}/src/wxrobot/textreply.cfg >wxrobot.log 2>&1 & 
if you wanna see log:
	tail -f wxrobot.log
```



### 	demo result：

​	

```
2018/08/29 09:44:24 wx.go:104: Please open link in browser: https://login.weixin.qq.com/qrcode/IesWCyGxZg==
2018/08/29 09:44:49 wx.go:129: login timeout, reconnecting...
2018/08/29 09:45:09 wx.go:133: scan success, please confirm login on your phone
2018/08/29 09:45:12 wx.go:136: login success
2018/08/29 09:45:19 wx.go:305: update 141 contacts
2018/08/29 09:45:19 wx.go:313: @c458675e3c522f5f0bc436f0a861ca16 => 微信安全中心
2018/08/29 09:45:19 wx.go:313: @c8e81be227ce9428490833eb837a5ee81f3b9b6beed19809cdd4a53976fd4104 => 快乐人生
......
```

