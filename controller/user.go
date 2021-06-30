package controller

import (
	"github.com/HikaruEgashira/simple-server/lib"
	"github.com/HikaruEgashira/simple-server/usecase"
	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	name := c.Query("name")

	text := usecase.Hello(name)

	lib.Render(c.Writer, "pages/user", map[string]string{"text": text})
}
