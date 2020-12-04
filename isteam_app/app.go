package isteam_app

import (
	"github.com/luanruisong/g-steam"
)

const (
	AppServerName = "iSteamApps"
)

type (
	iSteamApps struct {
		api steam.Api
	}
	appInfo struct {
		Appid uint   `json:"appid" xml:"appid" form:"appid"`
		Name  string `json:"name" xml:"name" form:"name"`
	}
	serverInfo struct {
		Addr     string `json:"addr" xml:"addr" form:"addr"`
		Gmsindex uint   `json:"gmsindex" xml:"gmsindex" form:"gmsindex"`
		Appid    uint   `json:"appid" xml:"appid" form:"appid"`
		Gamedir  string `json:"gamedir" xml:"gamedir" form:"gamedir"`
		Region   uint   `json:"region" xml:"region" form:"region"`
		Secure   bool   `json:"secure" xml:"secure" form:"secure"`
		Lan      bool   `json:"lan" xml:"lan" form:"lan"`
		Gameport uint   `json:"gameport" xml:"gameport" form:"gameport"`
		Specpot  uint   `json:"specpot" xml:"specpot" form:"specpot"`
	}

	upToDateInfo struct {
		UpToDate          bool   `json:"up_to_date" xml:"up_to_date" form:"up_to_date"`
		VersionIsListable bool   `json:"version_is_listable" xml:"version_is_listable" form:"version_is_listable"`
		RequiredVersion   uint   `json:"required_version" xml:"required_version" form:"required_version"`
		Message           string `json:"message" xml:"message" form:"message"`
	}
)

func (app *iSteamApps) apiServer() steam.Api {
	return app.api.Server(AppServerName)
}
func (app *iSteamApps) GetAppList() ([]appInfo, error) {
	var res struct {
		Applist struct {
			Apps []appInfo `json:"apps" form:"apps"`
		} `json:"applist" form:"applist"`
	}
	_, err := app.apiServer().Method("GetAppList").Version("v2").Get(&res)
	return res.Applist.Apps, err
}

func (app *iSteamApps) GetServersAtAddress(ip string) (bool, []serverInfo, error) {
	var res struct {
		Response struct {
			Success bool         `json:"success" form:"success"`
			Servers []serverInfo `json:"servers" form:"servers"`
		} `json:"response" form:"response"`
	}
	_, err := app.apiServer().
		Method("GetServersAtAddress").
		Version("v1").
		AddParam("addr", ip).
		Get(&res)
	return res.Response.Success, res.Response.Servers, err
}

func (app *iSteamApps) UpToDateCheck(appid uint, version string) (bool, upToDateInfo, error) {
	var res struct {
		Response struct {
			upToDateInfo
			Success bool `json:"success" form:"success"`
		} `json:"response" form:"response"`
	}
	_, err := app.apiServer().
		Method("UpToDateCheck").
		Version("v1").
		AddParam("appid", appid).
		AddParam("version", version).
		Get(&res)
	return res.Response.Success, res.Response.upToDateInfo, err
}

func New(c steam.Client) *iSteamApps {
	return &iSteamApps{api: c.Api()}
}
