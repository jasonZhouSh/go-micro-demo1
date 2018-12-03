package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-micro-demo1/user-service/client"
	"github.com/golang/glog"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
)

func GetTimeOrAddUser(c *gin.Context)  {
	name := c.Param("name")
	glog.Info("name:",name)
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	uResp,err := client.GetTimeOrAddUser(name,reg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data" : err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : uResp.Time,
	})
}