package commands

import (
	"password/priceCache"
	"password/users"
)

type UserCacheNeededCommands interface {
	GetSummaryUsingCache(user *users.User, cacher *priceCache.PriceCache)
}
