## 微信机器人-wxrobot
## Usage：

### 	download：

​		

```
cd ${GOPATH}/src
git clone https://github.com/godaner/wxrobot.git
```



### 	run：

​		

```
cd ${GOPATH}/src/wxrobot
go run main.go -textReplyPath ${GOPATH}/src/wxrobot/textreply.cfg
if you wanna run wxrobot in background:
	nohup go run main.go -textReplyPath ${GOPATH}/src/wxrobot/textreply.cfg >wxrobot.log 2>&1 & 
if you wanna see log:
	tail -f wxrobot.log
```



### 	result:

​	

```
API server listening at: 127.0.0.1:63805
2018/08/29 16:09:59 wx.go:89: Please open link in browser: https://login.weixin.qq.com/qrcode/Admjurf_7g==
2018/08/29 16:10:06 wx.go:118: scan success, please confirm login on your phone
2018/08/29 16:10:08 wx.go:121: login success
2018/08/29 16:10:12 wx.go:290: update 141 contacts
2018/08/29 16:10:12 wx.go:413: @79d6d9e7408e66f2401204a8e31f26ece703127f81e98851c43e7245290fa770: 
2018/08/29 16:10:30 wx.go:413: @79d6d9e7408e66f2401204a8e31f26ece703127f81e98851c43e7245290fa770:你们几点结束？
2018/08/29 16:15:58 wx.go:413: @79d6d9e7408e66f2401204a8e31f26ece703127f81e98851c43e7245290fa770: 
```

