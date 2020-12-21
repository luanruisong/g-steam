package iplayer_service

import (
	"fmt"

	steam "github.com/luanruisong/g-steam"
)

const (
	PlayerServerName = "IPlayerService"
)

type (
	IPlayerService interface {
		GetRecentlyPlayedGames(steamid string, count uint) (uint, []playedGame, error)
		GetOwnedGames(steamid string, includeAppInfo, includePlayedFreeGames bool, appidsFilter []uint) (uint, []playedGame, error)
		GetSteamLevel(steamid string) (uint, error)
		GetBadges(steamid string) (badge, error)
		GetCommunityBadgeProgress(steamid string, badgeid uint) ([]Quest, error)
	}
	iPlayerService struct {
		c steam.Client
	}
	playedGame struct {
		Appid                    uint   `json:"appid" xml:"appid" form:"appid"`
		Name                     string `json:"name" xml:"name" form:"name"`
		Playtime2weeks           uint   `json:"playtime_2weeks" xml:"playtime_2weeks" form:"playtime_2weeks"`
		PlaytimeForever          uint   `json:"playtime_forever" xml:"playtime_forever" form:"playtime_forever"`
		ImgIconUrl               string `json:"img_icon_url" xml:"img_icon_url" form:"img_icon_url"`
		ImgLogoUrl               string `json:"img_logo_url" xml:"img_logo_url" form:"img_logo_url"`
		HasCommunityVisibleStats bool   `json:"has_community_visible_stats" xml:"has_community_visible_stats" form:"has_community_visible_stats"`
	}

	badgesInfo struct {
		Badgeid        uint `json:"badgeid"`
		Level          uint `json:"level"`
		CompletionTime uint `json:"completion_time"`
		Xp             uint `json:"xp"`
		Scarcity       uint `json:"scarcity"`
	}
	badge struct {
		Badges                     []badgesInfo `json:"badges"`
		PlayerXp                   uint         `json:"player_xp"`
		PlayerLevel                uint         `json:"player_level"`
		PlayerXpNeededToLevelUp    uint         `json:"player_xp_needed_to_level_up"`
		PlayerXpNeededCurrentLevel uint         `json:"player_xp_needed_current_level"`
	}
	Quest struct {
		Questid   int  `json:"questid"`
		Completed bool `json:"completed"`
	}
)

func (app *iPlayerService) apiServer() steam.Api {
	return app.c.Api().Server(PlayerServerName)
}

func (app *iPlayerService) GetRecentlyPlayedGames(steamid string, count uint) (uint, []playedGame, error) {
	api := app.apiServer().
		Method("GetRecentlyPlayedGames").
		Version("v1").
		AddParam("steamid", steamid)
	if count > 0 {
		api = api.AddParam("count", count)
	}
	var res struct {
		Response struct {
			TotalCount uint         `json:"total_count" xml:"total_count" form:"total_count"`
			Games      []playedGame `json:"games" xml:"games" form:"games"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.TotalCount, res.Response.Games, err
}

func (app *iPlayerService) GetOwnedGames(steamid string, includeAppInfo, includePlayedFreeGames bool, appidsFilter []uint) (uint, []playedGame, error) {
	api := app.apiServer().
		Method("GetOwnedGames").
		Version("v1").
		AddParam("steamid", steamid).
		AddParam("include_appinfo", includeAppInfo).
		AddParam("include_played_free_games", includePlayedFreeGames)

	for i, v := range appidsFilter {
		api = api.AddParam(fmt.Sprintf("appids_filter[%d]", i), v)
	}

	var res struct {
		Response struct {
			GameCount uint         `json:"game_count" xml:"game_count" form:"game_count"`
			Games     []playedGame `json:"games" xml:"games" form:"games"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.GameCount, res.Response.Games, err
}

func (app *iPlayerService) GetSteamLevel(steamid string) (uint, error) {
	api := app.apiServer().
		Method("GetSteamLevel").
		Version("v1").
		AddParam("steamid", steamid)
	var res struct {
		Response struct {
			PlayerLevel uint `json:"player_level" xml:"player_level" form:"player_level"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.PlayerLevel, err
}

func (app *iPlayerService) GetBadges(steamid string) (badge, error) {
	api := app.apiServer().
		Method("GetBadges").
		Version("v1").
		AddParam("steamid", steamid)
	var res struct {
		Response badge `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response, err
}

func (app *iPlayerService) GetCommunityBadgeProgress(steamid string, badgeid uint) ([]Quest, error) {
	api := app.apiServer().
		Method("GetCommunityBadgeProgress").
		Version("v1").
		AddParam("steamid", steamid).
		AddParam("badgeid", badgeid)
	var res struct {
		Response struct {
			Quests []Quest `json:"quests" xml:"quests" form:"quests"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Quests, err
}

func New(c steam.Client) IPlayerService {
	return &iPlayerService{c: c}
}

func (pg playedGame) GetIcon() string {
	return FmtImg(pg.Appid, pg.ImgIconUrl)
}
func (pg playedGame) GetLogo() string {
	return FmtImg(pg.Appid, pg.ImgLogoUrl)
}
