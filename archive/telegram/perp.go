package telegram

import (
	"fmt"
)

func (handler *telegramHandler) perp(id int64) string {
	history := handler.msgHistory[id]
	switch len(history) {
	case 1:
		return "Exchange?"
	case 2:
		return "Symbol?"
	case 3:
		handler.msgHistory[id] = []string{}
		perp, err := handler.fundingUsecase.GetPerpData(history[1], history[2])
		if err != nil {
			fmt.Println(err)
			return "Cannot get the perpetual data. Please try again later."
		}

		msg := fmt.Sprintf("%s %s\n", perp.Exchange, perp.Symbol)
		msg += fmt.Sprintf("Price: %f (%.2f%%)\n", perp.Price, perp.PriceChangePercent)
		msg += fmt.Sprintf("Next Funding: %.4f", perp.FundingRate)
		return msg
	default:
		return "show error"
	}

}
