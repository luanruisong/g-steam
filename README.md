# g-steam

steam openid 2.0 golang 实现

```go


    //获取steam的render地址
    //path -> steam openid 登陆认证地址
    //callbackPath -> steam认证成功后 跳转的浏览器url
    path := steam.RenderTo(callbackPath)
    
    //创建接收返回对象
    openidRes := new(steam.OpenidRes)
    //绑定返回参数
    openidRes.Bind(*http.Request)//原生 包含validateSteamSing
    gin.Context.ShouldBind(openidRes) //gin框架
    //。。。其他绑定方式

    //验证是steam的请求 用于其他参数绑定方式
    //多次验证会返回失败
    openidRes.ValidateSteamSign() == nil 

    openidRes.GetSteamId() //获取steamid

```




