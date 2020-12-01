package steam

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetOpenId(t *testing.T) {
	render := NewRender("http://127.0.0.1:9099/")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		res := &OpenidRes{}
		_ = res.Bind(request)
		fmt.Println(res.GetSteamId())
	})
	fmt.Println(render.GetFullUrl())
	_ = http.ListenAndServe(":9099", nil)
}
