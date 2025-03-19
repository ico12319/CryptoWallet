package priceCache

import (
	"sync"
	"time"
)

// PriceCache will be implemented using a Singleton design pattern so we have one global object that
// every class can use
type PriceCache struct {
	prices   map[string]PriceInfo
	duration time.Duration
}

var instance *PriceCache
var once sync.Once

func newPriceCache(duration time.Duration) *PriceCache {
	return &PriceCache{prices: make(map[string]PriceInfo), duration: duration}
}

func (pc *PriceCache) GetPrice(assetId string) (float64, bool) {
	pInfo, exist := pc.prices[assetId]
	if !exist || time.Since(pInfo.timeStamp) > pc.duration {
		return 0, false
	}
	return pInfo.price, true
}

func (pc *PriceCache) SetPrice(assetId string, price float64) {
	pc.prices[assetId] = PriceInfo{price: price, timeStamp: time.Now()}
}

func GetInstance() *PriceCache {
	once.Do(func() {
		instance = newPriceCache(time.Minute)
	})
	return instance
}
