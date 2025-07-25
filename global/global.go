package global

import (
	"k8s.io/client-go/kubernetes"
	"kubeproj.com/config"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
