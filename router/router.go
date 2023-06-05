package router

import (
	"fmt"
	"net/http"
	"oauth2/service"
	"oauth2/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(util.CORS())
	r.LoadHTMLGlob("app/templates/**/*")
	r.StaticFS("/static", http.Dir("./app/static"))
	oauth2 := r.Group("/oauth2")
	{
		oauth2.POST("/register", util.HandlerWrapper(service.ClientRegister))
		oauth2.GET("/ask-auth", service.Authorization)
		oauth2.GET("/code", service.Code)
		oauth2.POST("/access-token", util.HandlerWrapper(service.AccessToken))
		oauth2.POST("/refresh-token", util.HandlerWrapper(service.RefreshToken))

	}
	if err := r.Run(":8078"); err != nil {
		fmt.Println(err)
		return
	}
}
