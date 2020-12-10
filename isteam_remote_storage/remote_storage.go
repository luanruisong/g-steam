package isteam_news

import (
	"fmt"

	steam "github.com/luanruisong/g-steam"
)

const (
	RemoteStorageServerName = "ISteamRemoteStorage"
)

type (
	iSteamRemoteStorage struct {
		c steam.Client
	}

	ugcFile struct {
		Filename string `json:"filename" xml:"filename" form:"filename"`
		Url      string `json:"url" xml:"url" form:"url"`
		Size     uint   `json:"size" xml:"size" form:"size"`
	}
)

func (app *iSteamRemoteStorage) apiServer() steam.Api {
	return app.c.Api().Server(RemoteStorageServerName)
}

func (app *iSteamRemoteStorage) GetCollectionDetails(collectioncount uint32, publishedfileids []uint64) (string, error) {
	api := app.apiServer().Method("GetCollectionDetails").Version("v0001").AddParam("collectioncount", collectioncount)
	for i, v := range publishedfileids {
		api = api.AddParam(fmt.Sprintf("publishedfileids[%d]", i), v)
	}
	return api.Post(nil)
}

func (app *iSteamRemoteStorage) GetPublishedFileDetails(itemcount uint32, publishedfileids []uint64) (string, error) {
	api := app.apiServer().Method("GetPublishedFileDetails").Version("v1").AddParam("itemcount", itemcount)
	for i, v := range publishedfileids {
		api = api.AddParam(fmt.Sprintf("publishedfileids[%d]", i), v)
	}
	return api.Post(nil)
}

func (app *iSteamRemoteStorage) GetUGCFileDetails(appid uint32, ugcid, steamid string) (uint, []ugcFile, error) {
	api := app.apiServer().Method("GetUGCFileDetails").Version("v1").AddParam("appid", appid).AddParam("ugcid", ugcid)
	if len(steamid) > 0 {
		api = api.AddParam("steamid", steamid)
	}
	var res struct {
		Status struct {
			Code uint `json:"code" xml:"code" form:"code"`
		} `json:"status" xml:"status" form:"status"`
		Data []ugcFile `json:"data" xml:"data" form:"data"`
	}
	_, err := api.Get(&res)
	return res.Status.Code, res.Data, err
}

func New(c steam.Client) *iSteamRemoteStorage {
	return &iSteamRemoteStorage{c: c}
}
