package sheds

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/shed"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ShedAddHandler(c *gin.Context) {
	form := &ent.Shed{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	shed, err := db.Client.Shed.Create().
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetHeight(form.Height).
		SetLength(form.Length).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedCateId(form.ShedCateId).
		SetShedCateName(form.ShedCateName).
		SetShedTypeId(form.ShedTypeId).
		SetShedTypeName(form.ShedTypeName).
		SetSquare(form.Square).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		SetWidth(form.Width).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(shed))
}

func ShedListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Shed.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	sheds, err := db.Client.Shed.Query().Where(where).Order(ent.Desc(shed.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   sheds,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func ShedDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Shed.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func ShedUpdateHandler(c *gin.Context) {
	form := &ent.Shed{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	shed, err := db.Client.Shed.UpdateOneID(form.ID).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetHeight(form.Height).
		SetLength(form.Length).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedCateId(form.ShedCateId).
		SetShedCateName(form.ShedCateName).
		SetShedTypeId(form.ShedTypeId).
		SetShedTypeName(form.ShedTypeName).
		SetSquare(form.Square).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		SetWidth(form.Width).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(shed))
}
