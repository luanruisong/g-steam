package steam

import (
	"fmt"
	"net/http"
	"testing"
)

func newTestClient() *client {
	return NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
}

func steamId() string {
	return "76561198421538055"
}

func TestGetOpenId(t *testing.T) {
	client := newTestClient()

	render := client.RenderTo("http://127.0.0.1:9099/")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		res, err := client.OpenidBindQuery(request.URL.Query())
		fmt.Println(res, err)
		fmt.Println(res.GetSteamId())

	})
	fmt.Println(render)
	_ = http.ListenAndServe(":9099", nil)
}
func TestGetPlayerSummaries(t *testing.T) {
	//http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=XXX&steamids=76561197960435530
	client := newTestClient()
	api := client.Api()
	raw, err := api.Server("ISteamUser").
		Method("GetPlayerSummaries").
		Version("v0002").
		AddParam("steamids", steamId()).
		Get(nil)
	if err == nil {
		t.Log(raw)
	} else {
		t.Error(err.Error())
	}
}

func TestGetFriendList(t *testing.T) {
	//http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=XXX&steamid=76561197960435530&relationship=friend
	client := newTestClient()
	api := client.Api()
	raw, err := api.Server("ISteamUser").
		Method("GetFriendList").
		Version("v0001").
		AddParam("steamid", steamId()).
		AddParam("relationship", "friend").
		Get(nil)
	if err == nil {
		t.Log(raw)
	} else {
		t.Error(err.Error())
	}
}

func TestGetOwnedGames(t *testing.T) {
	//https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/?key=XXXXXX&steamid=76561197960434622&format=json
	client := newTestClient()
	api := client.Api()
	raw, err := api.Server("IPlayerService").
		Method("GetOwnedGames").
		Version("v1").
		AddParam("steamid", steamId()).
		Get(nil)
	if err == nil {
		t.Log(raw)
	} else {
		t.Error(err.Error())
	}
}

func TestGetRecentlyPlayedGames(t *testing.T) {
	client := newTestClient()
	api := client.Api()
	raw, err := api.Server("IPlayerService").
		Method("GetRecentlyPlayedGames").
		Version("v0001").
		AddParam("steamid", steamId()).
		Get(nil)
	if err == nil {
		t.Log(raw)
	} else {
		t.Error(err.Error())
	}
}

func TestProfile(t *testing.T) {
	//编辑地址
	//{player.profileurl}/edit/settings
}
