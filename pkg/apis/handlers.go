package apis

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/api"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SubAPIHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	wheres := []predicate.API{api.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, api.NameContains(listParams.Q))
	}
	wheres = append(wheres, api.LevelEQ(listParams.Level))
	wheres = append(wheres, api.ParentIdEQ(listParams.ParentId))
	apis, err := db.Client.API.Query().Where(api.And(wheres...)).Order(ent.Desc(api.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	apids := make([]int64, len(apis))
	for i := range apis {
		apids[i] = apis[i].ID
	}

	subapis, err := db.Client.API.Query().Where(
		api.ParentIdIn(apids...),
	).All(c.Request.Context())

	m := map[int64][]*ent.API{}
	for i := range subapis {
		if m[subapis[i].ParentId] == nil {
			m[subapis[i].ParentId] = []*ent.API{subapis[i]}
		} else {
			m[subapis[i].ParentId] = append(m[subapis[i].ParentId], subapis[i])
		}
	}

	apiresp := make([]*subapiresp, 0)
	for i := range apis {
		apiresp = append(apiresp, &subapiresp{Api: apis[i], Children: m[apis[i].ID]})
	}

	c.JSON(http.StatusOK, resp.Success(apiresp))
}

type subapiresp struct {
	Api      *ent.API   `json:"api"`
	Children []*ent.API `json:"children"`
}
