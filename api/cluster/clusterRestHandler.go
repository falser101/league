package cluster

import (
	"github.com/gin-gonic/gin"
)

type ClusterRestHandler interface {
	Save(c *gin.Context)
	SaveClusters(c *gin.Context)
	ValidateKubeconfig(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	FindNoteByClusterId(c *gin.Context)
	Update(c *gin.Context)
	UpdateClusterDescription(c *gin.Context)
	UpdateClusterNote(c *gin.Context)
	FindAllForAutoComplete(c *gin.Context)
	DeleteCluster(c *gin.Context)
	GetClusterNamespaces(c *gin.Context)
	GetAllClusterNamespaces(c *gin.Context)
	FindAllForClusterPermission(c *gin.Context)
}

type ClusterRestHandlerImpl struct {
}

func NewClusterRestHandlerImpl() ClusterRestHandler {
	return &ClusterRestHandlerImpl{}
}

// DeleteCluster implements ClusterRestHandler.
func (ClusterRestHandlerImpl) DeleteCluster(c *gin.Context) {
	panic("unimplemented")
}

// FindAll implements ClusterRestHandler.
func (ClusterRestHandlerImpl) FindAll(c *gin.Context) {
	panic("unimplemented")
}

// FindAllForAutoComplete implements ClusterRestHandler.
func (ClusterRestHandlerImpl) FindAllForAutoComplete(c *gin.Context) {
	panic("unimplemented")
}

// FindAllForClusterPermission implements ClusterRestHandler.
func (ClusterRestHandlerImpl) FindAllForClusterPermission(c *gin.Context) {
	panic("unimplemented")
}

// FindById implements ClusterRestHandler.
func (ClusterRestHandlerImpl) FindById(c *gin.Context) {
	panic("unimplemented")
}

// FindNoteByClusterId implements ClusterRestHandler.
func (ClusterRestHandlerImpl) FindNoteByClusterId(c *gin.Context) {
	panic("unimplemented")
}

// GetAllClusterNamespaces implements ClusterRestHandler.
func (ClusterRestHandlerImpl) GetAllClusterNamespaces(c *gin.Context) {
	panic("unimplemented")
}

// GetClusterNamespaces implements ClusterRestHandler.
func (ClusterRestHandlerImpl) GetClusterNamespaces(c *gin.Context) {
	panic("unimplemented")
}

// Save implements ClusterRestHandler.
func (ClusterRestHandlerImpl) Save(c *gin.Context) {
	panic("unimplemented")
}

// SaveClusters implements ClusterRestHandler.
func (ClusterRestHandlerImpl) SaveClusters(c *gin.Context) {
	panic("unimplemented")
}

// Update implements ClusterRestHandler.
func (ClusterRestHandlerImpl) Update(c *gin.Context) {
	panic("unimplemented")
}

// UpdateClusterDescription implements ClusterRestHandler.
func (ClusterRestHandlerImpl) UpdateClusterDescription(c *gin.Context) {
	panic("unimplemented")
}

// UpdateClusterNote implements ClusterRestHandler.
func (ClusterRestHandlerImpl) UpdateClusterNote(c *gin.Context) {
	panic("unimplemented")
}

// ValidateKubeconfig implements ClusterRestHandler.
func (ClusterRestHandlerImpl) ValidateKubeconfig(c *gin.Context) {
	panic("unimplemented")
}
