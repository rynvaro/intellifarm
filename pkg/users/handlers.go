package users

import (
	"cattleai/cache"
	"cattleai/db"
	"cattleai/ent/api"
	"cattleai/ent/user"
	"cattleai/pkg/positionapis"
	"cattleai/resp"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type LoginParams struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	TenantId int64  `json:"tenantId" form:"tenantId" binding:"required"`
}

func UserLogin(c *gin.Context) {
	form := &LoginParams{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	rets, err := db.Client.User.Query().Where(
		user.And(
			user.NameEQ(form.Username),
			user.TenantId(form.TenantId),
		),
	).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, resp.Error(nil))
		return
	}

	if len(rets) != 1 {
		c.JSON(http.StatusNotFound, resp.Error(nil))
		return
	}

	currentUser := rets[0]
	// TODO encrypt password
	if form.Password != currentUser.Password {
		c.JSON(http.StatusInternalServerError, resp.Response(10400, "invalid password", nil))
		return
	}
	token := uuid.New().String()
	userJsonData, _ := json.Marshal(currentUser)
	cache.L1Cache.Add(token, string(userJsonData), time.Hour*24)

	if _, err := db.Client.Operation.Create().
		SetUserId(currentUser.ID).
		SetUserName(currentUser.Name).
		SetTenantId(currentUser.TenantId).
		SetTenantName(currentUser.TenantName).
		SetDeleted(0).
		SetMethod(c.Request.Method).
		SetIP(c.ClientIP()).
		SetAPI(c.Request.URL.String()).
		SetCreatedAt(time.Now().Unix()).
		Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
	}

	c.JSON(http.StatusOK, resp.Success(token))
}

func UserLogout(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	cache.L1Cache.Delete(token)
	c.JSON(http.StatusOK, resp.Default())
}

func UserInfo(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	currentUser, err := cache.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response(10401, "请重新登陆", nil))
		return
	}

	apids, err := positionapis.GetApisByPosition(currentUser.PositionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response(10402, "内部错误", nil))
		return
	}

	apis, err := db.Client.API.Query().Where(
		api.And(api.LevelEQ(1), api.IDIn(apids...)),
	).All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response(10402, "内部错误", nil))
		return
	}

	ret := map[string]interface{}{
		"user": currentUser,
		"apis": apis,
	}

	c.JSON(http.StatusOK, resp.Success(ret))
}
