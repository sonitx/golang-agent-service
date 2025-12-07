package server

import (
	"fmt"
	"main/handlers"
	"main/services"
	"main/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gopkg.in/yaml.v2"
)

var (
	configPath = "./configs"
)

type apiServer struct {
	baseHandler  *handlers.BaseHandler
	agentHandler *handlers.AgentHandler
}

func initConfig() (*utils.Config, error) {
	config := &utils.Config{}

	dir, err := os.ReadDir(configPath)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range dir {
		fileName := fileInfo.Name()
		if !strings.HasSuffix(fileName, ".yml") {
			continue
		}

		utils.ShowInfoLogs("Loading config: %s", fileName)
		file, err := os.Open(fmt.Sprintf("%s/%s", configPath, fileName))
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

	// service
	agentSvc := services.NewAgentService()

	// handler
	baseHandler := handlers.NewBaseHandler()
	agentHandler := handlers.NewAgentHandler(agentSvc)

	return &apiServer{
		baseHandler:  baseHandler,
		agentHandler: agentHandler,
	}
}

func (apiServer *apiServer) Start() {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	r.Use(c.Handler)

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
