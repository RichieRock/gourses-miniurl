package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/RichieRock/gourses-miniurl/api"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_AddUrl(t *testing.T) {
	const (
		payload = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE"}`
		expectedBody = `{"url": "https://github.com/gourses/miniurl/blob/main/LICENSE", "hash": "testvalue"}`
		expectedStatusCode = http.StatusOK
	)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(payload))
	rr := httptest.NewRecorder()

	r := httprouter.New()
	h := &strHandler{str: "testvalue"}
	api.Bind(r, h)
	r.ServeHTTP(rr, req)

	// check status code
	assert.Equal(t, expectedStatusCode, rr.Result().StatusCode)
	// read the body
	body, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	assert.JSONEq(t, expectedBody, string (body))
}

type strHandler struct {
	str string
}

func (h *strHandler) AddUrl(url string) (hash string, err error) {
	return h.str, nil
}