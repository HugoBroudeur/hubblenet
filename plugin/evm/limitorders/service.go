// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package limitorders

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type OrderBookAPI struct {
	db LimitOrderDatabase
}

func NewOrderBookAPI(database LimitOrderDatabase) *OrderBookAPI {
	return &OrderBookAPI{
		db: database,
	}
}

type OrderBookResponse struct {
	Orders []OrderMin
}

type OpenOrdersResponse struct {
	Orders []LimitOrder
}

type OrderMin struct {
	Market
	Price *big.Int
	Size  *big.Int
}

func (api *OrderBookAPI) GetDetailedOrderBookData(ctx context.Context) InMemoryDatabase {
	return api.db.GetOrderBookData()
}

func (api *OrderBookAPI) GetOrderBook(ctx context.Context, marketStr string) (*OrderBookResponse, error) {
	// market is a string cuz it's an optional param
	allOrders := api.db.GetAllOrders()
	orders := []OrderMin{}

	if len(marketStr) > 0 {
		market, err := strconv.Atoi(marketStr)
		if err != nil {
			return nil, fmt.Errorf("invalid market")
		}
		marketOrders := []LimitOrder{}
		for _, order := range allOrders {
			if order.Market == Market(market) {
				marketOrders = append(marketOrders, order)
			}
		}
		allOrders = marketOrders
	}

	for _, order := range allOrders {
		orders = append(orders, OrderMin{
			Market: order.Market,
			Price:  order.Price,
			Size:   order.GetUnFilledBaseAssetQuantity(),
		})
	}

	return &OrderBookResponse{Orders: orders}, nil
}

func (api *OrderBookAPI) GetOpenOrders(ctx context.Context, trader string) OpenOrdersResponse {
	traderOrders := []LimitOrder{}
	orderMap := api.db.GetOrderBookData().OrderMap
	for _, order := range orderMap {
		if strings.ToLower(order.UserAddress) == strings.ToLower(trader) {
			traderOrders = append(traderOrders, LimitOrder{
				Market: order.Market,
				Price: order.Price,
				BaseAssetQuantity: order.BaseAssetQuantity,
				FilledBaseAssetQuantity: order.FilledBaseAssetQuantity,
			})
		}
	}

	return OpenOrdersResponse{Orders: traderOrders}
}
