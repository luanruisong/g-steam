package steam

import (
	"fmt"
	"net/http"
	"testing"
)

func newTestClient() *client {
	return NewClient("appkey")
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
		AddParam("steamids", "123").
		Get(nil)
	fmt.Println(raw, err)
}
