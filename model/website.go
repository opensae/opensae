package model

// 基于Caddy实现一个反向代理，根据host和路由将流量调度给Service
type Website struct {
	Name  string
	Intro string
	VHost []string
}

