package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/apis/auth"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/apis/hello"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/apis/user"
)

func NewRooter(r *gin.Engine) {
	v1 := r.Group("/demo-docker/api/v1")
	hello.HelloRouter(v1)

	v1.POST("/users/create-or-update-user", user.CreateOrUpdateUser)
	v1.GET("/users", auth.Authorization(), user.ListUser)
}
