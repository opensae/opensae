package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/saltbo/opensae/core"
	"github.com/saltbo/opensae/model"
	"github.com/saltbo/opensae/rest/binding"
)

type FuncResource struct {
	engine *core.Engine
}

func NewFuncResource(engine *core.Engine) *FuncResource {
	return &FuncResource{
		engine: engine,
	}
}

func (r *FuncResource) Register(router *gin.RouterGroup) {
	router.GET("/fcs", r.list)
	router.POST("/fcs", r.create)
	router.DELETE("/fcs/:id", r.remove)
}

func (r *FuncResource) list(c *gin.Context) {
	ms := make([]model.Func, 0)
	r.engine.Database.Find(&ms)
	c.JSON(http.StatusOK, ms)
}

func (r *FuncResource) create(c *gin.Context) {
	p := new(binding.Func)
	if err := c.ShouldBind(p); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	mp := &model.Func{
		Service: p.Service,
		Name:    p.Name,
		Intro:   p.Intro,
	}
	if err := r.engine.Database.Create(mp).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *FuncResource) remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	mp := &model.Func{Model: gorm.Model{ID: uint(id)}}
	if err := r.engine.Database.Delete(mp).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}
