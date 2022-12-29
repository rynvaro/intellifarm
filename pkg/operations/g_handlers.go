package operations

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/operation"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func OperationAddHandler(c *gin.Context) {
	form := &ent.Operation{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	operation, err := db.Client.Operation.Create().
		SetAPI(form.API).
		SetIP(form.IP).
		SetMethod(form.Method).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(operation))
}

func OperationListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Operation.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	operations, err := db.Client.Operation.Query().Where(where).Order(ent.Desc(operation.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   operations,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func OperationDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Operation.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func OperationUpdateHandler(c *gin.Context) {
	form := &ent.Operation{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	operation, err := db.Client.Operation.UpdateOneID(form.ID).
		SetAPI(form.API).
		SetIP(form.IP).
		SetMethod(form.Method).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(operation))
}
