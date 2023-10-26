package hello

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunlun-qilian/confserver"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/apis/auth"
	"github.com/rookie-luochao/gin-demo/internal/utils"
	"net/http"
	//"runtime/trace"
)

func HelloRouter(r *gin.RouterGroup) {
	r.GET("/hello", auth.Authorization(), HelloGet)
	//r.GET("/hello/:age/:sex", auth.Authorization(), HelloGetWithPath)
	r.POST("/hello", auth.Authorization(), HelloPost)
}

type HelloResp struct {
	// 结果
	Data string `json:"data"`
}

type HelloGetParam struct {
	Name string `in:"query" name:"name"`
}

// @BasePath /
// PingExample godoc
// @Summary HelloGet
// @Schemes
// @Description HelloGet
// @Tags hello
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Token"
// @Param name query string true "Name"
// @Success 200 {object} HelloResp 成功
// @Failure 400 {object} utils.ErrorResp 失败
// @Failure 500 {object} utils.ErrorResp 失败
// @Router /demo-docker/api/v1/hello [get]
// @ID HelloGet
func HelloGet(ctx *gin.Context) {

	fmt.Println("Server header:", ctx.Request.Header)

	req := HelloGetParam{}
	err := confserver.Bind(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, HelloResp{
		Data: fmt.Sprintf("hello %s", req.Name),
	})
}

//type HelloGetPathParam struct {
//	HelloGetParam
//	Age int    `in:"path" name:"age"`
//	Sex string `in:"path" name:"sex"`
//}

// @BasePath /
// PingExample godoc
// @Summary HelloGetWithPath
// @Schemes
// @Description HelloGetWithPath
// @Tags hello
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Token"
// @Param name query string true "名字"
// @Param age path  int true "年龄"
// @Success 200 {object} HelloResp 成功
// @Failure 400 {object} utils.ErrorResp 失败
// @Failure 500 {object} utils.ErrorResp 失败
// @Router /demo-docker/api/v1/hello/{age} [get]
// @ID HelloGetWithPath
//func HelloGetWithPath(ctx *gin.Context) {
//
//	span := trace.GetTraceSpanFromContext(ctx)
//	span = trace.Start(span.Context(), ctx.HandlerName())
//	defer span.End()
//
//	fmt.Println("Server header:", ctx.Request.Header)
//
//	req := HelloGetPathParam{}
//	err := confserver.Bind(ctx, &req)
//	if err != nil {
//		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, HelloResp{
//		Data: fmt.Sprintf("hello %s age: %d sex:%s", req.Name, req.Age, req.Sex),
//	})
//}

type HelloPostParam struct {
	// Post名字
	HelloPostBody `in:"body"`
}

type HelloPostBody struct {
	Name string `json:"name"`
}

// @BasePath /
// PingExample godoc
// @Summary HelloGet
// @Schemes
// @Description HelloPost
// @Tags hello
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Token"
// @Param body body HelloPostParam true "body"
// @Success 200 {object} HelloResp 成功
// @Failure 400 {object} utils.ErrorResp 失败
// @Failure 500 {object} utils.ErrorResp 失败
// @Router /demo-docker/api/v1/hello [post]
// @ID HelloPost
func HelloPost(ctx *gin.Context) {
	req := HelloPostParam{}
	err := confserver.Bind(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResp{Result: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, HelloResp{
		Data: fmt.Sprintf("hello %s", req.Name),
	})
}
