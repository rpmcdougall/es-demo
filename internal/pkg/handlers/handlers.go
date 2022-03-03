package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rpmcdougall/es-demo/internal/pkg/repository"
	"go.uber.org/fx"
	"net/http"
)

type ApiHandler struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

type ApiHandlers struct {
	Repository repository.Repository
	Handlers   []ApiHandler
}

func (h *ApiHandlers) EsHandler() ApiHandler {
	return ApiHandler{
		Method: http.MethodGet,
		Path:   "/es/info",
		HandlerFunc: func(c *gin.Context) {
			info, err := h.Repository.GetInfo()
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, info)
		},
	}

}

func (h *ApiHandlers) register(handlers ...ApiHandler) *ApiHandlers {
	for _, handler := range handlers {
		h.Handlers = append(h.Handlers, handler)
	}
	return h
}

type NewApiHandlersParams struct {
	fx.In
	Repository *repository.LiveRepository
}

func NewApiHandlers(p NewApiHandlersParams) *ApiHandlers {
	handlers := &ApiHandlers{
		Repository: p.Repository,
		Handlers:   []ApiHandler{},
	}

	registered := handlers.register(
		handlers.EsHandler(),
	)

	return registered
}
