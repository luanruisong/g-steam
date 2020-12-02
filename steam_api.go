package steam

import (
	"errors"
	"fmt"
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

const apiUrl = `http://api.steampowered.com/%s/%s/%s/`

type api struct {
	param url.Values
	req   Req
	appKey,
	server,
	method,
	version string
}

func openApi(appKey string, req Req) *api {
	return &api{
		req:    req,
		appKey: appKey,
		param:  url.Values{},
	}
}

func (s *api) Server(server string) *api {
	s.server = server
	return s
}

func (s *api) Method(method string) *api {
	s.method = method
	return s
}

func (s *api) Version(version string) *api {
	s.version = version
	return s
}

func (s *api) AddParam(k, v string) *api {
	s.param.Add(k, v)
	return s
}

func (s *api) AddParams(m map[string]string) *api {
	for i, v := range m {
		s.param.Add(i, v)
	}
	return s
}

func (s *api) buildUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf(apiUrl, s.server, s.method, s.version))
}

func (s *api) valid() error {
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

func (s *api) do(d doFunc, resPtr interface{}) (raw string, err error) {
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

	raw, err = d(u.String(), nil, nil)
	if err == nil && len(raw) > 0 && resPtr != nil {
		err = jsoniter.UnmarshalFromString(raw, resPtr)
	}
	return
}

func (s *api) Get(resPtr interface{}) (raw string, err error) {
	return s.do(s.req.Get, resPtr)
}

func (s *api) Post(resPtr interface{}) (raw string, err error) {
	return s.do(s.req.Post, resPtr)
}