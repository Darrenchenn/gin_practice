package util

import (
	_const "awesomeProject/src/const"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseParameters(c *gin.Context, obj interface{}) error {
	err := c.ShouldBind(obj)
	if err != nil {
		return err
	}

	return nil
}

func SetOutput(c *gin.Context, obj interface{}) {
	var sj = _const.Status{Code: 200, Message: "OK", Data: obj}
	c.JSON(http.StatusOK, &sj)
}

func Serialize(data interface{}) (string, error) {
	byte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func ReverseSerialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}
