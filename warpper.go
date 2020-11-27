package steam

import (
	"errors"
	"io/ioutil"
	"net/http"
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
	t = reflect.TypeOf(OpenidRes{})
}

type (
	renderParam struct {
		returnTo string
	}

	//http://127.0.0.1:9099/?
	//openid.ns=http://specs.openid.net/auth/2.0
	//openid.mode=id_res
	//openid.op_endpoint=https://steamcommunity.com/openid/login
	//openid.claimed_id=https://steamcommunity.com/openid/id/76561198421538055
	//openid.identity=https://steamcommunity.com/openid/id/76561198421538055
	//openid.return_to=http://127.0.0.1:9099/callback
	//openid.response_nonce=2020-11-27T08:35:56Z3eJ+NfRY4xZwPFPH1wT1E3gg/sk=
	//openid.assoc_handle=1234567890
	//openid.signed=signed,op_endpoint,claimed_id,identity,return_to,response_nonce,assoc_handle
	//openid.sig=QZmrDUWx8cP/QA2YIJTok7/mUfM=

	OpenidRes struct {
		raw           url.Values
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

func (rp *renderParam) GetFullUrl() string {
	u, _ := url.Parse(loginUrl)
	query := u.Query()
	query.Add("openid.ns", nsUrl)
	query.Add("openid.mode", "checkid_setup")
	query.Add("openid.return_to", rp.returnTo)
	query.Add("openid.realm", rp.returnTo)
	query.Add("openid.identity", nsSelect)
	query.Add("openid.claimed_id", nsSelect)
	u.RawQuery = query.Encode()
	return u.String()
}

func (r *OpenidRes) scan() {
	v := reflect.ValueOf(r).Elem()
	for i := 0; i < t.NumField(); i++ {
		currTag := t.Field(i).Tag
		tName := currTag.Get("openid")
		if len(tName) > 0 {
			va := r.raw.Get(tName)
			if len(va) > 0 {
				v.Field(i).SetString(va)
			}
		}
	}
}

func (r *OpenidRes) Bind(req *http.Request) error {
	query := req.URL.Query()
	r.raw = query
	r.scan()
	if err := r.validateSteamSign(); err != nil {
		return err
	}
	return nil
}

func NewRender(returnTo string) *renderParam {
	return &renderParam{
		returnTo: returnTo,
	}
}

func RenderTo(callback string) string {
	return NewRender(callback).GetFullUrl()
}

func (r *OpenidRes) validateSteamSign() error {
	u, _ := url.Parse(loginUrl)
	const mode = "openid.mode"
	query := u.Query()
	for i := range r.raw {
		if i != mode {
			query.Add(i, r.raw.Get(i))
		} else {
			query.Add(mode, "check_authentication")
		}
	}
	u.RawQuery = query.Encode()
	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	m := map[string]string{}
	for _, v := range strings.Split(string(b), "\n") {
		if x := strings.Split(v, ":"); len(v) == 2 {
			m[x[0]] = x[1]
		}
	}
	if t, ok := m["is_valid"]; !ok || t != "true" {
		return errors.New("validate failed")
	}
	return nil
}

func (r *OpenidRes) GetSteamId() string {
	//http://steamcommunity.com/openid/id/<steamid>
	idx := strings.LastIndex(r.ClaimedId, "/") + 1
	return r.ClaimedId[idx:]
}
