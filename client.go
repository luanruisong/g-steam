package steam

import "net/url"

type (
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

func (c *client) OpenidBindQuery(param url.Values) (res *openidRes, err error) {
	return openidBindQuery(param, c.req)
}

func (c *client) OpenidBindMap(param map[string]string) (res *openidRes, err error) {
	return openidBindMap(param, c.req)
}
