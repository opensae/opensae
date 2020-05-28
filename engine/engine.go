package engine

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/saltbo/opensae/engine/api"
	"github.com/saltbo/opensae/engine/db"
	"github.com/saltbo/opensae/engine/sae"
	"github.com/saltbo/opensae/engine/sfh"
)

type Engine struct {
	Router   *gin.Engine
	Database *mongo.Client
	Cron     *cron.Cron
}

func New(dsn string) (*Engine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		Router:   gin.Default(),
		Database: dbClient,
	}

	return engine, nil
}

func (e *Engine) Boot() error {
	// 云引擎
	saeRouter := e.Router.Group("/sae")
	sae.RouteRegister(saeRouter, e.Database)

	// 数据库
	storageRouter := e.Router.Group("/db")
	db.RouteRegister(storageRouter, e.Database)

	// 系统API
	apiRouter := e.Router.Group("/api")
	api.RouteRegister(apiRouter, e.Database)

	// 静态文件托管
	static := sfh.New()
	e.Router.NoRoute(static.Render)

	return e.Router.Run(":8081")
}
