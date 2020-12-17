package cache

import (
	"cattleai/ent"
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

var L1Cache = cache.New(5*time.Minute, 10*time.Minute)

func GetUserByToken(token string) (*ent.User, error) {
	obj, ok := L1Cache.Get(token)
	if !ok {
		return nil, errors.New("not logged in")
	}
	currentUser := &ent.User{}
	json.Unmarshal([]byte(obj.(string)), currentUser)
	return currentUser, nil
}
