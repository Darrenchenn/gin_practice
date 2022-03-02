package controller

import (
	_const "awesomeProject/src/const"
	"awesomeProject/src/model"
	"awesomeProject/src/util"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var log = logrus.WithFields(logrus.Fields{"pkg": "handler"})

func Test() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req _const.TestRequest
		if err := util.ParseParameters(context, &req); err != nil {
			if err := context.Error(err); err != nil {
				log.Errorf("failed to add error:[%s]", err.Error())
				return
			}
			return
		}

		util.SetOutput(context, req)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req _const.TestRequest
		if err := util.ParseParameters(context, &req); err != nil {
			if err := context.Error(err); err != nil {
				log.Errorf("failed to add error:[%s]", err.Error())
				return
			}
			return
		}
		log.Debugf("request parameters:[%v]", req)
		resp := model.CreateUser(&req)

		util.SetOutput(context, resp)
	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req _const.TestRequest
		if err := util.ParseParameters(context, &req); err != nil {
			if err := context.Error(err); err != nil {
				log.Errorf("failed to add error:[%s]", err.Error())
				return
			}
			return
		}
		log.Debugf("request parameters:[%v]", req)
		resp := model.GetUser(&req)

		util.SetOutput(context, resp)
	}
}
