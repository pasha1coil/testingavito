package handler

import (
	"github.com/pasha1coil/testingavito/pkg/service"

	"github.com/gin-gonic/gin"
	_ "github.com/pasha1coil/testingavito/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	command := router.Group("/main")
	{
		command.POST("/adduser", h.AddUser)
		command.POST("/addsegment", h.AddSegm)
		command.DELETE("/delsegment", h.DelSegm)
		command.POST("/insertsegmentuser", h.InsSegmUsr)
		command.DELETE("/delsegmentuser", h.DelSegmUsr)
		command.POST("/getusersegment", h.GetUsrSegm)
		command.GET("/history", h.GetSlugHistoryCsv)
	}
	router.Static("/static", "./static")
	//http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
