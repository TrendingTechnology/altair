package mock

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	entity "github.com/codefluence-x/altair/entity"
)

type ErrorResponse struct {
	Errors []entity.ErrorObject `json:"errors"`
}

type MockErrorIoReader struct {
}

func (m MockErrorIoReader) Read(x []byte) (int, error) {
	return 0, errors.New("read error")
}

func PerformRequest(r http.Handler, method, path string, body io.Reader, reqModifiers ...func(req *http.Request)) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	for _, f := range reqModifiers {
		f(req)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
