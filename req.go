package steam

import (
	"io/ioutil"
	"net/http"
)

type (
	Req interface {
		Get(urlStr string) (string, error)
		Post(urlStr string) (string, error)
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

func (d *defReq) Post(urlStr string) (body string, err error) {

	var (
		rsp *http.Response
		b   []byte
	)
	rsp, err = http.Post(urlStr, "", nil)
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

func NewDefReq() *defReq {
	return &defReq{
		c: http.DefaultClient,
	}
}
