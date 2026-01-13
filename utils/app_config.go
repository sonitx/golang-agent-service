package utils

var AppConfig *Config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	ModelConfig    ModelConfig  `yaml:"models"`
	AgentConfig    []AgentItem  `yaml:"agents"`
	PostgresConfig PostgresConf `yaml:"postgres"`
	AgenticNodes   AgenticNodes `yaml:"agentic-nodes"`
}

type ModelConfig struct {
	Gemini struct {
		Enable    bool   `yaml:"enable"`
		ModelName string `yaml:"model-name"`
		APIKey    string `yaml:"apikey"`
	} `yaml:"gemini"`
	Ollama struct {
		Enable        bool   `yaml:"enable"`
		ModelName     string `yaml:"model-name"`
		ServerAddress string `yaml:"server-address"`
	} `yaml:"ollama"`
}

type AgentItem struct {
	Enable      bool   `yaml:"enable"`
	Key         string `yaml:"key"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type PostgresConf struct {
	Enable      bool   `yaml:"enable"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	SslMode     string `yaml:"sslmode"`
	MaxIdleConn int    `yaml:"max-idle-conn"`
	MaxOpenConn int    `yaml:"max-open-conn"`
}

type AgenticNodes struct {
	Router AgenticModel `yaml:"router"`
	Direct AgenticModel `yaml:"direct"`
	Logic  AgenticModel `yaml:"logic"`
	RAG    AgenticModel `yaml:"rag"`
}

type AgenticModel struct {
	ModelType string `yaml:"model-type"`
	ModelName string `yaml:"model-name"`
	APIKey    string `yaml:"api-key"`
}
