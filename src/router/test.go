package router

import (
	_const "awesomeProject/src/const"
	"awesomeProject/src/controller"

	"github.com/gin-gonic/gin"
)

func LoadTest(e *gin.Engine) {
	e.GET(_const.RouterTest, controller.Test())
	e.POST(_const.RouterCreateUser, controller.CreateUser())
	e.GET(_const.RouterGetUser, controller.GetUser())
}
