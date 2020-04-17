package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/saltbo/opensae/engine/api/binding"
	"github.com/saltbo/opensae/model"
)

type FuncResource struct {
	dbc *mongo.Client
}

func NewFuncResource(dbc *mongo.Client) *FuncResource {
	return &FuncResource{
		dbc: dbc,
	}
}

func (r *FuncResource) Register(router *gin.RouterGroup) {
	router.GET("/fcs", r.list)
	router.POST("/fcs", r.create)
	router.DELETE("/fcs/:id", r.remove)
}

func (r *FuncResource) list(c *gin.Context) {
	ms := make([]model.Func, 0)
	//r.engine.Database.Find(&ms)
	c.JSON(http.StatusOK, ms)
}

func (r *FuncResource) create(c *gin.Context) {
	p := new(binding.Func)
	if err := c.ShouldBind(p); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//mp := &model.Func{
	//	Service: p.Service,
	//	Name:    p.Name,
	//	Intro:   p.Intro,
	//}
	//if err := r.engine.Database.Create(mp).Error; err != nil {
	//	c.AbortWithError(http.StatusBadRequest, err)
	//	return
	//}

	c.Status(http.StatusOK)
}

func (r *FuncResource) remove(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	c.AbortWithError(http.StatusBadRequest, err)
	//	return
	//}

	//mp := &model.Func{Model: gorm.Model{ID: uint(id)}}
	//if err := r.engine.Database.Delete(mp).Error; err != nil {
	//	c.AbortWithError(http.StatusBadRequest, err)
	//	return
	//}

	c.Status(http.StatusOK)
}
