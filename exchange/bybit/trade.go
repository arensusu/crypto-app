package bybit

import (
	"crypto-exchange/domain"
	"errors"
	"strconv"
	"time"

	"github.com/hirokisan/bybit/v2"
)

func (ex *Bybit) MarketTrade(p domain.MarketTradeParam) error {
	id := "susu" + strconv.FormatInt(time.Now().UnixMilli()%1000, 10)
	param := bybit.V5CreateOrderParam{
		Category:    bybit.CategoryV5Linear,
		Symbol:      bybit.SymbolV5(p.Symbol),
		Side:        bybit.Side(p.Side),
		OrderType:   bybit.OrderTypeMarket,
		Qty:         p.Quantity,
		OrderLinkID: &id,
	}
	res, err := ex.Client.V5().Order().CreateOrder(param)
	if err != nil {
		return err
	}

	if res.RetMsg != "OK" {
		return errors.New(res.RetMsg)
	}

	return nil
}
