package apis

import (
	"sky/common/tools/app"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	app.OK(c, "", "ok")
}
