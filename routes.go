package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func routes(s *server) *router.Router {
	r := router.New()

	r.GET("/healthcheck", healthcheck)
	r.GET("/request-echo", requestHandler)

	v1 := r.Group("/v1")

	slV1 := v1.Group("/setlist")
	// due to initial unknown implementation direction and request capabilities of
	// existing chat bots, all routes will be accessed via GETs with all data added as
	// query parameters.
	slV1.GET("/", s.getSetlist)
	slV1.GET("/create", s.createSetlist)
	slV1.GET("/clear", s.clearSetlist)
	slV1.GET("/save", s.saveSetlist)
	slV1.GET("/delete", s.deleteSetlist)

	upV1 := slV1.Group("/update")
	upV1.GET("/", s.updateSetlist)
	upV1.GET("/add_song", s.addSong)
	upV1.GET("/remove_song", s.removeSong)

	return r
}

// getSetlist handles requests to retrieve a setlist. If a name is provided, a persisted
// setlist will be returned if found. If no name is provided, it will return the current
// temporary setlist if it contains any songs. If neither is true, it will return a
// message indicating such.
func (s *server) getSetlist(rctx *fasthttp.RequestCtx) {
	action := "get setlist"
	name := rctx.QueryArgs().Peek("name")

	ctx, cancel := context.WithTimeout(rctx, 5*time.Second)
	defer cancel()

	// TODO: Return non-error response when setlist is not found
	sl, err := setlist(ctx, s.db, name)
	if err != nil {
		log.Printf("%s - getting setlist: %s", action, err)
		// TODO: Create error types for automatic formatting
		rctx.Error(`{"error":"failed to get setlist"}`, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(sl)
	if err != nil {
		log.Printf("%s - marshalling setlist data: %s", action, err)
		rctx.Error(fmt.Sprintf("preparing response data: %s", err), http.StatusInternalServerError)
		return
	}

	// TODO: Determine a good way of formatting data for output to chat.
	if _, err := rctx.Write(b); err != nil {
		log.Printf("%s - writing response: %s", action, err)
		rctx.Error(fmt.Sprintf("preparing response data: %s", err), http.StatusInternalServerError)
		return
	}
}

// createSetlist handles requests to create a persisted setlist. A name must be provided
// otherwise the request will be handled as a no-op.
func (s *server) createSetlist(ctx *fasthttp.RequestCtx) {}

// deleteSetlist handles requests to remove a setlist. The name provided must be an exact
// match in order for the delete to be processed successfully.
func (s *server) deleteSetlist(ctx *fasthttp.RequestCtx) {}

// clearSetlist handles requests to clear all songs from a particular setlist. If no name
// is provided, then the current temporary setlist will be cleared.
func (s *server) clearSetlist(ctx *fasthttp.RequestCtx) {}

// saveSetlist handles requests to save the current temporary setlist as a persisted
// setlist with the provided name. A name is required for this request to be processed
// successfully.
func (s *server) saveSetlist(ctx *fasthttp.RequestCtx) {}

// updateSetlist handles requests to update a setlist. Currently, the only field that can
// be updated is a setlist's name. This should help in situations where a setlist was
// created with an incorrect name or the requester simply wants to change it. Both the
// existing and desired names must be provided.
func (s *server) updateSetlist(ctx *fasthttp.RequestCtx) {}

// addSong handles requests to append a song to a setlist. If no setlist name is provided
// the song will be added to the temporary setlist.
func (s *server) addSong(ctx *fasthttp.RequestCtx) {}

// removeSong handles requests to remove a song from a setlist. If no setlist name is
// provided the song will be removed from the temporary setlist if it exists on the
// setlist.
func (s *server) removeSong(ctx *fasthttp.RequestCtx) {}

// findSong handles requests to look up a song. It looks up a particular song in a user's
// stored song list and if not found, searches chorus to see if its chart is available to
// download. It returns an applicable message based on its findings.
// It currently is not in use and its implementation depends on a couple of factors:
//   - Whether or not chorus allows external services to search its catalog
//   - Whether or not a feature allowing users to upload their list of songs is
//     implemented.
// func findSong(ctx *fasthttp.RequestCtx) {}

// findSongs handles requests to look up potenally multiple songs. It returns a list of
// matched songs by partial name or artist name. Other parameters TBD. Also TBD is
// pagination support in case there are more results provided than can be reasonably
// returned.
// It is also currently not in use and also depends on the above factors in the comment
// for findSong
// func findSongs(ctx *fasthttp.RequestCtx) {}

// healthcheck handles requests to inquire whether the service is running or not. It
// currently returns no data, only an HTTP status code of 200 if successful.
func healthcheck(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(http.StatusOK)
}

// requestHandler is a temporary debug handler that echos back information to the caller
// about their request.
func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
