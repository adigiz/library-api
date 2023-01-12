package testutils

import (
	"encoding/json"
	"library-api/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ServeReq(opts *server.RouterOpts, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.NewRouter(opts)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}
