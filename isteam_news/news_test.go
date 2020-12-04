package isteam_news

import (
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() *iSteamNews {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetNewsForApp(t *testing.T) {
	getTestApps().GetNewsForApp(440, 0, 0, "")
}
