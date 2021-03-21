package materials

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/material"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func MaterialAddHandler(c *gin.Context) {
	form := &ent.Material{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	material, err := db.Client.Material.Create().
		SetCategoryId(form.CategoryId).
		SetCategoryName(form.CategoryName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetMaterialId(form.MaterialId).
		SetCode(form.Code).
		SetInventory(form.Inventory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(material))
}

func MaterialListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Material.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	materials, err := db.Client.Material.Query().Where(where).Order(ent.Desc(material.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   materials,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func MaterialDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Material.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func MaterialUpdateHandler(c *gin.Context) {
	form := &ent.Material{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	material, err := db.Client.Material.UpdateOneID(form.ID).
		SetCategoryId(form.CategoryId).
		SetCategoryName(form.CategoryName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetMaterialId(form.MaterialId).
		SetCode(form.Code).
		SetInventory(form.Inventory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(material))
}
