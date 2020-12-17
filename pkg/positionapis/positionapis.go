package positionapis

import (
	"cattleai/db"
	"cattleai/ent/positionapi"
	"context"
	"strconv"
	"strings"
)

func GetApisByPosition(positionId int64) ([]int64, error) {
	apis, err := db.Client.PositionApi.Query().Where(positionapi.PositionId(positionId)).All(context.Background())
	if err != nil {
		return nil, err
	}
	if len(apis) == 0 {
		return nil, err
	}
	ids := make([]int64, 0)
	arr := strings.Split(apis[0].Apis, ",")
	for i := range arr {
		id, _ := strconv.ParseInt(arr[i], 10, 64)
		ids = append(ids, id)
	}
	return ids, nil
}
