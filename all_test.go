package openweathermap

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/otiai10/marmoset"

	. "github.com/otiai10/mint"
)

func TestNew(t *testing.T) {

	mock := server()
	c := New("valid")
	c.BaseURL = mock.URL
	Expect(t, c).TypeOf("*openweathermap.Client")

	res, err := c.ByCityName("tokyo", nil)
	Expect(t, err).ToBe(nil)
	Expect(t, res.Code).ToBe("200")
}

func server() *httptest.Server {
	r := marmoset.NewRouter()
	r.GET("/data/2.5/forecast", func(rw http.ResponseWriter, req *http.Request) {
		if req.FormValue("apikey") != "valid" {
			f, err := os.Open("test/response.401.json")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			io.Copy(rw, f)
		}
		f, err := os.Open("test/response.200.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		io.Copy(rw, f)
	})
	s := httptest.NewServer(r)
	return s
}
