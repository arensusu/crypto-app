package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (handler *telegramHandler) setSearchHistoryToInlineKeyboard(id int64) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := [][]tgbotapi.InlineKeyboardButton{}
	pairs, err := handler.fundingUsecase.GetLastFiveFundingSearched(id)
	if err != nil {
		return tgbotapi.InlineKeyboardMarkup{}
	}

	for _, pair := range pairs {
		text := fmt.Sprintf("%s %s", pair.Exchange, pair.Symbol)
		data := fmt.Sprintf("/funding %s %s", pair.Exchange, pair.Symbol)
		inlineKeyboard = append(inlineKeyboard, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(text, data)))
	}
	return tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)
}

func (handler *telegramHandler) funding(id int64) string {
	history := handler.msgHistory[id]
	switch len(history) {
	case 1:
		inlineKeyboard := handler.setSearchHistoryToInlineKeyboard(id)
		msg := tgbotapi.NewMessage(id, "Recent Search")
		msg.ReplyMarkup = inlineKeyboard
		if _, err := handler.tgbot.Send(msg); err != nil {
			log.Println(err)
		}
		return "Or enter exchange."
	case 2:
		return "Symbol?"
	case 3:
		handler.msgHistory[id] = []string{}
		data, err := handler.fundingUsecase.GetFundingData(history[1], history[2])
		if err != nil {
			return "Cannot get the funding rate. Please try again."
		}

		if err := handler.fundingUsecase.AddFundingSearched(id, history[1], history[2]); err != nil {
			fmt.Println(err)
		}

		msg := "Funding Rate\n"
		msg += fmt.Sprintf("\n%s %s\n", data.Exchange, data.Symbol)
		msg += fmt.Sprintf("Total of last 100: %.4f%%\n", data.Last100)
		msg += fmt.Sprintf("Total of last 30: %.4f%%\n", data.Last30)
		msg += fmt.Sprintf("Last: %.4f%%\n", data.Last)
		return msg
	default:
		return "show error"
	}

}
