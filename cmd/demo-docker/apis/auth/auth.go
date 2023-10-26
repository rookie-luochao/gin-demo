package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/confserver"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthorizationParam struct {
	Authorization string `name:"Authorization,omitempty" in:"header" `
}

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a := AuthorizationParam{}
		err := confserver.Bind(ctx, &a)
		if err != nil {
			logrus.Warn("Authorization is nil")
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.Next()
	}
}
