package main

import (
	"github.com/TameshkCloud/todoberry-server/config"
	"github.com/TameshkCloud/todoberry-server/router"
	"github.com/TameshkCloud/todoberry-server/services"
	_ "github.com/TameshkCloud/todoberry-server/docs" // docs generated using swagger cli
)

// start gin server and load application
// -------------------------------------
// Swagger Documentation annotations:
// 
// 
// 
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	// load config
	config.Init()

	// services
	services.Init()

	// start gin server
	router.RunGin()
}