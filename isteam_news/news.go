package isteam_news

import (
	steam "github.com/luanruisong/g-steam"
)

const (
	NewsServerName = "ISteamNews"
)

type (
	ISteamNews interface {
		GetNewsForApp(appid, enddate, count uint, feeds string) (uint, []NewsItemInfo, error)
	}
	iSteamNews struct {
		c steam.Client
	}
	NewsItemInfo struct {
		Gid           string `json:"gid" xml:"gid" form:"gid"`
		Title         string `json:"title" xml:"title" form:"title"`
		IsExternalUrl bool   `json:"is_external_url" xml:"is_external_url" form:"is_external_url"`
		Author        string `json:"author" xml:"author" form:"author"`
		Contents      string `json:"contents" xml:"contents" form:"contents"`
		Feedlabel     string `json:"feedlabel" xml:"feedlabel" form:"feedlabel"`
		Date          int64  `json:"date" xml:"date" form:"date"`
		Feedname      string `json:"feedname" xml:"feedname" form:"feedname"`
		FeedType      uint   `json:"feed_type" xml:"feed_type" form:"feed_type"`
		Appid         uint   `json:"appid" xml:"appid" form:"appid"`
	}
)

func (app *iSteamNews) apiServer() steam.Api {
	return app.c.Api().Server(NewsServerName)
}

func (app *iSteamNews) GetNewsForApp(appid, enddate, count uint, feeds string) (uint, []NewsItemInfo, error) {
	api := app.apiServer().Method("GetNewsForApp").Version("v0002").AddParam("appid", appid)
	if enddate > 0 {
		api = api.AddParam("enddate", enddate)
	}
	if count > 0 {
		api = api.AddParam("count", count)
	}
	if len(feeds) > 0 {
		api = api.AddParam("feeds", feeds)
	}

	var res struct {
		Appid     uint           `json:"appid" xml:"appid" form:"appid"`
		Count     uint           `json:"count" xml:"count" form:"count"`
		Newsitems []NewsItemInfo `json:"newsitems" xml:"newsitems" form:"newsitems"`
	}
	_, err := api.Get(&res)
	return res.Count, res.Newsitems, err
}

func New(c steam.Client) ISteamNews {
	return &iSteamNews{c: c}
}
