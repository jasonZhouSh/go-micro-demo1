package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-micro-demo1/time-service/client"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
)

func GetTime(c *gin.Context)  {
	name := c.Param("name")
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	tResp,err := client.GetUserTime(name,reg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data" : err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : tResp.User.CrateTime,
	})
}