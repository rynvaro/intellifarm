package cattles

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattle"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CattleSearchParams struct {
	ShedId       int64 `json:"shedId" form:"shedId"`
	InGroup      int   `json:"inGroup" form:"inGroup"` // 是否在群
	CattleCateId int   `json:"cattleCateId" form:"cattleCateId"`
	params.ListParams
}

func CattleListHandler(c *gin.Context) {
	listParams := &CattleSearchParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	fmt.Printf("--%+v\n", listParams)
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.Cattle.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattles, err := db.Client.Cattle.Query().Where(where).Order(ent.Desc(cattle.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattles,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}
