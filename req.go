package steam

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	Req interface {
		Get(urlStr string, param, header map[string]string) (string, error)
		Post(urlStr string, param, header map[string]string) (string, error)
	}
	defReq struct {
		c *http.Client
	}
)

func (d *defReq) do(method, urlStr string, param, header map[string]string) (res string, err error) {
	var (
		reader *strings.Reader
		body   url.Values
		req    *http.Request
		rsp    *http.Response
		b      []byte
	)
	if len(param) > 0 {
		body = url.Values{}
		for i, v := range param {
			body.Add(i, v)
		}
		reader = strings.NewReader(body.Encode())
	}
	if req, err = http.NewRequest(method, urlStr, reader); err != nil {
		return
	}
	for i, v := range header {
		req.Header.Add(i, v)
	}
	if rsp, err = d.c.Do(req); err != nil {
		return
	}
	defer rsp.Body.Close()
	if b, err = ioutil.ReadAll(rsp.Body); err != nil {
		return
	}
	res = string(b)
	return
}
func (d *defReq) Get(urlStr string, param, header map[string]string) (string, error) {
	return d.do(http.MethodGet, urlStr, param, header)
}

func (d *defReq) Post(urlStr string, param, header map[string]string) (string, error) {
	return d.do(http.MethodPost, urlStr, param, header)
}

func NewDefReq() *defReq {
	return &defReq{
		c: http.DefaultClient,
	}
}
