package main

import (
	"github.com/ankitgd/gdproxy"
	"log"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":2357", "proxy listen address")
	flag.Parse()

	proxy := gdproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest(gdproxy.MethodIs("GET")).DoFunc(func(req *http.Request, ctx *gdproxy.ProxyCtx) (*http.Request, *http.Response){
		fmt.Println(*req)
		return req, nil
	})

	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *gdproxy.ProxyCtx)*http.Response{
		fmt.Println(resp)
		return resp
	})

	log.Fatal(http.ListenAndServe(*addr, proxy))
}
