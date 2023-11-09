package api

import (
	"crypto-exchange/pkg/crossexchange"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CrossExchangeServer struct {
	exchange *crossexchange.CrossExchangeSingleSymbol
}

func NewCrossExchangeServer(r *gin.Engine, exchange *crossexchange.CrossExchangeSingleSymbol) {
	server := &CrossExchangeServer{exchange: exchange}
	r.GET("/crossexchange/:symbol", server.GetExchangeData)

}

func (s *CrossExchangeServer) GetExchangeData(c *gin.Context) {
	symbol := c.Param("symbol")

	data := s.exchange.GetExchangeData(symbol)

	c.JSON(http.StatusOK, data)
}
