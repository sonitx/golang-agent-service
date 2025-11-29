package server

import (
	"fmt"
	"main/handlers"
	"main/utils"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/yaml.v2"
)

var (
	configPath = "./configs"
)

type apiServer struct {
	baseHandler *handlers.BaseHandler
}

func initConfig() (*utils.Config, error) {
	config := &utils.Config{}

	dir, err := os.ReadDir(configPath)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range dir {
		file, err := os.Open(fmt.Sprintf("%s/%s", configPath, fileInfo.Name()))
		if err != nil {
			return nil, err
		}
		defer file.Close()

		d := yaml.NewDecoder(file)

		if err := d.Decode(&config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func Initialize() *apiServer {
	var err error

	// config
	utils.AppConfig, err = initConfig()
	if err != nil {
		panic(err)
	}

	// handler
	baseHandler := handlers.NewBaseHandler()

	return &apiServer{
		baseHandler: baseHandler,
	}
}

func (apiServer *apiServer) Start() {
	r := chi.NewRouter()

	// Init global middleware
	r.Use(middleware.RequestID)                 // Assign request ID to context
	r.Use(middleware.Logger)                    // Log requests
	r.Use(middleware.Recoverer)                 // Recover from panics
	r.Use(middleware.Timeout(60 * time.Second)) // Set timeout

	// Init routes
	apiServer.initApiRoutes(r)

	// Start server
	utils.ShowInfoLogs("🚀 Server running on :%d", utils.AppConfig.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", utils.AppConfig.Server.Port), r)
}
