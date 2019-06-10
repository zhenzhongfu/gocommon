package util

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func SetUserID(c *gin.Context, userID string) error {
	str := com.ToStr(userID)
	//TODO 这里的session放redis
	c.Set("TOKEN_USERID", str)
	return nil
}

func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get("TOKEN_USERID")
	if !ok {
		return 0, fmt.Errorf("can not get userid\n")
	}
	return com.StrTo(id.(string)).MustInt(), nil
}

func SetEmptyUserID(c *gin.Context, userID string) error {
	uid, err := GetUserID(c)
	if err != nil {
		if uid == 0 {
			str := com.ToStr(userID)
			//TODO 这里的session放redis
			c.Set("TOKEN_USERID", str)
			return nil
		} else {
			return nil
		}
	}

	return nil
}
