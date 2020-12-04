package isteam_economy

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	steam "github.com/luanruisong/g-steam"
)

const (
	EconomyServerName = "ISteamEconomy"
)

type (
	iSteamEconomy struct {
		api steam.Api
	}
	action struct {
		Name string `json:"name" xml:"name" form:"name"`
		Link string `json:"link" xml:"link" form:"link"`
	}
	tag struct {
		InternalName string `json:"internal_name" xml:"internal_name" form:"internal_name"`
		Name         string `json:"name" xml:"name" form:"name"`
		Category     string `json:"category" xml:"category" form:"category"`
		Color        string `json:"color" xml:"color" form:"color"`
		CategoryName string `json:"category_name" xml:"category_name" form:"category_name"`
	}
	assetClassInfo struct {
		Classid                     string            `json:"classid" xml:"classid" form:"classid"`
		IconUrl                     string            `json:"icon_url" xml:"icon_url" form:"icon_url"`
		IconUrlLarge                string            `json:"icon_url_large" xml:"icon_url_large" form:"icon_url_large"`
		IconDragLarge               string            `json:"icon_drag_large" xml:"icon_drag_large" form:"icon_drag_large"`
		Name                        string            `json:"name" xml:"name" form:"name"`
		MarketName                  string            `json:"market_name" xml:"market_name" form:"market_name"`
		MarketHashName              string            `json:"market_hash_name" xml:"market_hash_name" form:"market_hash_name"`
		NameColor                   string            `json:"name_color" xml:"name_color" form:"name_color"`
		BackgroundColor             string            `json:"background_color" xml:"background_color" form:"background_color"`
		Type                        string            `json:"type" xml:"type" form:"type"`
		Tradable                    string            `json:"tradable" xml:"tradable" form:"tradable"`
		Marketable                  string            `json:"marketable" xml:"marketable" form:"marketable"`
		Commodity                   string            `json:"commodity" xml:"commodity" form:"commodity"`
		MarketTradableRestriction   string            `json:"market_tradable_restriction" xml:"market_tradable_restriction" form:"market_tradable_restriction"`
		MarketMarketableRestriction string            `json:"market_marketable_restriction" xml:"market_marketable_restriction" form:"market_marketable_restriction"`
		Fraudwarnings               string            `json:"fraudwarnings" xml:"fraudwarnings" form:"fraudwarnings"`
		Actions                     map[string]action `json:"actions" xml:"actions" form:"actions"`
		MarketActions               map[string]action `json:"market_actions" xml:"market_actions" form:"market_actions"`
		Tags                        map[string]tag    `json:"market_actions" xml:"market_actions" form:"market_actions"`
		//appdata
		AppData struct {
			DefIndex   string `json:"def_index" xml:"def_index" form:"def_index"`
			Quality    string `json:"quality" xml:"quality" form:"quality"`
			Slot       string `json:"slot" xml:"slot" form:"slot"`
			FilterData map[string]struct {
				ElementIds map[string]string `json:"element_ids" xml:"element_ids" form:"element_ids"`
			} `json:"filter_data" xml:"filter_data" form:"filter_data"`
			PlayerClassIds map[string]string `json:"player_class_ids" xml:"player_class_ids" form:"player_class_ids"`
		} `json:"app_data" xml:"app_data" form:"app_data"`
	}

	assetPriceInfo struct {
		Prices         map[string]uint `json:"prices" xml:"prices" form:"prices"`
		OriginalPrices map[string]uint `json:"original_prices" xml:"original_prices" form:"original_prices"`
		Name           string          `json:"name" xml:"name" form:"name"`
		Date           string          `json:"date" xml:"date" form:"date"`
		Class          []struct {
			Name  string `json:"name" xml:"name" form:"name"`
			Value string `json:"value" xml:"value" form:"value"`
		} `json:"class" xml:"class" form:"class"`
		ClassId string   `json:"class_id" xml:"class_id" form:"class_id"`
		Tags    []string `json:"tags" xml:"tags" form:"tags"`
		TagIds  []uint64 `json:"tag_ids" xml:"tag_ids" form:"tag_ids"`
	}
)

func (app *iSteamEconomy) apiServer() steam.Api {
	return app.api.Server(EconomyServerName)
}

func (app *iSteamEconomy) GetAssetClassInfo(appid uint, language string, classCount uint, classId, instanceid []uint64) (succ bool, m map[string]assetClassInfo, err error) {
	var tmp map[string]interface{}
	api := app.apiServer().Method("GetAssetClassInfo").Version("v0001").
		AddParam("appid", appid).
		AddParam("class_count", classCount)
	if len(language) > 0 {
		api = api.AddParam("language", language)
	}
	for i, v := range classId {
		api = api.AddParam(fmt.Sprintf("classid%d", i), v)
	}
	for i, v := range instanceid {
		api = api.AddParam(fmt.Sprintf("instanceid%d", i), v)
	}
	_, err = api.Get(&tmp)
	if err == nil {
		if x, ok := tmp["result"]; ok {
			result, ok := x.(map[string]interface{})
			if ok {
				if succ, ok = result["success"].(bool); ok && succ {
					delete(result, "success")
					str, _ := jsoniter.MarshalToString(result)
					_ = jsoniter.UnmarshalFromString(str, &m)
				}
			}

		}
	}
	return
}

func (app *iSteamEconomy) GetAssetPrices(appid uint, language, currency string) (bool, []assetPriceInfo, error) {
	api := app.apiServer().Method("GetAssetPrices").Version("v0001").AddParam("appid", appid)
	if len(language) > 0 {
		api = api.AddParam("language", language)
	}
	if len(currency) > 0 {
		api = api.AddParam("currency", currency)
	}
	var res struct {
		Result struct {
			Success bool             `json:"success" xml:"success" form:"success"`
			Assets  []assetPriceInfo `json:"assets" xml:"assets" form:"assets"`
		} `json:"result" xml:"result" form:"result"`
	}
	_, err := api.Get(&res)
	return res.Result.Success, res.Result.Assets, err
}
func New(c steam.Client) *iSteamEconomy {
	return &iSteamEconomy{api: c.Api()}
}
