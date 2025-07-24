package api

import "kubeproj.com/api/example"

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
