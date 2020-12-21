package isteam_economy

import (
	"testing"

	steam "github.com/luanruisong/g-steam"
)

func getTestApps() ISteamEconomy {
	client := steam.NewClient("3C6A47B5B1E591DB30DA99B2E043571B")
	return New(client)
}

func TestGetAssetClassInfo(t *testing.T) {
	apps := getTestApps()
	succ, infoList, err := apps.GetAssetClassInfo(440, "en", 2, []uint64{195151, 16891096}, []uint64{})
	if err == nil {
		t.Log(succ, len(infoList))
	} else {
		t.Error(err.Error())
	}

}
func TestGetAssetPrices(t *testing.T) {
	apps := getTestApps()
	succ, list, err := apps.GetAssetPrices(440, "en", "USD")
	if err == nil {
		t.Log(succ, len(list), list[0].Prices, list[0].OriginalPrices)
	} else {
		t.Error(err.Error())
	}
}
