package isteam_user

import (
	"testing"

	jsoniter "github.com/json-iterator/go"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() ISteamUser {
	client := steam.NewClient("24BA3E91FBD8C1C9678907A552C0AD37")
	return New(client)
}

func TestGetFriendList(t *testing.T) {
	app := getTestApps()
	res, err := app.GetFriendList("76561198421538055", "all")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(res)
	}
}

func TestGetPlayerBans(t *testing.T) {
	app := getTestApps()
	res, err := app.GetPlayerBans("76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(res)
	}
}

func TestGetPlayerSummaries(t *testing.T) {
	app := getTestApps()
	res, err := app.GetPlayerSummaries("76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(res)
		b, _ := jsoniter.MarshalIndent(res, "", "")
		t.Log(string(b))
	}
}

func TestGetUserGroupList(t *testing.T) {
	app := getTestApps()
	succ, res, err := app.GetUserGroupList("76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(succ, res)
	}
}
func TestResolveVanityURL(t *testing.T) {
	app := getTestApps()
	succ, steamId, msg, err := app.ResolveVanityURL("userVanityUrlName")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(succ, steamId, msg)
	}
}
