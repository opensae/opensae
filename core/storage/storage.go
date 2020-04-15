package storage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Operation(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		s.find()
	case http.MethodPost:
		s.create()
	case http.MethodPut:
		s.update()
	case http.MethodDelete:
		s.delete()

	}
}

func (s *Storage) find() {

}

func (s *Storage) create() {

}

func (s *Storage) update() {

}

func (s *Storage) delete() {

}
