package main

import (
	"net/http"

	"github.com/nats-io/nats"
	"gopkg.in/sohlich/nats-proxy.v1"
)

func main() {
	proxyConn, _ := nats.Connect(nats.DefaultURL)
	proxy, _ := natsproxy.NewNatsProxy(proxyConn)
	defer proxyConn.Close()
	http.ListenAndServe(":8080", proxy)

	// proxy.AddHook(".*", func(r *natsproxy.Response) {
	// 	// Exchange the jwt token for
	// 	// reference token to hide user information
	// 	jwt := r.GetHeader().Get("X-Auth")
	// 	refToken := auth.GetTokenFor(jwt)
	// 	r.GetHeader().Set("X-Auth", refToken)
	// })
}

