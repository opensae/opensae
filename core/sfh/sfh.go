package sfh

import (
	"github.com/gin-gonic/gin"
)

// sfh: Static file hosting

type SFH struct {
}

func New() *SFH {
	return &SFH{}
}

func (s *SFH) Render(c *gin.Context) {

}