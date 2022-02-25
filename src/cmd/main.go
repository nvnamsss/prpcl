package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nvnamsss/prpcl/adapters/database"
	_ "github.com/nvnamsss/prpcl/cmd/docs"
	"github.com/nvnamsss/prpcl/configs"
	"github.com/nvnamsss/prpcl/controllers"
	"github.com/nvnamsss/prpcl/logger"
	"github.com/nvnamsss/prpcl/repositories"
	"github.com/nvnamsss/prpcl/services"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

// @title prpcl
// @version 1.0
// @description prpcl API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name Nam Nguyen
// @contact.email nvnam.c@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /prpcl/v1
func main() {
	db := database.NewDatabase()
	if err := db.Open(configs.Config.MySQL.ConnectionString(), gorm.Config{}); err != nil {
		logger.Fatalf(err, "Creating connection to DB: %v", err)
	}

	var r = gin.Default()
	corsConfig := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "TimezoneOffset"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})

	var (
		wagerRepository            = repositories.NewWagerRepository(db)
		wagerTransactionRepository = repositories.NewWagerTransactionRepository(db)
	)

	var (
		wagerService = services.NewWagerService(wagerRepository, wagerTransactionRepository, db)
	)

	var (
		wagerController = controllers.NewWagerController(wagerService)
	)

	r.Use(
		corsConfig,
	)

	wager := r.Group("/")
	{
		wager.POST("/wagers", wagerController.Place)
		wager.POST("/buy/:id", wagerController.Buy)
		wager.GET("/wagers", wagerController.List)
	}

	if configs.Config.RunMode == gin.DebugMode && configs.Config.Env != "PRODUCTION" {
		r.GET("/prpcl/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &http.Server{
		Addr:    configs.Config.AddressListener(),
		Handler: r,
	}

	go func() {
		logger.Infof("Starting Server on %v", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf(err, "Opening HTTP server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logger.Infof("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Shutdown error: %v", err)
	}

	os.Exit(0)
}

func init() {
	if _, err := configs.New(); err != nil {
		os.Exit(1)
	}
}
