package handler

import (
	"diploma/internal/handler/utils"
	"diploma/internal/service"
	"diploma/pkg/logger"
	"diploma/pkg/pgsql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   service.Services
	db *pgsql.Postgres
	logger *logger.Logger
	responder *utils.Responder
}

func NewHandler(service service.Services, db *pgsql.Postgres, logger *logger.Logger) *Handler {
	return &Handler{
		service:   service,
		db : db,
		logger: logger,
		responder: utils.NewResponder(logger),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

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

			users.POST("/client", h.PostUsers)

			users.PUT("/client", h.PutUsers)

			users.DELETE("/client", h.DeleteUsers)

		}
		site := api.Group("/sites")
		{
			site.GET("/site", h.GetSites)

			site.POST("/site", h.PostSites)

			site.PUT("/site", h.PutSites)

			site.DELETE("/site", h.DeleteSites)
		}

		text := api.Group("/texts")
		{
			text.GET("/text", h.GetTexts)

			text.POST("/text", h.PostTexts)

			text.PUT("/text", h.PutTexts)

			text.DELETE("/text", h.DeleteTexts)

		}
		page := api.Group("/pages")
		{
			page.GET("/page", h.GetPages)

			page.POST("/page", h.PostPages)

			page.PUT("/page", h.PutPages)

			page.DELETE("/page", h.DeletePages)
		}
	}

	return router
}
