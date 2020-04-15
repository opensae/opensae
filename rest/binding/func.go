package binding

type Func struct {
	Service string `form:"service" binding:"required"`
	Name    string `form:"name" binding:"required"`
	Intro   string `form:"intro"`
}
