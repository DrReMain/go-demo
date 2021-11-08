package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/remainlab/go-vue/common"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":8080"))
}
