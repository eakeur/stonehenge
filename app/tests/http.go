package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"stonehenge/app/core/types/id"
	accessManager "stonehenge/app/gateway/access"
	"stonehenge/app/gateway/api/middlewares"
	"stonehenge/app/gateway/api/rest"
)

const JWTTokenMock = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
const testTokenKey = "test"
const testTokenExpirationTime = 100000000000

var (
	responseBuilder = rest.ResponseBuilder{
		Logger: zerolog.New(os.Stdout),
	}
	access = accessManager.NewManager(testTokenExpirationTime, []byte(testTokenKey))
)

type Route struct {
	Method       string
	Pattern      string
	Handler      rest.Handler
	RequiresAuth bool
}

func (r Route) ServeHTTP(request *http.Request) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	middleware := middlewares.NewMiddleware(access, responseBuilder)
	router.Use(middleware.CORS)
	router.Use(middleware.RequestTracer)
	if r.RequiresAuth {
		router.Use(middleware.Authorization)
	}
	router.Method(r.Method, r.Pattern, r.Handler)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, request)
	return rec
}

func GetResponseBuilder() rest.ResponseBuilder {
	return responseBuilder
}

func EncodeURLParameters(params map[string]string) string {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}
	return "?" + q.Encode()
}

func CreateRequestWithBody(method, target string, body interface{}) *http.Request {
	var b *bytes.Reader
	if body != nil {
		parsed, _ := json.Marshal(body)
		b = bytes.NewReader(parsed)
	}
	return httptest.NewRequest(method, target, b)
}

func CreateRequestWithParams(method, target string, params map[string]string) *http.Request {
	route := target
	if params != nil {
		encoded := EncodeURLParameters(params)
		route += encoded
	}
	return httptest.NewRequest(method, route, nil)
}

func AuthenticateRequest(r *http.Request, accountID id.External) *http.Request {
	acc, _ := access.Create(accountID, "")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", acc.Token))
	return r
}
