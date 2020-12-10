package isteam_webapi_util

import (
	steam "github.com/luanruisong/g-steam"
)

const (
	WebApiUtilServerName = "ISteamWebAPIUtil"
)

type (
	iSteamWebapiUtil struct {
		c steam.Client
	}

	serverTime struct {
		Servertime       int64  `json:"servertime" xml:"servertime" form:"servertime"`
		Servertimestring string `json:"servertimestring" xml:"servertimestring" form:"servertimestring"`
	}
	ApiList struct {
		Interfaces []struct {
			Name    string `json:"name"`
			Methods []struct {
				Name       string        `json:"name"`
				Version    int           `json:"version"`
				Httpmethod string        `json:"httpmethod"`
				Parameters []interface{} `json:"parameters"`
			} `json:"methods"`
		} `json:"interfaces"`
	}
)

func (app *iSteamWebapiUtil) apiServer() steam.Api {
	return app.c.Api().Server(WebApiUtilServerName)
}

func (app *iSteamWebapiUtil) GetServerInfo() (serverTime, error) {
	api := app.apiServer().
		Method("GetServerInfo").
		Version("v0001")
	var res serverTime
	_, err := api.Get(&res)
	return res, err
}

func (app *iSteamWebapiUtil) GetSupportedAPIList() (ApiList, error) {
	api := app.apiServer().
		Method("GetSupportedAPIList").
		Version("v0001")
	var res struct {
		ApiList ApiList `json:"apilist" xml:"apilist" form:"apilist"`
	}
	_, err := api.Get(&res)
	return res.ApiList, err
}

func New(c steam.Client) *iSteamWebapiUtil {
	return &iSteamWebapiUtil{c: c}
}
