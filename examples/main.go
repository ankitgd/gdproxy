package main

import (
	"github.com/ankitgd/gdproxy"
	"log"
	"flag"
	//"fmt"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()

	proxy := gdproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().DoFunc(func(r *http.Request,ctx *gdproxy.ProxyCtx) (*http.Request, *http.Response){
		r.Header.Set("X-GoProxy","1")
		return r, nil
	})

	proxy.OnResponse().DoFunc(func(r *http.Response, ctx *gdproxy.ProxyCtx)*http.Response{
		println(ctx.Req.Host,"->",r.Header.Get("Content-Type"))
		return r
	})

	log.Fatal(http.ListenAndServe(*addr, proxy))
}
