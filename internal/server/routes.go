package server

import (
	"net/http"

	docs "go-privfile/docs"
	"go-privfile/internal/adapters"
	"go-privfile/internal/domain/shortner"
	"go-privfile/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	helloHandler    adapters.HelloHandler
	shortnerHandler adapters.ShortnerHandler
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	s.instantiateDependencies()

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/hello")
	})

	hello := r.Group("/")
	{
		hello.GET("/hello", helloHandler.Hello)
		hello.GET("/health", helloHandler.Health)
	}

	shortner := r.Group("/")
	{
		shortner.POST("/shorten", shortnerHandler.Shorten)
		shortner.GET("/:code", shortnerHandler.Redirect)
	}

	return r
}

func (s *Server) instantiateDependencies() {
	helloHandler = adapters.NewHelloHandler(s.db.GetDB())

	shortnerRepository := shortner.NewShortnerRepository(s.db.GetDB())
	shortnerService := shortner.NewShortnerService(shortnerRepository)
	shortnerHandler = adapters.NewShortnerHandler(shortnerService)
}
