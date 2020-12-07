package isteam_user_stats

import (
	"fmt"

	steam "github.com/luanruisong/g-steam"
)

const (
	UserStatsServerName = "ISteamUserStats"
)

type (
	iSteamUserStats struct {
		api steam.Api
	}
	achievementInfo struct {
		Name    string  `json:"name" xml:"name" form:"name"`
		Percent float64 `json:"percent" xml:"percent" form:"percent"`
	}
	globalState struct {
		Total string `json:"total" xml:"total" form:"total"`
	}
	achievementsInfo struct {
		Apiname    string `json:"apiname" xml:"apiname" form:"apiname"`
		Achieved   uint   `json:"achieved" xml:"achieved" form:"achieved"`
		Unlocktime int64  `json:"unlocktime" xml:"unlocktime" form:"unlocktime"`
	}
	singlePlayerStats struct {
		SteamId  string `json:"steamID" xml:"steamID" form:"steamID"`
		GameName string `json:"gameName" xml:"gameName" form:"gameName"`
	}
	playerStats struct {
		singlePlayerStats
		Success      bool               `json:"success" xml:"success" form:"success"`
		Achievements []achievementsInfo `json:"achievements" xml:"achievements" form:"achievements"`
	}
	schemaStats struct {
		Name         string `json:"name" xml:"name" form:"name"`
		DefaultValue int    `json:"defaultvalue" xml:"defaultvalue" form:"defaultvalue"`
		DisplayName  string `json:"displayName" xml:"displayName" form:"displayName"`
	}
	schemaAchievements struct {
		Name         string `json:"name" xml:"name" form:"name"`
		DefaultValue int    `json:"defaultvalue" xml:"defaultvalue" form:"defaultvalue"`
		DisplayName  string `json:"displayName" xml:"displayName" form:"displayName"`
		Hidden       int    `json:"hidden" xml:"hidden" form:"hidden"`
		Description  string `json:"description" xml:"description" form:"description"`
		Icon         string `json:"icon" xml:"icon" form:"icon"`
		Icongray     string `json:"icongray" xml:"icongray" form:"icongray"`
	}
	gameSchema struct {
		GameName           string `json:"gameName" xml:"gameName" form:"gameName"`
		GameVersion        string `json:"gameVersion" xml:"gameVersion" form:"gameVersion"`
		AvailableGameStats struct {
			Stats        []schemaStats        `json:"stats" xml:"stats" form:"stats"`
			Achievements []schemaAchievements `json:"achievements" xml:"achievements" form:"achievements"`
		} `json:"availableGameStats" xml:"availableGameStats" form:"availableGameStats"`
	}
)

func (app *iSteamUserStats) apiServer() steam.Api {
	return app.api.Server(UserStatsServerName)
}

func (app *iSteamUserStats) GetGlobalAchievementPercentagesForApp(gameId uint) ([]achievementInfo, error) {
	api := app.apiServer().
		Method("GetGlobalAchievementPercentagesForApp").
		Version("v0002").
		AddParam("gameid", gameId)
	var res struct {
		AchievementPercenTages struct {
			Achievements []achievementInfo `json:"achievements" xml:"achievements" form:"achievements"`
		} `json:"achievementpercentages" xml:"achievementpercentages" form:"achievementpercentages"`
	}
	_, err := api.Get(&res)
	return res.AchievementPercenTages.Achievements, err
}

func (app *iSteamUserStats) GetGlobalStatsForGame(appId uint, names []string, startDate, endDate int64) (uint, map[string]globalState, error) {
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
			GlobalStats map[string]globalState `json:"globalstats" xml:"globalstats" form:"globalstats"`
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

func (app *iSteamUserStats) GetPlayerAchievements(appId uint, steamid, l string) (playerStats, error) {
	api := app.apiServer().
		Method("GetPlayerAchievements").
		Version("v1").
		AddParam("appid", appId).AddParam("steamid", steamid)
	if len(l) > 0 {
		api = api.AddParam("l", l)
	}
	var res struct {
		PlayerStats playerStats `json:"playerstats" xml:"playerstats" form:"playerstats"`
	}
	_, err := api.Get(&res)
	return res.PlayerStats, err
}

func (app *iSteamUserStats) GetSchemaForGame(appId uint, l string) (gameSchema, error) {
	api := app.apiServer().
		Method("GetSchemaForGame").
		Version("v2").
		AddParam("appid", appId)
	if len(l) > 0 {
		api = api.AddParam("l", l)
	}
	var res struct {
		Game gameSchema `json:"game" xml:"game" form:"game"`
	}
	_, err := api.Get(&res)
	return res.Game, err
}

func (app *iSteamUserStats) GetUserStatsForGame(appId uint, steamid string) (singlePlayerStats, error) {
	api := app.apiServer().
		Method("GetUserStatsForGame").
		Version("v2").
		AddParam("appid", appId).AddParam("steamid", steamid)
	var res struct {
		PlayerStats singlePlayerStats `json:"playerstats" xml:"playerstats" form:"playerstats"`
	}
	_, err := api.Get(&res)
	return res.PlayerStats, err
}

func New(c steam.Client) *iSteamUserStats {
	return &iSteamUserStats{api: c.Api()}
}
