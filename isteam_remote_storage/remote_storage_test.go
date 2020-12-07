package isteam_news

import (
	"fmt"
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() *iSteamRemoteStorage {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetCollectionDetails(t *testing.T) {
	app := getTestApps()

	fmt.Println(app.GetCollectionDetails(10, []uint64{1, 2, 3}))

}

func TestGetPublishedFileDetails(t *testing.T) {
	app := getTestApps()

	fmt.Println(app.GetPublishedFileDetails(10, []uint64{1, 2, 3}))

}

func TestGetUGCFileDetails(t *testing.T) {
	app := getTestApps()

	fmt.Println(app.GetUGCFileDetails(440, 1, 0))

}
