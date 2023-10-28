package openapidoc

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func OpenapiGet(ctx *gin.Context) {
	jsonFile, err := os.Open("./cmd/demo-docker/openapi.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	ctx.JSON(http.StatusOK, result)
}
