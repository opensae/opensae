package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Resource interface {
	Register(router *gin.RouterGroup)
}

func RouteRegister(r *gin.RouterGroup, dbc *mongo.Client) {
	resources := []Resource{
		NewFuncResource(dbc),
	}

	for _, resource := range resources {
		resource.Register(r)
	}
}
