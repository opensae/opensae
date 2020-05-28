package model

// 由一个docker image启动Scale个容器
type Service struct {
	Image string
	Scale uint32
	Envs  map[string]string
}

