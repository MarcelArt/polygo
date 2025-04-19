package fiber

import (
	"fmt"
	"os"
	"strings"
)

const routeFileTemplate = `// This file is auto generated by polygo
package routes

import (
	"log"
	"os"
	"time"

	"${moduleName}/config"
	"${moduleName}/database"
	"${moduleName}/middlewares"
	"${moduleName}/repositories"
	api_routes "${moduleName}/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	file, err := os.OpenFile("./model-craft.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "[${time}] ${status} - ${method} ${path} - Query: ${queryParams} - Request: ${body} - Response: ${resBody}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} - Query: ${queryParams} - Request: ${body} - Response: ${resBody}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	if config.Env.ServerENV != "prod" {
		app.Get("/swagger/*", swagger.HandlerDefault)     // default
		app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			URL:         "http://example.com/doc.json",
			DeepLinking: false,
		}))
	}	

	app.Get("/metrics", monitor.New())

	authMiddleware := middlewares.NewAuthMiddleware(repositories.NewUserRepo(database.GetDB()))

	api := app.Group("/api")
	api_routes.SetupUserRoutes(api, authMiddleware)
	api_routes.SetupAuthorizedDeviceRoutes(api, authMiddleware)
}
`

const apiUserRouteFileTemplate = `// This file is auto generated by polygo
package api_routes

import (
	"${moduleName}/database"
	api_handlers "${moduleName}/handlers/api"
	"${moduleName}/middlewares"
	"${moduleName}/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.NewUserHandler(
		repositories.NewUserRepo(database.GetDB()),
		repositories.NewAuthorizedDeviceRepo(database.GetDB()),
	)

	g := api.Group("/user")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)

	g.Post("/", h.Create)
	g.Post("/login", h.Login)
	g.Post("/refresh", h.Refresh)

	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}
`

const apiAuthDeviceRouteFileTemplate = `// This file is auto generated by polygo
package api_routes

import (
	"${moduleName}/database"
	api_handlers "${moduleName}/handlers/api"
	"${moduleName}/middlewares"
	"${moduleName}/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthorizedDeviceRoutes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.NewAuthorizedDeviceHandler(repositories.NewAuthorizedDeviceRepo(database.GetDB()))

	g := api.Group("/authorized-device")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)
	g.Post("/", auth.ProtectedAPI, h.Create)
	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}
`

func (fp FiberProject) createRouteFiles() error {
	routeFileBody := strings.ReplaceAll(routeFileTemplate, "${moduleName}", fp.ModuleName)
	apiUserRouteFileBody := strings.ReplaceAll(apiUserRouteFileTemplate, "${moduleName}", fp.ModuleName)
	apiAuthDeviceRouteFileBody := strings.ReplaceAll(apiAuthDeviceRouteFileTemplate, "${moduleName}", fp.ModuleName)

	if err := os.Mkdir(fmt.Sprintf("%s/routes", fp.Directory), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(fmt.Sprintf("%s/routes/api", fp.Directory), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("%s/routes/route.go", fp.Directory), []byte(routeFileBody), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(fmt.Sprintf("%s/routes/api/user.route.go", fp.Directory), []byte(apiUserRouteFileBody), 0644); err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/routes/api/authorized_device.route.go", fp.Directory), []byte(apiAuthDeviceRouteFileBody), 0644)
}
