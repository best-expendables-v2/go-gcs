package gogcs

import (
	"github.com/best-expendables-v2/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type GCSConfig struct {
	Bucket    string `envconfig:"GCS_BUCKET" required:"true"`
	ProjectID string `envconfig:"GCS_PROJECT_ID" required:"true"`
	BaseUrl   string `envconfig:"GCS_BASE_URL" required:"true"`
}

func LoadGSCConfig() GCSConfig {
	var config GCSConfig
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}
	envconfig.MustProcess("", &config)
	return config
}
