package model

import (
	"github.com/jinzhu/gorm"
)

type Func struct {
	gorm.Model

	Service   string
	Name      string
	Intro     string
	TimerSpec string
	Upstream  string
}
