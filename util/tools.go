package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrWrapper(err error) (int, *Status) {
	return ErrorMap[err].HttpCode, ErrorMap[err].Status
}

func HandlerWrapper(f func(c *gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := f(c)
		if err == nil && data == nil {
			return
		}
		if err == nil {
			c.JSON(http.StatusOK, data)
			return
		}
		for k, _ := range ErrorMap {
			if err == k {
				c.JSON(ErrWrapper(k))
				return
			}
		}
		c.JSON(ErrWrapper(CommonErr))
		return
	}
}

func GetMD5String(b []byte) string {
	hash := md5.New()
	hash.Write(b)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetSHAStringPass(b []byte) string {
	hash := sha1.New()
	hash.Write(b)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
