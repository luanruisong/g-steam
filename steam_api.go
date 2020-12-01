package steam

import "fmt"

const apiUrl = `http://api.steampowered.com/%s/%s/%s`

type Api interface {
	SetServer(string) Api
	SetWay(string) Api
	SetVersion(string) Api
	GetUrl() string
	Request() string
	Response() interface{}
}

type Steam struct {
	server,
	way,
	version,
	query string
}

func (s *Steam) GetUrl() string {
	panic("implement me")
}

func (s *Steam) SetServer(server string) Api {
	s.server = server
	return s
}

func (s *Steam) SetWay(way string) Api {
	s.way = way
	return s
}

func (s *Steam) SetVersion(version string) Api {
	s.version = version
	return s
}

func (s *Steam) Request() string {
	_ = fmt.Sprintf(apiUrl, s.server, s.way, s.version)
	return ""
}

func (s *Steam) Response() interface{} {
	panic("implement me")
}
