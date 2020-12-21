package isteam_app

import (
	"testing"

	"github.com/luanruisong/g-steam"
)

func getTestApps() ISteamApps {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetAppList(t *testing.T) {
	apps := getTestApps()
	infoList, err := apps.GetAppList()
	if err == nil {
		t.Log(len(infoList))
	} else {
		t.Error(err.Error())
	}

}

func TestGetServersAtAddress(t *testing.T) {
	apps := getTestApps()
	succ, list, err := apps.GetServersAtAddress("64.94.100.204")
	if err == nil {
		t.Log(succ, len(list))
	} else {
		t.Error(err.Error())
	}
}

func TestUpToDateCheck(t *testing.T) {
	apps := getTestApps()
	succ, info, err := apps.UpToDateCheck(580, "123")
	if err == nil {
		t.Log(succ, info)
	} else {
		t.Error(err.Error())
	}
}
