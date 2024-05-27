package adapters

import (
	"go-privfile/internal/adapters/dtos"
	"go-privfile/internal/domain/shortner"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShortnerHandler interface {
	Shorten(c *gin.Context)
	Redirect(c *gin.Context)
}

type shortnerHandler struct {
	service shortner.ShortnerService
}

func NewShortnerHandler(service shortner.ShortnerService) ShortnerHandler {
	return &shortnerHandler{
		service: service,
	}
}

// @Summary Shorten
// @Description Shorten URL
// @Tags Shortner
// @Accept json
// @Produce json
// @Param shorten body dtos.ShortnerRequest true "Shortner Request"
// @Success 201 {object} shortner.Shortner
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /shorten [post]
func (h *shortnerHandler) Shorten(c *gin.Context) {
	var req dtos.ShortnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dtos.NewErrorResponse(err.Error(), 400))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, dtos.NewErrorResponse(err.Error(), 400))
		return
	}

	shortner := shortner.Shortner{
		Url:      req.Url,
		Redirect: req.AutoRedirect,
	}

	if err := h.service.Store(&shortner); err != nil {
		c.JSON(500, dtos.NewErrorResponse(err.Error(), 500))
		return
	}

	c.JSON(201, shortner)
}

// @Summary Redirect
// @Description Redirect URL
// @Tags Shortner
// @Accept json
// @Produce json
// @Param code path string true "Shortner Code"
// @Param body query bool false "Return Shortner Body"
// @Success 301 {object} shortner.Shortner
// @Failure 404 {object} dtos.ErrorResponse
// @Router /{code} [get]
func (h *shortnerHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	body, err := strconv.ParseBool(c.DefaultQuery("body", "false"))
	if err != nil {
		c.JSON(400, dtos.NewErrorResponse("Invalid 'body' query parameter", 400))
		return
	}

	shortner, err := h.service.Find(code)
	if err != nil {
		c.JSON(404, dtos.NewErrorResponse(err.Error(), 404))
		return
	}

	if !body && shortner.Redirect {
		c.Redirect(301, shortner.Url)
		return
	}

	c.JSON(200, shortner)
}
