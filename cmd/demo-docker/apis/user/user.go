package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/confserver"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/controller"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/global"
	"github.com/rookie-luochao/gin-demo/internal/utils"
	"net/http"
)

// ListUser
// @BasePath /api/v1
// PingExample godoc
// @Summary ListUser
// @Schemes
// @Description 获取用户
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param param query controller.ListUserParams true "查询条件"
// @Success 200 {object} controller.ListUserResp 成功
// @Failure 400 {object} utils.ErrorResp 失败
// @Failure 500 {object} utils.ErrorResp 失败
// @Router /demo-docker/api/v1/users [get]
// @ID ListUser
func ListUser(c *gin.Context) {
	param := controller.ListUserParams{}
	err := confserver.Bind(c, &param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}
	resp, err := global.UserMgr.ListUser(confserver.RequestContextFromGinContext(c), &param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

type CreateOrUpdateUserBody struct {
	controller.CreateOrUpdateUserData `in:"body"`
}

// CreateOrUpdateUser
// @BasePath /api/v1
// PingExample godoc
// @Summary CreateOrUpdateUser
// @Schemes
// @Description 创建用户
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Body body CreateOrUpdateUserBody true "活动"
// @Success 200 {object} utils.ErrorResp 成功
// @Failure 400 {object} utils.ErrorResp 失败
// @Failure 500 {object} utils.ErrorResp 失败
// @Router /demo-docker/api/v1/users [post]
// @ID CreateOrUpdateUser
func CreateOrUpdateUser(c *gin.Context) {
	body := CreateOrUpdateUserBody{}
	err := confserver.Bind(c, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}

	err = global.UserMgr.CreateOrUpdateUser(body.CreateOrUpdateUserData)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
