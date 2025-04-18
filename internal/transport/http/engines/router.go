package engines

import (
	"github.com/Square-POC/SquarePosBE/internal/controllers"
	"github.com/Square-POC/SquarePosBE/internal/transport/http/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Engine struct {
	controller *controllers.ControllerV1
}

func NewEngine(
	controller *controllers.ControllerV1,
) *Engine {
	return &Engine{
		controller: controller,
	}
}

func (e *Engine) GetEngine() *gin.Engine {
	engine := gin.New()

	pprof.Register(engine)

	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	engine.GET("/auth", e.controller.GoogleLogin)
	engine.GET("/callback", e.controller.GoogleCallback)

	engine.GET("/auth/google", e.controller.GoogleLogin)
	engine.GET("/auth/google/callback", e.controller.GoogleCallback)

	v1Group := engine.Group("/api/v1")
	v1Group.Use(middlewares.AuthMiddleware())
	{
		v1Group.POST("/earn", e.controller.AccumulateLoyaltyController)
	}

	return engine
}
