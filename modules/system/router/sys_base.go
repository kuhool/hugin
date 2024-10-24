package router

import (
	"github.com/gin-gonic/gin"
	"yxgin/modules/instance"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := instance.ApiGroupApp.SystemApiGroup.BaseApi
	{

		baseRouter.POST("captcha", baseApi.Captcha)
		//baseRouter.POST("upload", baseApi.Upload)

	}
	return baseRouter
}
