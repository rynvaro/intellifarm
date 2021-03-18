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
		SetLevel(form.Level).
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
		SetTenantName(form.TenantName).
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
	listParams.Level = c.MustGet("level").(int)
	log.Debug().Msg(listParams.ToString())
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
	err := db.Client.User.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
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
		SetLevel(form.Level).
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
		SetTenantName(form.TenantName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(user))
}
