package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

const (
	lh = "http://localhost"
)

func TestGetSetlist(t *testing.T) {
	testCases := []struct {
		name               string
		params             string
		db                 func(t *testing.T) *Mockdber
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name:   "successfully returns setlist by name when found",
			params: "?name=My%20Awesome%20Playlist",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, "My Awesome Playlist").
					Return(&Setlist{
						Name: "My Awesome Playlist",
						Songs: []*Song{
							{
								Artist: "Dragonforce",
								Name:   "Through the Fire and Flames",
							},
						},
					}, nil)

				return db
			},
			expectedBody: `{` +
				`"name":"My Awesome Playlist",` +
				`"songs":[{"artist":"Dragonforce","name":"Through the Fire and Flames"}]` +
				`}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:   "returns error text when calling db errors",
			params: "?name=My%20Awesome%20Playlist",
			db: func(t *testing.T) *Mockdber {
				db := NewMockdber(t)

				db.On("find", mock.Anything, "My Awesome Playlist").
					Return(nil, fmt.Errorf("something broke"))

				return db
			},
			expectedBody:       `{"error":"failed to get setlist"}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := newServer()
			s.db = tc.db(t)
			client := newTestServer(t, s.getSetlist)

			resp, err := client.Get(lh + tc.params)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			if tc.expectedBody != "" {
				require.NotNil(t, resp.Body)
				respBody, _ := io.ReadAll(resp.Body)
				assert.Equal(t, tc.expectedBody, string(respBody))
			}
		})
	}
}

// newTestServer configures an in memory listener (server) with the provided handler and
// returns a client to use to make requests against the server with.
func newTestServer(t *testing.T, h fasthttp.RequestHandler) *http.Client {
	t.Helper()

	s := &fasthttp.Server{
		Handler: h,
	}

	ln := fasthttputil.NewInmemoryListener()

	go func() {
		s.Serve(ln) // nolint: errcheck
	}()

	return &http.Client{Transport: &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}}
}
