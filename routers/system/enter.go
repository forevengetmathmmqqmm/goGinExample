package system

type RoutersTrc struct {
	CommentRouter
	UserRouter
	MusicianRouter
	ActivityRouter
	BannerRouter
	RoleRouter
	AccessRouter
}

var Routers *RoutersTrc

func init() {
	Routers = new(RoutersTrc)
}
