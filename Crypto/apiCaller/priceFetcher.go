package apiCaller

type PriceFetcher interface {
	UpdatePrice() error
	UpdatePrices() error
}
