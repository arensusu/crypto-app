package api

import (
	"crypto-exchange/exchange/domain"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FundingHandler struct {
	usecase domain.PairUsecase
}

func NewFundingHandler(r *mux.Router, fundingUsecase domain.PairUsecase) {
	handler := &FundingHandler{fundingUsecase}
	r.HandleFunc("/funding/{exchange}/{symbol}", handler.GetFundingHistory)
	r.HandleFunc("/perp/{exchange}/{symbol}", handler.GetPerpetual)
}

func ResponseWithJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Print(err)
	}
}

func (handler *FundingHandler) GetFundingHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchange := vars["exchange"]
	symbol := vars["symbol"]

	fundingData, err := handler.usecase.GetFundingData(exchange, symbol)
	if err != nil {
		ResponseWithJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ResponseWithJson(w, http.StatusOK, fundingData)
}

func (handler *FundingHandler) GetPerpetual(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exchange := vars["exchange"]
	symbol := vars["symbol"]

	perpData, err := handler.usecase.GetPerpData(exchange, symbol)
	if err != nil {
		ResponseWithJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ResponseWithJson(w, http.StatusOK, perpData)
}
