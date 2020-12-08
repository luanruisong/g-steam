# g-steam steam web api golang 实现

## 安装
`
  go get -u github.com/luanruisong/g-steam
`

## 秘钥申请
   [https://steamcommunity.com/dev/revokekey](https://steamcommunity.com/dev/revokekey)
## 使用

### 基础用法

```go
    //创建client
    client := steam.NewClient("appkey")
    //获取steam的render地址
    //path -> steam openid 登陆认证地址
    //callbackPath -> steam认证成功后 跳转的浏览器url
    path := client.RenderTo(callbackPath)
    
    //创建接收返回对象
    res, err := client.OpenidBindQuery(request.URL.Query())
    fmt.Println(res, err)
    fmt.Println(res.GetSteamId()) //获取steamid

    //创建api对象
    api := client.Api()
    //raw 原始返回
    raw, err := api.Server("ISteamUser").//设置服务接口
        Method("GetPlayerSummaries").//设置访问函数
        Version("v0002").//设置版本号
        AddParam("steamids", "76561198421538055").//设置参数 （key参数不设置默认会添加client的appKey）
        Get(nil) //发起请求，另外支持传入结构体指针用于接收参数
    fmt.Println(raw, err) //打印

```

### 封装实现

```go
    //统一创建client
    client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
    //使用client 创建相关的server
    appServer := isteam_app.New(client)
    //调用server包装函数
    appInfoList, err := apps.GetAppList()
    fmt.Println(appInfoList,err)
```




