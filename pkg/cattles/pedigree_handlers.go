package cattles

import (
	"cattleai/db"
	"cattleai/ent/cattle"
	"cattleai/pkg/params"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func PedigreeHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}

	cattleSelf := db.Client.Cattle.GetX(c.Request.Context(), int(id.Id))

	mother := cattleSelf.Mother
	father := cattleSelf.Father

	grandpa := ""
	grandma := ""

	grandfather := ""
	grandmother := ""

	if mother != "" {
		motherCattles := db.Client.Cattle.Query().Where(cattle.EarNumberEQ(mother)).AllX(c.Request.Context())
		if len(motherCattles) > 0 {
			grandpa = motherCattles[0].Father
			grandma = motherCattles[0].Mother
		}

	}

	if father != "" {
		fatherCattles := db.Client.Cattle.Query().Where(cattle.EarNumberEQ(father)).AllX(c.Request.Context())
		if len(fatherCattles) > 0 {
			grandfather = fatherCattles[0].Father
			grandmother = fatherCattles[0].Mother
		}

	}

	ret := &PedigreeResp{
		Self:        cattleSelf.EarNumber,
		Father:      father,
		Mother:      mother,
		Grandfather: grandfather,
		Grandmother: grandmother,
		Grandma:     grandma,
		Grandpa:     grandpa,
	}

	c.JSON(http.StatusOK, resp.Success(ret))
}

type PedigreeResp struct {
	Self        string `json:"self"`
	Father      string `json:"father"`
	Mother      string `json:"mother"`
	Grandpa     string `json:"grandpa"`
	Grandma     string `json:"grandma"`
	Grandfather string `json:"grandfather"`
	Grandmother string `json:"grandmother"`
}
