package api

import (
	"kubeproj.com/api/example"
	"kubeproj.com/api/kubernetes"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     kubernetes.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
