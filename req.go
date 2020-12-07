package steam

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Req interface {
		Get(urlStr string) (string, error)
		Post(urlStr, body string) (string, error)
	}
	defReq struct {
		c *http.Client
	}
)

func (d *defReq) Get(urlStr string) (body string, err error) {

	var (
		rsp *http.Response
		b   []byte
	)
	rsp, err = http.Get(urlStr)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	if b, err = ioutil.ReadAll(rsp.Body); err != nil {
		return
	}
	body = string(b)
	return
}

func (d *defReq) Post(urlStr, body string) (resBody string, err error) {

	var (
		rsp *http.Response
		b   []byte
	)
	rsp, err = http.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	if b, err = ioutil.ReadAll(rsp.Body); err != nil {
		return
	}
	resBody = string(b)
	return
}

func NewDefReq() *defReq {
	return &defReq{
		c: http.DefaultClient,
	}
}
