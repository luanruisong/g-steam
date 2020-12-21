package isteam_user_stats

import (
	"fmt"

	steam "github.com/luanruisong/g-steam"
)

const (
	UserStatsServerName = "ISteamUserStats"
)

type (
	ISteamUserStats interface {
		GetGlobalAchievementPercentagesForApp(gameId uint) ([]AchievementInfo, error)
		GetGlobalStatsForGame(appId uint, names []string, startDate, endDate int64) (uint, map[string]GlobalState, error)
		GetNumberOfCurrentPlayers(appId uint) (uint, uint, error)
		GetPlayerAchievements(appId uint, steamid, l string) (PlayerStats, error)
		GetSchemaForGame(appId uint, l string) (GameSchema, error)
		GetUserStatsForGame(appId uint, steamid string) (SinglePlayerStats, error)
	}
	iSteamUserStats struct {
		c steam.Client
	}
	AchievementInfo struct {
		Name    string  `json:"name" xml:"name" form:"name"`
		Percent float64 `json:"percent" xml:"percent" form:"percent"`
	}
	GlobalState struct {
		Total string `json:"total" xml:"total" form:"total"`
	}
	AchievementsInfo struct {
		Apiname    string `json:"apiname" xml:"apiname" form:"apiname"`
		Achieved   uint   `json:"achieved" xml:"achieved" form:"achieved"`
		Unlocktime int64  `json:"unlocktime" xml:"unlocktime" form:"unlocktime"`
	}
	SinglePlayerStats struct {
		SteamId  string `json:"steamID" xml:"steamID" form:"steamID"`
		GameName string `json:"gameName" xml:"gameName" form:"gameName"`
	}
	PlayerStats struct {
		SinglePlayerStats
		Success      bool               `json:"success" xml:"success" form:"success"`
		Achievements []AchievementsInfo `json:"achievements" xml:"achievements" form:"achievements"`
	}
	SchemaStats struct {
		Name         string `json:"name" xml:"name" form:"name"`
		DefaultValue int    `json:"defaultvalue" xml:"defaultvalue" form:"defaultvalue"`
		DisplayName  string `json:"displayName" xml:"displayName" form:"displayName"`
	}
	SchemaAchievements struct {
		Name         string `json:"name" xml:"name" form:"name"`
		DefaultValue int    `json:"defaultvalue" xml:"defaultvalue" form:"defaultvalue"`
		DisplayName  string `json:"displayName" xml:"displayName" form:"displayName"`
		Hidden       int    `json:"hidden" xml:"hidden" form:"hidden"`
		Description  string `json:"description" xml:"description" form:"description"`
		Icon         string `json:"icon" xml:"icon" form:"icon"`
		Icongray     string `json:"icongray" xml:"icongray" form:"icongray"`
	}
	GameSchema struct {
		GameName           string `json:"gameName" xml:"gameName" form:"gameName"`
		GameVersion        string `json:"gameVersion" xml:"gameVersion" form:"gameVersion"`
		AvailableGameStats struct {
			Stats        []SchemaStats        `json:"stats" xml:"stats" form:"stats"`
			Achievements []SchemaAchievements `json:"achievements" xml:"achievements" form:"achievements"`
		} `json:"availableGameStats" xml:"availableGameStats" form:"availableGameStats"`
	}
)

func (app *iSteamUserStats) apiServer() steam.Api {
	return app.c.Api().Server(UserStatsServerName)
}

func (app *iSteamUserStats) GetGlobalAchievementPercentagesForApp(gameId uint) ([]AchievementInfo, error) {
	api := app.apiServer().
		Method("GetGlobalAchievementPercentagesForApp").
		Version("v0002").
		AddParam("gameid", gameId)
	var res struct {
		AchievementPercenTages struct {
			Achievements []AchievementInfo `json:"achievements" xml:"achievements" form:"achievements"`
		} `json:"achievementpercentages" xml:"achievementpercentages" form:"achievementpercentages"`
	}
	_, err := api.Get(&res)
	return res.AchievementPercenTages.Achievements, err
}

func (app *iSteamUserStats) GetGlobalStatsForGame(appId uint, names []string, startDate, endDate int64) (uint, map[string]GlobalState, error) {
	api := app.apiServer().
		Method("GetGlobalStatsForGame").
		Version("v0001").
		AddParam("appid", appId).
		AddParam("count", len(names))

	for i, v := range names {
		api = api.AddParam(fmt.Sprintf("name[%d]", i), v)
	}
	if startDate > 0 {
		api = api.AddParam("startdate", startDate)
	}
	if endDate > 0 {
		api = api.AddParam("enddate", endDate)
	}

	var res struct {
		Response struct {
			Result      uint                   `json:"result" xml:"result" form:"result"`
			GlobalStats map[string]GlobalState `json:"globalstats" xml:"globalstats" form:"globalstats"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Result, res.Response.GlobalStats, err
}

func (app *iSteamUserStats) GetNumberOfCurrentPlayers(appId uint) (uint, uint, error) {
	api := app.apiServer().
		Method("GetNumberOfCurrentPlayers").
		Version("v1").
		AddParam("appid", appId)
	var res struct {
		Response struct {
			Result      uint `json:"result" xml:"result" form:"result"`
			PlayerCount uint `json:"player_count" xml:"player_count" form:"player_count"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Result, res.Response.PlayerCount, err
}

func (app *iSteamUserStats) GetPlayerAchievements(appId uint, steamid, l string) (PlayerStats, error) {
	api := app.apiServer().
		Method("GetPlayerAchievements").
		Version("v1").
		AddParam("appid", appId).AddParam("steamid", steamid)
	if len(l) > 0 {
		api = api.AddParam("l", l)
	}
	var res struct {
		PlayerStats PlayerStats `json:"playerstats" xml:"playerstats" form:"playerstats"`
	}
	_, err := api.Get(&res)
	return res.PlayerStats, err
}

func (app *iSteamUserStats) GetSchemaForGame(appId uint, l string) (GameSchema, error) {
	api := app.apiServer().
		Method("GetSchemaForGame").
		Version("v2").
		AddParam("appid", appId)
	if len(l) > 0 {
		api = api.AddParam("l", l)
	}
	var res struct {
		Game GameSchema `json:"game" xml:"game" form:"game"`
	}
	_, err := api.Get(&res)
	return res.Game, err
}

func (app *iSteamUserStats) GetUserStatsForGame(appId uint, steamid string) (SinglePlayerStats, error) {
	api := app.apiServer().
		Method("GetUserStatsForGame").
		Version("v2").
		AddParam("appid", appId).AddParam("steamid", steamid)
	var res struct {
		PlayerStats SinglePlayerStats `json:"playerstats" xml:"playerstats" form:"playerstats"`
	}
	_, err := api.Get(&res)
	return res.PlayerStats, err
}

func New(c steam.Client) ISteamUserStats {
	return &iSteamUserStats{c: c}
}
