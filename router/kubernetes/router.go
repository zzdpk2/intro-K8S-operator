package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubeproj.com/api"
)

type KubernetesRouter struct {
}

func (*KubernetesRouter) InitKubernetesSRouter(r *gin.Engine) {
	group := r.Group("/kubernetes")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.GET("/listPod", apiGroup.GetPodList)
}
