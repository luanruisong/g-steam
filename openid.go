package steam

import (
	"errors"
	"net/url"
	"reflect"
	"strings"
)

const (
	loginUrl = "http://steamcommunity.com/openid/login"
	nsUrl    = "http://specs.openid.net/auth/2.0"
	nsSelect = nsUrl + "/identifier_select"
)

var t reflect.Type

func init() {
	t = reflect.TypeOf(openidRes{})
}

type (
	openidRes struct {
		Ns            string `form:"openid.ns" json:"openid.ns" openid:"openid.ns"`
		Mode          string `form:"openid.mode" json:"openid.mode" openid:"openid.mode"`
		Error         string `form:"openid.error" json:"openid.error" openid:"openid.error"`
		OpEndpoint    string `json:"openid.op_endpoint" form:"openid.op_endpoint" openid:"openid.op_endpoint"`
		Identity      string `json:"openid.identity" form:"openid.identity" openid:"openid.identity"`
		ClaimedId     string `json:"openid.claimed_id" form:"openid.claimed_id" openid:"openid.claimed_id"`
		ReturnTo      string `json:"openid.return_to" form:"openid.return_to" openid:"openid.return_to"`
		ResponseNonce string `json:"openid.response_nonce" form:"openid.response_nonce" openid:"openid.response_nonce"`
		AssocHandle   string `json:"openid.assoc_handle" form:"openid.assoc_handle" openid:"openid.assoc_handle"`
		Signed        string `json:"openid.signed" form:"openid.signed" openid:"openid.signed"`
		Sig           string `json:"openid.sig" form:"openid.sig" openid:"openid.sig"`
	}
)

func renderTo(callback string) string {
	u, _ := url.Parse(loginUrl)
	query := u.Query()
	query.Add("openid.ns", nsUrl)
	query.Add("openid.mode", "checkid_setup")
	query.Add("openid.return_to", callback)
	query.Add("openid.realm", callback)
	query.Add("openid.identity", nsSelect)
	query.Add("openid.claimed_id", nsSelect)
	u.RawQuery = query.Encode()
	return u.String()
}

func (r *openidRes) scan(query url.Values) {
	v := reflect.ValueOf(r).Elem()
	for i := 0; i < t.NumField(); i++ {
		currTag := t.Field(i).Tag
		tName := currTag.Get("openid")
		if len(tName) > 0 {
			va := query.Get(tName)
			if len(va) > 0 {
				v.Field(i).SetString(va)
			}
		}
	}
}

func (r *openidRes) scanMap(m map[string]string) {
	v := reflect.ValueOf(r).Elem()
	for i := 0; i < t.NumField(); i++ {
		currTag := t.Field(i).Tag
		tName := currTag.Get("openid")
		if len(tName) > 0 {
			va := m[tName]
			if len(va) > 0 {
				v.Field(i).SetString(va)
			}
		}
	}
}

func bindQuery(query url.Values) *openidRes {
	r := new(openidRes)
	r.scan(query)
	return r
}

func bindMap(m map[string]string) *openidRes {
	r := new(openidRes)
	r.scanMap(m)
	return r
}

func (r *openidRes) validateSteamSign(req Req) error {
	if len(r.Error) > 0 {
		return errors.New(r.Error)
	}
	u, _ := url.Parse(loginUrl)
	v := reflect.ValueOf(r).Elem()
	const mode = "openid.mode"
	query := u.Query()
	for i := 0; i < t.NumField(); i++ {
		currTag := t.Field(i).Tag
		tName := currTag.Get("openid")
		if tName != mode {
			query.Add(tName, v.Field(i).String())
		}
	}
	query.Add(mode, "check_authentication")
	u.RawQuery = query.Encode()

	res, err := req.Post(u.String(), nil, nil)
	if err != nil {
		return err
	}
	m := map[string]string{}
	for _, v := range strings.Split(res, "\n") {
		if x := strings.Index(v, ":"); x > 0 {
			key := strings.TrimSpace(v[:x])
			value := strings.TrimSpace(v[x+1:])
			m[key] = value
		}
	}
	if t, ok := m["is_valid"]; !ok || t != "true" {
		return errors.New("validate failed")
	}
	return nil
}

func (r *openidRes) GetSteamId() string {
	//http://steamcommunity.com/openid/id/<steamid>
	idx := strings.LastIndex(r.ClaimedId, "/") + 1
	return r.ClaimedId[idx:]
}

func openidBindQuery(param url.Values, req Req) (res *openidRes, err error) {
	//绑定返回参数
	if openidRes := bindQuery(param); openidRes != nil {
		if err = openidRes.validateSteamSign(req); err == nil {
			res = openidRes
		}
	}
	return
}

func openidBindMap(param map[string]string, req Req) (res *openidRes, err error) {
	//绑定返回参数
	if openidRes := bindMap(param); openidRes != nil {
		if err = openidRes.validateSteamSign(req); err == nil {
			res = openidRes
		}
	}
	return
}
