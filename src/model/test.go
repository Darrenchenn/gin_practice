package model

import (
	_const "awesomeProject/src/const"
	"awesomeProject/src/redis_client"
	"awesomeProject/src/util"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{"pkg": "model"})

func CreateUser(req *_const.TestRequest) *_const.TestResponse {
	resp := &_const.TestResponse{Status: util.SuccessStatus}
	userID, err := redis_client.Get(_const.RedisKeyUserID)
	if err == redis.Nil {
		current, err := redis_client.Incr(_const.RedisKeyUserID)
		if err != nil {
			log.Errorf("err : [%s]", err.Error())
			resp.Status = util.NewError(err)
			return resp
		}
		log.Infof("debug current:[%v]", current)
		userID = strconv.FormatInt(current, 10)
	} else if err != nil {
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	} else {
		current, err := redis_client.Incr(_const.RedisKeyUserID)
		if err != nil {
			log.Errorf("err : [%s]", err.Error())
			resp.Status = util.NewError(err)
			return resp
		}
		userID = strconv.FormatInt(current, 10)
	}

	key := fmt.Sprintf("%s:%s:%s", _const.RedisRootPrefix, req.Name, userID)
	val, err := util.Serialize(req)
	if err != nil {
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	}

	log.Debugf("debug set user key:[%s]", key)
	log.Debugf("debug set user val:[%s]", val)

	if err := redis_client.Set(key, val); err != nil {
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	}

	return resp
}

func GetUser(req *_const.TestRequest) *_const.TestResponse {
	resp := &_const.TestResponse{Status: util.SuccessStatus}

	userKey := fmt.Sprintf("%s:%s:%s", _const.RedisRootPrefix, req.Name, req.UUID)

	user, err := redis_client.Get(userKey)
	if err == redis.Nil {
		err = errors.New("user is not exist")
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	} else if err != nil {
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	}

	err = util.ReverseSerialize([]byte(user), req)
	if err != nil {
		log.Errorf("err : [%s]", err.Error())
		resp.Status = util.NewError(err)
		return resp
	}

	resp.Status.Data = req
	return resp
}
