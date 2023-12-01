package api

import (
	"go-league/api/cluster"

	"github.com/gin-gonic/gin"
)

var clusterRouter *cluster.ClusterRouterImpl = cluster.NewClusterRouterImpl(cluster.NewClusterRestHandlerImpl())

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		clusterRouter.AddRoutes(v1)
	}
}
