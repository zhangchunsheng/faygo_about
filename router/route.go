package router

import (
	"faygo_about/handler"
	"faygo_about/middleware"
	"github.com/henrylee2cn/faygo"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
	)
}
