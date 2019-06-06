package router
import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	health_controller "github.com/TameshkCloud/todoberry-server/application/controller/health"
)

//RunGin run gin
func RunGin() {

	// init gin
	r := gin.Default()

	// setup routes

	// testing health APIs!
	healthRoutes := r.Group("/health")
	{
		healthRoutes.GET("/welcome", health_controller.SayHello)	
		healthRoutes.GET("/check/redis", health_controller.CheckRedis)	
	}
	

	// use ginSwagger middleware
	// get more information on api:
	// https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
	// or check HelloController.go
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// gin middleware config
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://golab.rezaseyf.ir:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Run(":3000")
	fmt.Println("server is running in 127.0.0.1:3000")
}