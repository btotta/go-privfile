package server

import (
	"net/http"

	docs "go-privfile/docs"
	"go-privfile/internal/adapters"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	helloHandler adapters.HelloHandler
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	s.instantiateDependencies()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/hello")
	})

	hello := r.Group("/")
	{
		hello.GET("/hello", helloHandler.Hello)
		hello.GET("/health", helloHandler.Health)
	}

	return r
}

func (s *Server) instantiateDependencies() {
	helloHandler = adapters.NewHelloHandler(s.db.GetDB())
}
