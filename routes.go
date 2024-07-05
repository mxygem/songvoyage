package main

import (
	"fmt"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func routes() *router.Router {
	r := router.New()

	r.GET("/healthcheck", healthcheck)
	r.GET("/request-echo", requestHandler)

	// v1 := r.Group("/v1")
	// setlist
	// slV1 := v1.Group("/setlist")
	// slV1.GET("")

	return r
}

// createSetlist handles requests to create a persisted setlist. A name must be provided
// otherwise the request will be handled as a no-op.
func createSetlist(ctx *fasthttp.RequestCtx) {}

// getSetlist handles requests to retrieve a setlist. If a name is provided, a persisted
// setlist will be returned if found. If no name is provided, it will return the current
// temporary setlist if it contains any songs. If neither is true, it will return a
// message indicating such.
func getSetlist(ctx *fasthttp.RequestCtx) {}

// updateSetlist handles requests to update a setlist. Currently, the only field that can
// be updated is a setlist's name. This should help in situations where a setlist was
// created with an incorrect name or the requester simply wants to change it. Both the
// existing and desired names must be provided.
func updateSetlist(ctx *fasthttp.RequestCtx) {}

// deleteSetlist handles requests to remove a setlist. The name provided must be an exact
// match in order for the delete to be processed successfully.
func deleteSetlist(ctx *fasthttp.RequestCtx) {}

// clearSetlist handles requests to clear all songs from a particular setlist. If no name
// is provided, then the current temporary setlist will be cleared.
func clearSetlist(ctx *fasthttp.RequestCtx) {}

// saveSetlist handles requests to save the current temporary setlist as a persisted
// setlist with the provided name. A name is required for this request to be processed
// successfully.
func saveSetlist(ctx *fasthttp.RequestCtx) {}

// addSong handles requests to append a song to a setlist. If no setlist name is provided
// the song will be added to the temporary setlist.
func addSong(ctx *fasthttp.RequestCtx) {}

// removeSong handles requests to remove a song from a setlist. If no setlist name is
// provided the song will be removed from the temporary setlist if it exists on the
// setlist.
func removeSong(ctx *fasthttp.RequestCtx) {}

// findSong handles requests to look up a song. It looks up a particular song in a user's
// stored song list and if not found, searches chorus to see if its chart is available to
// download. It returns an applicable message based on its findings.
// It currently is not in use and its implementation depends on a couple of factors:
//   - Whether or not chorus allows external services to search its catalog
//   - Whether or not a feature allowing users to upload their list of songs is
//     implemented.
// func findSong(ctx *fasthttp.RequestCtx) {}

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
