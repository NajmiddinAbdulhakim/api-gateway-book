package api

import (
	han "github.com/NajmiddinAbdulhakim/ude/api-gateway/api/handler"
	"github.com/NajmiddinAbdulhakim/ude/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/ude/api-gateway/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/NajmiddinAbdulhakim/ude/api-gateway/api/docs" // swag
)

type Option struct {
	Conf           config.Config
	ServiceManager service.IServiceManager
}

func New(o *Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := han.New(&han.HandlerConfig{
		ServiceManager: o.ServiceManager,
		Cfg:            o.Conf,
	})

	book := router.Group("/book")
	{
		book.POST("/create", handler.CreateBook)
		book.GET(`/:id`, handler.GetBookById)
		book.PUT(`/:id`, handler.UpdateBook)
		book.DELETE(`/:id`, handler.DeleteBookById)
	}

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run()
	return router
}
