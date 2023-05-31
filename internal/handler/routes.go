package handler

import (
	"diploma/internal/handler/utils"
	"diploma/internal/service"
	"diploma/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   service.Services
	logger    *logger.Logger
	responder *utils.Responder
}

func NewHandler(service service.Services,logger *logger.Logger) *Handler {
	return &Handler{
		service:   service,
		logger:    logger,
		responder: utils.NewResponder(logger),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} 
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "x-data"}
	router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
	}
	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/client", h.GetUsers)

			users.PUT("/client", h.PutUsers)

			users.DELETE("/client", h.DeleteUsers)

		}
		api.GET("sites/", h.GetListSites)
		site := api.Group("/sites")
		{
			site.GET("/site", h.GetSite)

			site.POST("/site", h.PostSite)

			site.DELETE("/site", h.DeleteSite)

		}
		text := api.Group("/texts")
		{
			text.GET("/text", h.GetText)

			text.POST("/text", h.PostText)

			text.PUT("/text", h.PutText)
		}
		api.PUT("/passwordChange", h.ChangePassword)
		api.POST("/parse", h.ParseSite)
		api.GET("/download", h.DocHandler)
	}
	
	return router
}
