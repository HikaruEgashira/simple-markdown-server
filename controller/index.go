package controller

import (
	"github.com/HikaruEgashira/simple-server/lib"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	lib.Render(c.Writer, "pages/index")
}
