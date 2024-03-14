package server

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/highlight/highlight/sdk/highlight-go"
	hlog "github.com/highlight/highlight/sdk/highlight-go/log"

	"github.com/noraincode/fiber-seed/internal/app/config"
	middlewares "github.com/noraincode/fiber-seed/internal/app/middlewares/response"
	"github.com/noraincode/fiber-seed/internal/app/server/routes"
)

type Server struct {
	Config *config.Config
	App    *fiber.App
}

func setUpMiddlewares(app *fiber.App) {
	// app.Use(highlightFiber.Middleware())
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())
	app.Use(limiter.New(limiter.Config{
		Max: 20,
	}))
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(requestid.New(requestid.Config{
		Header:     "X-Trace-Id",
		ContextKey: "trace_id",
	}))
	app.Use(middlewares.UnifiedResponse())
}

func setUpHighlight() {
	highlight.SetProjectID("ldw7r95g")
	highlight.Start(
		highlight.WithServiceName("my-app"),
		highlight.WithServiceVersion("git-sha"),
	)
	defer highlight.Stop()

	hlog.Init()
}

func NewServer(r routes.IRouter, cfg *config.Config) (*Server, error) {
	setUpHighlight()

	app := fiber.New(fiber.Config{
		AppName:     "letschedule",
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	setUpMiddlewares(app)

	r.Register(app)

	srv := &Server{
		Config: cfg,
		App:    app,
	}

	return srv, nil
}

func (srv *Server) Run() error {
	log.Info("Starting the server...")

	go func() {
		if err := srv.App.Listen(":" + srv.Config.App.Port); err != nil && err != http.ErrServerClosed {
			log.Errorf("Failed to start the server with error", err)
		}
	}()

	return nil
}

func (srv *Server) Shutdown() {
	if err := srv.App.Shutdown(); err != nil {
		log.Errorf("Shutdown the server with error", err)
	}
	log.Info("The server was shutdown normally, bye bye.")
}
