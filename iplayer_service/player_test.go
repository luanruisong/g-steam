package iplayer_service

import (
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() IPlayerService {
	client := steam.NewClient("24BA3E91FBD8C1C9678907A552C0AD37")
	return New(client)
}

func TestGetRecentlyPlayedGames(t *testing.T) {
	app := getTestApps()
	count, res, err := app.GetRecentlyPlayedGames("76561198421538055", 10)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(count, res)
	}
}

func TestGetOwnedGames(t *testing.T) {
	app := getTestApps()
	count, res, err := app.GetOwnedGames("76561199110641233", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(count, res)
	}
}

func TestGetOwnedGamesExtend(t *testing.T) {
	app := getTestApps()
	res, err := app.GetOwnedGamesExtend("76561198421538055", true, false, nil)
	//res, err := app.GetOwnedGamesExtend("76561199110641233", true, false, nil)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Visible", res.Visible, "count", res.GameCount, "games", res.Games)
	}
}

func TestGetSteamLevel(t *testing.T) {
	app := getTestApps()
	level, err := app.GetSteamLevel("76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(level)
	}
}

func TestGetBadges(t *testing.T) {
	app := getTestApps()
	level, err := app.GetBadges("76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(level)
	}
}

func TestGetCommunityBadgeProgress(t *testing.T) {
	app := getTestApps()
	level, err := app.GetCommunityBadgeProgress("76561198421538055", 2)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(level)
	}
}
