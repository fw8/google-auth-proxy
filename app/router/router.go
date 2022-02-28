package Router

import (
	"fmt"
	"net/http/httputil"

	s "github.com/fw8/google-auth-proxy/app/strategy"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(rp *httputil.ReverseProxy) *httprouter.Router {
	r := httprouter.New()
	r.GET("/*path", s.AddAuthHeader(rp))
	r.POST("/*path", s.AddAuthHeader(rp))
	r.DELETE("/*path", s.AddAuthHeader(rp))
	r.PUT("/*path", s.AddAuthHeader(rp))
	r.OPTIONS("/*paths", s.FwdOptionsReq(rp))
	fmt.Println("Routes Registered")
	return r
}
