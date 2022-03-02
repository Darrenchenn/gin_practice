package middleware

import (
	"awesomeProject/src/redis_client"

	"github.com/gin-gonic/gin"
)

func InitRedis() gin.HandlerFunc {
	return func(c *gin.Context) {
		redis_client.NewClient("10.30.12.17:6379", "")

		c.Next()

		err := redis_client.Close()
		if err != nil {
			log.Errorf("close redis client err:[%s]", err.Error())
			return
		}
	}
}
