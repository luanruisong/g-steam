package isteam_news

import (
	"fmt"
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() ISteamNews {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetNewsForApp(t *testing.T) {
	fmt.Println(getTestApps().GetNewsForApp(440, 0, 0, ""))
}
