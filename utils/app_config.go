package utils

var AppConfig *Config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	ModelConfig ModelConfig `yaml:"models"`
	AgentConfig []AgentItem `yaml:"agents"`
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
