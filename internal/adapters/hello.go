package adapters

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloHandler interface {
	Hello(c *gin.Context)
	Health(c *gin.Context)
}

type helloHandler struct {
	db *gorm.DB
}

func NewHelloHandler(db *gorm.DB) HelloHandler {
	return &helloHandler{
		db: db,
	}
}

type Hello struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// @Summary Hello
// @Description Hello World
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {object} Hello
// @Router /hello [get]
func (h *helloHandler) Hello(c *gin.Context) {
	c.JSON(200, Hello{
		Message:   "Hello World",
		Timestamp: time.Now(),
	})
}

// @Summary Health
// @Description Health Check
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {object} Hello
// @Router /health [get]
func (h *helloHandler) Health(c *gin.Context) {

	if err := h.db.Exec("SELECT 1").Error; err != nil {
		c.JSON(500, Hello{
			Message:   "Database is not available",
			Timestamp: time.Now(),
		})

		return
	}

	c.JSON(200, Hello{
		Message:   "Health Check",
		Timestamp: time.Now(),
	})
}
