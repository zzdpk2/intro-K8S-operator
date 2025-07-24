package router

import "kubeproj.com/router/example"

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
}

var RouterGroupApp = new(RouterGroup)
