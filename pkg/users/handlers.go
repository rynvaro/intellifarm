package users

import (
	"cattleai/apicommon"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	token := "cattleaitoken"
	c.JSON(http.StatusOK, apicommon.SuccessResponse(token))
}

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, apicommon.SuccessResponse("info"))
}
