package isteam_user_stats

import (
	"testing"

	jsoniter "github.com/json-iterator/go"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() ISteamUserStats {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetGlobalAchievementPercentagesForApp(t *testing.T) {
	app := getTestApps()
	res, err := app.GetGlobalAchievementPercentagesForApp(440)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(len(res), res)
	}
}

func TestGetGlobalStatsForGame(t *testing.T) {
	app := getTestApps()
	//Example: http://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v0001/?format=xml&appid=17740&count=1&name[0]=global.map.emp_isle
	resc, res, err := app.GetGlobalStatsForGame(17740, []string{"global.map.emp_isle"}, 0, 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(resc, res)
	}
}

func TestGetNumberOfCurrentPlayers(t *testing.T) {
	app := getTestApps()
	//Example: http://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v0001/?format=xml&appid=17740&count=1&name[0]=global.map.emp_isle
	result, playerCount, err := app.GetNumberOfCurrentPlayers(440)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result, playerCount)
	}
}
func TestGetPlayerAchievements(t *testing.T) {
	app := getTestApps()
	//Example: http://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v0001/?format=xml&appid=17740&count=1&name[0]=global.map.emp_isle
	result, err := app.GetPlayerAchievements(440, "76561198421538055", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(jsoniter.MarshalToString(result))
	}
}

func TestGetSchemaForGame(t *testing.T) {
	app := getTestApps()
	//Example: http://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v0001/?format=xml&appid=17740&count=1&name[0]=global.map.emp_isle
	result, err := app.GetSchemaForGame(440, "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result.GameName, result.GameVersion)
		t.Log(len(result.AvailableGameStats.Stats))
		t.Log(len(result.AvailableGameStats.Achievements))
	}
}

func TestGetUserStatsForGame(t *testing.T) {
	app := getTestApps()
	//Example: http://api.steampowered.com/ISteamUserStats/GetGlobalStatsForGame/v0001/?format=xml&appid=17740&count=1&name[0]=global.map.emp_isle
	result, err := app.GetUserStatsForGame(440, "76561198421538055")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(jsoniter.MarshalToString(result))
	}
}
