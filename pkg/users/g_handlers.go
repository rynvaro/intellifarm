package users

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/user"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func UserAddHandler(c *gin.Context) {
	form := &ent.User{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	user, err := db.Client.User.Create().
		SetAddress(form.Address).
		SetAge(form.Age).
		SetDutyName(form.DutyName).
		SetEducation(form.Education).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetGender(form.Gender).
		SetIdCard(form.IdCard).
		SetJobTitle(form.JobTitle).
		SetJoinedAt(form.JoinedAt).
		SetMajor(form.Major).
		SetName(form.Name).
		SetOnJobState(form.OnJobState).
		SetPhone(form.Phone).
		SetPositionId(form.PositionId).
		SetPositionName(form.PositionName).
		SetRemarks(form.Remarks).
		SetPassword(form.Password).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(user))
}

func UserListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.User.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	users, err := db.Client.User.Query().Where(where).Order(ent.Desc(user.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   users,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func UserDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	user, err := db.Client.User.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(user))
}

func UserUpdateHandler(c *gin.Context) {
	form := &ent.User{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	user, err := db.Client.User.UpdateOneID(form.ID).
		SetAddress(form.Address).
		SetAge(form.Age).
		SetDutyName(form.DutyName).
		SetEducation(form.Education).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetGender(form.Gender).
		SetIdCard(form.IdCard).
		SetJobTitle(form.JobTitle).
		SetJoinedAt(form.JoinedAt).
		SetMajor(form.Major).
		SetName(form.Name).
		SetOnJobState(form.OnJobState).
		SetPhone(form.Phone).
		SetPositionId(form.PositionId).
		SetPositionName(form.PositionName).
		SetRemarks(form.Remarks).
		SetPassword(form.Password).
		SetTenantId(form.TenantId).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(user))
}
