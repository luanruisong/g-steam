package isteam_webapi_util

import (
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() *iSteamWebapiUtil {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetServerInfo(t *testing.T) {
	app := getTestApps()
	res, err := app.GetServerInfo()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(res)
	}
}

func TestGetSupportedAPIList(t *testing.T) {
	app := getTestApps()
	res, err := app.GetSupportedAPIList()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(res)
	}
}
