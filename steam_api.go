package steam

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

const apiUrl = `http://api.steampowered.com/%s/%s/%s/`

type (
	Api interface {
		Server(server string) Api
		Method(method string) Api
		Version(version string) Api
		AddParam(k string, v interface{}) Api
		AddParams(m map[string]interface{}) Api
		Url() string
		Get(resPtr interface{}) (raw string, err error)
	}
	defApi struct {
		param url.Values
		req   Req
		appKey,
		server,
		method,
		version string
	}
)

func openApi(appKey string, req Req) Api {
	return &defApi{
		req:    req,
		appKey: appKey,
		param:  url.Values{},
	}
}

func (s *defApi) Server(server string) Api {
	s.server = server
	return s
}

func (s *defApi) Method(method string) Api {
	s.method = method
	return s
}

func (s *defApi) Version(version string) Api {
	s.version = version
	return s
}

func (s *defApi) AddParam(k string, v interface{}) Api {
	s.param.Add(k, fmt.Sprintf("%v", v))
	return s
}

func (s *defApi) AddParams(m map[string]interface{}) Api {
	for i, v := range m {
		s.AddParam(i, v)
	}
	return s
}

func (s *defApi) buildUrl() (u *url.URL, err error) {
	if u, err = url.Parse(fmt.Sprintf(apiUrl, s.server, s.method, s.version)); err == nil {
		if _, ok := s.param["key"]; !ok {
			s.param.Add("key", s.appKey)
		}
		u.RawQuery = s.param.Encode()
	}
	return
}

func (s *defApi) Url() string {
	u, _ := s.buildUrl()
	return u.String()
}

func (s *defApi) valid() error {
	if len(s.server) == 0 ||
		len(s.method) == 0 ||
		len(s.version) == 0 {
		return errors.New("param error")
	}
	if len(s.appKey) == 0 {
		return errors.New("appkey is nil")
	}
	return nil
}

func (s *defApi) Get(resPtr interface{}) (raw string, err error) {
	if err = s.valid(); err != nil {
		return
	}
	var u *url.URL
	if u, err = s.buildUrl(); err != nil {
		return
	}
	raw, err = s.req.Get(u.String())
	if err == nil && len(raw) > 0 && resPtr != nil {
		var decoder func(string, interface{}) error

		switch u.Query().Get("format") {
		case "xml":
			decoder = func(s string, i interface{}) error {
				return xml.Unmarshal([]byte(s), i)
			}
		default:
			decoder = jsoniter.UnmarshalFromString
		}
		err = decoder(raw, resPtr)
	}
	return
}
