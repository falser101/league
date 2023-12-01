package cluster

import "github.com/gin-gonic/gin"

type ClusterRouterImpl struct {
	clusterRestHandler ClusterRestHandler
}

func NewClusterRouterImpl(handler ClusterRestHandler) *ClusterRouterImpl {
	return &ClusterRouterImpl{
		clusterRestHandler: handler,
	}
}
func (impl *ClusterRouterImpl) AddRoutes(router *gin.RouterGroup) {
	clusterRouter := router.Group("/cluster")
	clusterRouter.GET("", impl.clusterRestHandler.FindAll)
	clusterRouter.POST("", impl.clusterRestHandler.Save)
	clusterRouter.DELETE("/:cluster", impl.clusterRestHandler.DeleteCluster)
	clusterRouter.PUT("/:cluster", impl.clusterRestHandler.Update)
}
