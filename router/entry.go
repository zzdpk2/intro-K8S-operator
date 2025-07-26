package router

import (
	"kubeproj.com/router/example"
	"kubeproj.com/router/kubernetes"
)

type RouterGroup struct {
	ExampleRouterGroup    example.ExampleRouter
	KubernetesRouterGroup kubernetes.KubernetesRouter
}

var RouterGroupApp = new(RouterGroup)
