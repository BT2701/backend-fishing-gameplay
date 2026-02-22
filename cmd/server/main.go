package main

import (
	"log"

	"github.com/BT2701/backend-fishing-gameplay/adapter/repository/mongo"
	"github.com/BT2701/backend-fishing-gameplay/adapter/repository/redis"
	"github.com/BT2701/backend-fishing-gameplay/internal/delivery/http"
	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/config"
	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/logger"
	infmongo "github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence/mongo"
	infredis "github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/persistence/redis"
	"github.com/BT2701/backend-fishing-gameplay/internal/infrastructure/server"
	"github.com/BT2701/backend-fishing-gameplay/internal/usecase"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Close()

	zapLogger := logger.Get()

	// Load configuration
	cfg := config.Load()

	// Connect to MongoDB with retry
	mongoClient, err := infmongo.ConnectWithRetryZap(
		cfg.Mongo.URI,
		cfg.Mongo.Database,
		cfg.Mongo.Timeout,
		cfg.Mongo.MaxRetries,
		cfg.Mongo.RetryDelay,
		zapLogger,
	)
	if err != nil {
		zapLogger.Fatal("Failed to connect to MongoDB after retries", zap.Error(err))
	}
	defer infmongo.Close(mongoClient)

	mongoDB := mongoClient.Database(cfg.Mongo.Database)

	// Connect to Redis with retry
	redisClient, err := infredis.ConnectWithRetryZap(
		cfg.Redis.Addr,
		cfg.Redis.Password,
		cfg.Redis.DB,
		cfg.Redis.MaxRetries,
		cfg.Redis.RetryDelay,
		zapLogger,
	)
	if err != nil {
		zapLogger.Fatal("Failed to connect to Redis after retries", zap.Error(err))
	}

	// Initialize repositories
	roomRepo := mongo.NewRoomRepository(mongoDB)
	playerRepo := mongo.NewPlayerRepository(mongoDB)
	fishRepo := mongo.NewFishRepository(mongoDB)
	gunRepo := mongo.NewGunRepository(mongoDB)
	rtpRepo := redis.NewRTPRepository(redisClient)
	gameConfigMongoRepo := mongo.NewGameConfigRepository(mongoDB)

	// Initialize cache repositories (with fallback to MongoDB)
	gameConfigRepo := redis.NewGameConfigCacheRepository(redisClient, gameConfigMongoRepo, cfg.Redis.CacheTTL)

	// Initialize usecases
	roomUsecase := usecase.NewRoomUsecase(roomRepo, playerRepo)
	fishUsecase := usecase.NewFishUsecase(roomRepo, fishRepo)
	shootUsecase := usecase.NewShootUsecase(roomRepo, playerRepo, fishRepo, gunRepo, rtpRepo)
	rtpUsecase := usecase.NewRTPUsecase(rtpRepo)
	skillUsecase := usecase.NewSkillUsecase(playerRepo)
	gameConfigUsecase := usecase.NewGameConfigUsecase(gameConfigRepo)

	// Initialize HTTP server
	srv := server.New(cfg.Server.Host, cfg.Server.Port, zapLogger)

	// Setup routes
	http.SetupRoutes(srv.GetApp(), roomUsecase, fishUsecase, shootUsecase, rtpUsecase, skillUsecase, gameConfigUsecase)

	// Start server
	if err := srv.Start(); err != nil {
		zapLogger.Fatal("Failed to start server", zap.Error(err))
	}
}
