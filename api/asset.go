package api

import (
	"net/http"

	"github.com/arensusu/crypto-app/pkg/assets"

	"github.com/gin-gonic/gin"
)

type AssetsServer struct {
	exchange *assets.AssetsFinder
}

func NewAssetsServer(r *gin.Engine, exchange *assets.AssetsFinder) {
	server := &AssetsServer{exchange: exchange}
	r.GET("/assets", server.GetExchangeData)

}

func (s *AssetsServer) GetExchangeData(c *gin.Context) {

	data := s.exchange.GetAssets()

	c.JSON(http.StatusOK, data)
}
