package cors

import (
	"testing"

	"github.com/valyala/fasthttp"
)

func newRequestCtx(method, uri string) *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.SetMethod(method)
	ctx.Request.SetRequestURI(uri)
	return ctx
}

func BenchmarkWithout(b *testing.B) {
	ctx := newRequestCtx("GET", "http://example.com/foo")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testHandler(ctx)
	}
}

func BenchmarkDefault(b *testing.B) {
	ctx := newRequestCtx("GET", "http://example.com/foo")
	ctx.Request.Header.Add("Origin", "somedomain.com")
	handler := Default().Handler(testHandler)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}

func BenchmarkAllowedOrigin(b *testing.B) {
	ctx := newRequestCtx("GET", "http://example.com/foo")
	ctx.Request.Header.Add("Origin", "somedomain.com")
	c := New(Options{
		AllowedOrigins: []string{"somedomain.com"},
	})
	handler := c.Handler(testHandler)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}

func BenchmarkPreflight(b *testing.B) {
	ctx := newRequestCtx("OPTIONS", "http://example.com/foo")
	ctx.Request.Header.Add("Access-Control-Request-Method", "GET")
	handler := Default().Handler(testHandler)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}

func BenchmarkPreflightHeader(b *testing.B) {
	ctx := newRequestCtx("OPTIONS", "http://example.com/foo")
	ctx.Request.Header.Add("Access-Control-Request-Method", "GET")
	ctx.Request.Header.Add("Access-Control-Request-Headers", "Accept")
	handler := Default().Handler(testHandler)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler(ctx)
	}
}
