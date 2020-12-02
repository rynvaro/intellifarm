package users

import (
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	token := "cattleaitoken"
	c.JSON(http.StatusOK, resp.Success(token))
}

func UserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Default())
}

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Success("info"))
}
