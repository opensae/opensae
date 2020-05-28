package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	dbc *mongo.Client
}

func RouteRegister(r *gin.RouterGroup, dbc *mongo.Client) {
	s := &Storage{
		dbc: dbc,
	}
	r.Any("/:doc", s.Operation)
}

func (s *Storage) Operation(c *gin.Context) {
	app := c.GetString("app")
	doc := c.Param("doc")

	// 检查是否有权限操作doc

	collection := s.dbc.Database("app:" + app).Collection(doc)
	switch c.Request.Method {
	case http.MethodGet:
		s.find(c, collection)
	case http.MethodPost:
		s.create(c, collection)
	case http.MethodPut:
		s.update(c, collection)
	case http.MethodDelete:
		s.delete(c, collection)
	}
}

func (s *Storage) find(c *gin.Context, coll *mongo.Collection) {
	//coll.Find()
}

func (s *Storage) create(c *gin.Context, coll *mongo.Collection) {
	//coll.InsertMany()
}

func (s *Storage) update(c *gin.Context, coll *mongo.Collection) {
	//coll.UpdateMany()
}

func (s *Storage) delete(c *gin.Context, coll *mongo.Collection) {
	//coll.DeleteMany()
}
