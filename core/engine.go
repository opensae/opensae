package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"

	"github.com/saltbo/opensae/core/func"
	"github.com/saltbo/opensae/core/sfh"
	"github.com/saltbo/opensae/core/storage"
)

type Engine struct {
	Router   *gin.Engine
	Database *gorm.DB
	Cron     *cron.Cron
}

func New(driver, dsn string) (*Engine, error) {
	database, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	engine := &Engine{
		Router:   gin.Default(),
		Database: database,
	}

	return engine, nil
}

func (e *Engine) Boot() error {
	// 存储
	s := storage.New()
	e.Router.Any("/storage/:doc", s.Operation)

	// 函数计算
	f := _func.New()
	e.Router.Any("/fc/:service/:name", f.HTTPTrigger)

	// 静态文件托管
	static := sfh.New()
	e.Router.NoRoute(static.Render)

	return e.Router.Run()
}
