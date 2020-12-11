# g-steam 

steam web api golang realize

![](https://img.shields.io/badge/macOS-Development-d0d1d4)
![](https://img.shields.io/badge/golang-1.15-blue)
![](https://img.shields.io/badge/godoc-reference-3C57C4)
![](https://img.shields.io/badge/version-0.1-r)

## :rocket:Installation

`
  go get -u github.com/luanruisong/g-steam
`

## :bell:Premise

   Before using, you must go to [steam](https://steamcommunity.com/dev/revokekey) to apply for your key.
   
## :anchor:Usage

### Basic usage

```go
    //Create client
    client := steam.NewClient("appkey")
    //Get the render address of steam
    //path -> steam openid Login authentication address
    //callbackPath -> steam browser url to redirect to after successful authentication
    path := client.RenderTo(callbackPath)
    
    //Create receiving return object
    res, err := client.OpenidBindQuery(request.URL.Query())
    fmt.Println(res, err)
    fmt.Println(res.GetSteamId()) //Get steamid

    //Create api object
    api := client.Api()
    //raw return
    raw, err := api.Server("ISteamUser").//Set up service interface
        Method("GetPlayerSummaries").//Set access function
        Version("v0002").//Set version
        AddParam("steamids", "76561198421538055").//Setting parameters (If the key parameter is not set, the client's appKey will be added by default)
        Get(nil) //Initiate a request, and support the incoming structure pointer to receive parameters
    fmt.Println(raw, err) 

```

### Package API

In addition to providing basic encapsulation, we also encapsulate common APIs to make it more convenient to use.

```go
    //Use client to create related server
    appService := isteam_app.New(client)
    iplyerSercer := iplayer_service.New(client)
    economyService := isteam_economy.New(client)
    newsService := isteam_news.New(client)
    remoteStorageService := isteam_remote_storage.New(client)
    userService := isteam_user.New(client)
    userStatsService := isteam_user_stats.New(client)
    util := isteam_webapi_util.New(client)
    //Call the server wrapper function
```

## :tada:Contribute code

Open source projects are inseparable from everyone’s support. If you have a good idea, encountered some bugs and fixed them, and corrected the errors in the document, please submit a Pull Request~
   1. Fork this project to your own repo
   2. Clone the project in the past, that is, the project in your warehouse, to your local
   3. Modify the code
   4. Push to your own library after commit
   5. Initiate a PR (pull request) request and submit it to the `provide` branch
   6. Waiting to merge

## :closed_book:License

Distributed under MIT License, please see license file within the code for more details.



