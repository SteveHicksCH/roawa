package config

import (
	"log"
	"os"

	"github.com/ian-kent/gofigure"
)

type config struct {
	gofigure          interface{} `order:"env,flag"`
	APIKey            string      `env:"ROA_WEBAPP_API_KEY" flag:"API key authorising use of CH API"`
	TemplateDirectory string      `env:"ROAWA_TEMPLATES_DIR" flag:"Directory for application templates"`
	ListenAddress     string      `env:"ROAWA_LISTEN_ADDRESS" flag:"Listen address and port for application web server"`
}

var cfg *config

// Get configures the application and returns the configuration
func Get() (*config, error) {

	if cfg != nil {
		return cfg, nil
	}

	log.Print("Creating a new configuration variable")
	cfg = &config{}
	// Enable multiple values in environment variables...
	err := os.Setenv("GOFIGURE_ENV_ARRAY", "1")
	if err != nil {
		// Not that Setenv can fail....
		return nil, err
	}
	err = gofigure.Gofigure(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	setDefaultValues()

	return cfg, nil
}

func setDefaultValues() {
	if cfg.TemplateDirectory == "" {
		cfg.TemplateDirectory = os.Getenv("GOPATH") + string(os.PathSeparator) + "src" + string(os.PathSeparator) + "github.com" + string(os.PathSeparator) + "shicks" + string(os.PathSeparator) + "roawa" + string(os.PathSeparator) + "http" + string(os.PathSeparator) + "templates" + string(os.PathSeparator)
	}
	if cfg.ListenAddress == "" {
		// Specify local server here to avoid MAC asking me about accepting incoming network connections (need to make this configurable anyway)
		cfg.ListenAddress = "127.0.0.1:8080"
	}
}
