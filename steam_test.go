package steam

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetOpenId(t *testing.T) {
	client := NewClient("key123")

	render := client.RenderTo("http://127.0.0.1:9099/")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		res, err := client.OpenidBindQuery(request.URL.Query())
		fmt.Println(res.GetSteamId(), err)
	})
	fmt.Println(render)
	_ = http.ListenAndServe(":9099", nil)
}
