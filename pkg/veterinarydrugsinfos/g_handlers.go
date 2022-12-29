package veterinarydrugsinfos

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/veterinarydrugsinfo"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func VeterinaryDrugsInfoAddHandler(c *gin.Context) {
	form := &ent.VeterinaryDrugsInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	veterinarydrugsinfo, err := db.Client.VeterinaryDrugsInfo.Create().
		SetCode(form.Code).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStopAt(form.StopAt).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(veterinarydrugsinfo))
}

func VeterinaryDrugsInfoListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.VeterinaryDrugsInfo.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	veterinarydrugsinfos, err := db.Client.VeterinaryDrugsInfo.Query().Where(where).Order(ent.Desc(veterinarydrugsinfo.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   veterinarydrugsinfos,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func VeterinaryDrugsInfoDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.VeterinaryDrugsInfo.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func VeterinaryDrugsInfoUpdateHandler(c *gin.Context) {
	form := &ent.VeterinaryDrugsInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	veterinarydrugsinfo, err := db.Client.VeterinaryDrugsInfo.UpdateOneID(form.ID).
		SetCode(form.Code).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStopAt(form.StopAt).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(veterinarydrugsinfo))
}
