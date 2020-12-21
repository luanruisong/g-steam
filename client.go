package steam

import (
	"net/url"
)

type (
	Client interface {
		SetReq(req Req)
		RenderTo(callback string) string
		MaybeRenderTo(url, callback string) string
		OpenidBindQuery(param url.Values) (res *openidRes, err error)
		OpenidBindMap(param map[string]string) (res *openidRes, err error)
		Api() Api
	}

	client struct {
		appKey string
		req    Req
	}
)

func NewClient(appKey string) *client {
	return &client{
		appKey: appKey,
		req:    NewDefReq(),
	}
}

func (c *client) SetReq(req Req) {
	c.req = req
}

func (c *client) RenderTo(callback string) string {
	return renderTo(callback)
}

// maybe open other to login steam
func (c *client) MaybeRenderTo(url, callback string) string {
	return renderTo(callback, url)
}

func (c *client) OpenidBindQuery(param url.Values) (res *openidRes, err error) {
	return openidBindQuery(param, c.req)
}

func (c *client) OpenidBindMap(param map[string]string) (res *openidRes, err error) {
	return openidBindMap(param, c.req)
}

func (c *client) Api() Api {
	return openApi(c.appKey, c.req)
}
