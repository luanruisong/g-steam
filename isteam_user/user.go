package isteam_user

import (
	"strings"

	steam "github.com/luanruisong/g-steam"
)

const (
	UserServerName = "ISteamUser"
)

type (
	ISteamUser interface {
		GetFriendList(steamId string, relationship string) ([]SteamFriend, error)
		GetPlayerBans(steamIds ...string) ([]PlayerBans, error)
		GetPlayerSummaries(steamIds ...string) ([]UserProfile, error)
		GetUserGroupList(steamId string) (bool, []GroupInfo, error)
		ResolveVanityURL(vanityurl string) (bool, string, string, error)
	}
	iSteamUser struct {
		c steam.Client
	}

	SteamFriend struct {
		SteamId      string `json:"steamid" xml:"steamid" form:"steamid"`
		Relationship string `json:"relationship" xml:"relationship" form:"relationship"`
		FriendSince  int64  `json:"friend_since" xml:"friend_since" form:"friend_since"`
	}
	PlayerBans struct {
		SteamId          string `json:"SteamId" xml:"SteamId" form:"SteamId"`
		CommunityBanned  bool   `json:"CommunityBanned" xml:"CommunityBanned" form:"CommunityBanned"`
		VACBanned        bool   `json:"VACBanned" xml:"VACBanned" form:"VACBanned"`
		DaysSinceLastBan uint   `json:"DaysSinceLastBan" xml:"DaysSinceLastBan" form:"DaysSinceLastBan"`
		NumberOfVACBans  uint   `json:"NumberOfVACBans" xml:"NumberOfVACBans" form:"NumberOfVACBans"`
		NumberOfGameBans uint   `json:"NumberOfGameBans" xml:"NumberOfGameBans" form:"NumberOfGameBans"`
		EconomyBan       string `json:"EconomyBan" xml:"EconomyBan" form:"EconomyBan"`
	}
	UserProfile struct {
		SteamId                  string `json:"steamid" xml:"steamid" form:"steamid"`
		CommunityVisibilityState uint8  `json:"communityvisibilitystate" xml:"communityvisibilitystate" form:"communityvisibilitystate"`
		ProfileState             uint8  `json:"profilestate" xml:"profilestate" form:"profilestate"`
		PersonaName              string `json:"personaname" xml:"personaname" form:"personaname"`
		LastLogOff               int64  `json:"lastlogoff" xml:"lastlogoff" form:"lastlogoff"`
		ProfileUrl               string `json:"profileurl" xml:"profileurl" form:"profileurl"`
		Avatar                   string `json:"avatar" xml:"avatar" form:"avatar"`
		AvatarMedium             string `json:"avatarmedium" xml:"avatarmedium" form:"avatarmedium"`
		AvatarFull               string `json:"avatarfull" xml:"avatarfull" form:"avatarfull"`
		PersonaState             uint8  `json:"personastate" xml:"personastate" form:"personastate"`
		CommentPermission        uint8  `json:"commentpermission" xml:"commentpermission" form:"commentpermission"`
		RealName                 string `json:"realname" xml:"realname" form:"realname"`
		PrimaryClanid            string `json:"primaryclanid" xml:"primaryclanid" form:"primaryclanid"`
		TimeCreated              int64  `json:"timecreated" xml:"timecreated" form:"timecreated"`
		LocCountryCode           string `json:"loccountrycode" xml:"loccountrycode" form:"loccountrycode"`
		LocStateCode             string `json:"locstatecode" xml:"locstatecode" form:"locstatecode"`
		LocCityId                uint   `json:"loccityid" xml:"loccityid" form:"loccityid"`
		GameId                   string `json:"gameid" xml:"gameid" form:"gameid"`
		GameExtraInfo            string `json:"gameextrainfo" xml:"gameextrainfo" form:"gameextrainfo"`
		GameServerIp             string `json:"gameserverip" xml:"gameserverip" form:"gameserverip"`
	}
	GroupInfo struct {
		Gid string `json:"gid" xml:"gid" form:"gid"`
	}
)

func (app *iSteamUser) apiServer() steam.Api {
	return app.c.Api().Server(UserServerName)
}

func (app *iSteamUser) GetFriendList(steamId string, relationship string) ([]SteamFriend, error) {
	api := app.apiServer().
		Method("GetFriendList").
		Version("v1").
		AddParam("steamid", steamId).
		AddParam("relationship", relationship)
	var res struct {
		Friendslist struct {
			Friends []SteamFriend `json:"friends" xml:"friends" form:"friends"`
		} `json:"friendslist" xml:"friendslist" form:"friendslist"`
	}
	_, err := api.Get(&res)
	return res.Friendslist.Friends, err
}

func (app *iSteamUser) GetPlayerBans(steamIds ...string) ([]PlayerBans, error) {
	api := app.apiServer().
		Method("GetPlayerBans").
		Version("v1").
		AddParam("steamids", strings.Join(steamIds, ","))
	var res struct {
		Players []PlayerBans `json:"players" xml:"players" form:"players"`
	}
	_, err := api.Get(&res)
	return res.Players, err
}

func (app *iSteamUser) GetPlayerSummaries(steamIds ...string) ([]UserProfile, error) {
	api := app.apiServer().
		Method("GetPlayerSummaries").
		Version("v0002").
		AddParam("steamids", strings.Join(steamIds, ","))
	var res struct {
		Response struct {
			Players []UserProfile `json:"players" xml:"players" form:"players"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Players, err
}

func (app *iSteamUser) GetUserGroupList(steamId string) (bool, []GroupInfo, error) {
	api := app.apiServer().
		Method("GetUserGroupList").
		Version("v1").
		AddParam("steamid", steamId)
	var res struct {
		Response struct {
			Success bool        `json:"success" xml:"success" form:"success"`
			Groups  []GroupInfo `json:"groups" xml:"groups" form:"groups"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Success, res.Response.Groups, err
}

func (app *iSteamUser) ResolveVanityURL(vanityurl string) (bool, string, string, error) {
	api := app.apiServer().
		Method("ResolveVanityURL").
		Version("v0001").
		AddParam("vanityurl", vanityurl)
	var res struct {
		Response struct {
			Success int    `json:"success" xml:"success" form:"success"`
			Steamid string `json:"steamid" xml:"steamid" form:"steamid"`
			Message string `json:"message" xml:"message" form:"message"`
		} `json:"response" xml:"response" form:"response"`
	}
	_, err := api.Get(&res)
	return res.Response.Success == 1, res.Response.Steamid, res.Response.Message, err
}

func New(c steam.Client) ISteamUser {
	return &iSteamUser{c: c}
}
