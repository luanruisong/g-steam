package steam

import (
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
		AddParam(k, v string) Api
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

func (s *defApi) AddParam(k, v string) Api {
	s.param.Add(k, v)
	return s
}

func (s *defApi) AddParams(m map[string]string) Api {
	for i, v := range m {
		s.param.Add(i, v)
	}
	return s
}

func (s *defApi) buildUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf(apiUrl, s.server, s.method, s.version))
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

func (s *defApi) do(resPtr interface{}) (raw string, err error) {
	if err = s.valid(); err != nil {
		return
	}
	var u *url.URL
	if u, err = s.buildUrl(); err != nil {
		return
	}
	if _, ok := s.param["key"]; !ok {
		s.param.Add("key", s.appKey)
	}
	u.RawQuery = s.param.Encode()

	raw, err = s.req.Get(u.String())
	if err == nil && len(raw) > 0 && resPtr != nil {
		err = jsoniter.UnmarshalFromString(raw, resPtr)
	}
	return
}

func (s *defApi) Get(resPtr interface{}) (raw string, err error) {
	return s.do(resPtr)
}
