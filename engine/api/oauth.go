package api

import (
	"github.com/gin-gonic/gin"
)

// 负责登录，拿到身份标识
type Oauth struct {
}

// 如何鉴权
// 登录=》获取身份
// 该身份决定了是否可读写哪些数据

func (r *Oauth) Register(router *gin.RouterGroup) {
	// 登录，GitHub，Google，微信
	// 用户注册
	// 密码登录
	// 密码找回
	// 图片验证码
	// 邮箱验证码

	router.POST("/oauth/signin")
	router.POST("/oauth/signout")
	router.POST("/oauth/")
}
