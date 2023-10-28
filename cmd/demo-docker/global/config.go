package global

import (
	"github.com/kunlun-qilian/conflogger"
	"github.com/kunlun-qilian/confserver"
	"github.com/rookie-luochao/confx"
	"github.com/rookie-luochao/gin-demo/cmd/demo-docker/controller"
)

func init() {
	confx.SetConfX("demo-docker", "..", confx.DockerConfig{
		BuildImage:   "golang:1.20-bullseye",
		RuntimeImage: "debian:11",
		GoProxy: confx.GoProxyConfig{
			ProxyOn: true,
			Host:    "https://goproxy.cn,direct",
		},
		Openapi: true,
	})
	confx.ConfP(&Config)
}

var UserMgr *controller.UserMgr

var Config = struct {
	Logger  *conflogger.Log
	Server  *confserver.Server
	TestEnv string `env:""`

	TestPort int `env:"opt,expose"`
}{
	Server: &confserver.Server{
		Port: 80,
		Mode: "debug",
	},
	TestEnv:  "123",
	TestPort: 9090,
}
