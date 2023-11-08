package api

import (
	"crypto-exchange/pkg/crossexchange"
	"net/http"

	"github.com/gorilla/mux"
)

type CrossExchangeServer struct {
	exchange *crossexchange.CrossExchangeSingleSymbol
}

func NewCrossExchangeServer(r *mux.Router, exchange *crossexchange.CrossExchangeSingleSymbol) {
	server := &CrossExchangeServer{exchange: exchange}
	r.HandleFunc("/crossexchange/{symbol}", server.GetExchangeData).Methods("GET")

}

func (s *CrossExchangeServer) GetExchangeData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	data := s.exchange.GetExchangeData(params["symbol"])

	ResponseWithJson(w, http.StatusOK, data)
}
