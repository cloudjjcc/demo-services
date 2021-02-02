package handler

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/cloudjjcc/demo-services/product/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ http.Handler = (*ApiHandler)(nil)

type ApiHandler struct {
	*gin.Engine
	srv micro.Service
}

func NewApiHandler(srv micro.Service) *ApiHandler {
	router := gin.Default()
	apiHandler := &ApiHandler{Engine: router, srv: srv}
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/prods", apiHandler.GetProds)
	}
	return apiHandler
}

func (api *ApiHandler) GetProds(ctx *gin.Context) {
	prodService := proto.NewProductService("product", api.srv.Client())
	list, err := prodService.GetProducts(context.Background(), &proto.ProductRequest{
		Query: "name",
		Page: &proto.Page{
			PageNum:  1,
			PageSize: 20,
		},
	})
	if err != nil {
		ResponseWithErr(ctx, err)
		return
	}
	ResponseWithData(ctx, list)
}
