package iplayer_service

import "fmt"

func FmtImg(appid uint, id string) string {
	const base = "http://media.steampowered.com/steamcommunity/public/images/apps/%d/%s.jpg"
	return fmt.Sprintf(base, appid, id)
}
