package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AnnaVyvert/safe-concept-server/cmd/server/routes"
	"github.com/AnnaVyvert/safe-concept-server/cmd/server/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoutes(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	_ = require

	_ = utils.Load()

	mux := http.NewServeMux()
	routes.DefineRoutes(mux)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	type Request struct {
		method string
		url    string
		header http.Header
		body   io.Reader
	}

	type Expected struct {
		code    int
		headers http.Header
		body    []byte
	}

	file_content := "mock_file_content"
	
	testCases := []struct {
		desc     string
		request  Request
		expected Expected
	}{
		{
			desc:     "root",
			request:  Request{method: http.MethodGet, url: "/"},
			expected: Expected{code: http.StatusOK, body: []byte("root")},
		},
		{
			desc: "get_file",
			request: Request{method: http.MethodGet, url: "/get_file", header: http.Header{
				http.CanonicalHeaderKey("token"): []string{"todo_token"},
			}},
			expected: Expected{code: http.StatusOK, body: []byte(file_content)},
		},
		{
			desc: "put_file",
			request: Request{method: http.MethodPut, url: "/put_file", header: http.Header{
				http.CanonicalHeaderKey("token"): []string{"todo_token"},
			}},
			expected: Expected{code: http.StatusOK, body: nil},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			request := tC.request
			expected := tC.expected

			t.Log("REQUEST URL:", testServer.URL+request.url)

			req, err := http.NewRequest(request.method, testServer.URL+request.url, request.body)
			assert.NoError(err)

			res, err := http.DefaultClient.Do(req)
			assert.NoError(err)
			defer res.Body.Close()

			if expected.body != nil {
				responseBody, err := io.ReadAll(res.Body)
				assert.NoError(err, "can not read body", res.StatusCode)
				assert.Equal(expected.body, responseBody)
			}
			assert.Equal(http.StatusOK, res.StatusCode)
		})
	}
}
