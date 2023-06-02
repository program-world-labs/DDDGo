package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		App     `mapstructure:"app"`
		Swagger `mapstructure:"swagger"`
		GCP     `mapstructure:"gcp"`
		PG      `mapstructure:"postgres"`
		Redis   `mapstructure:"redis"`
		HTTP    `mapstructure:"http"`
		Log     `mapstructure:"logger"`
		Enviroment
	}

	// App -.
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	}
	// Swagger -.
	Swagger struct {
		BackendHost  string `mapstructure:"backend_host"`
		FrontendHost string `mapstructure:"frontend_host"`
	}

	// GCP -.
	GCP struct {
		Project     string `mapstructure:"project"`
		Emulator    bool   `mapstructure:"emulator"`
		Credentials string `mapstructure:"credentials"`
		Firestore   string `mapstructure:"firestore"`
		Auth        string `mapstructure:"auth"`
	}

	// PG -.
	PG struct {
		PoolMax int    `mapstructure:"pool_max"`
		URL     string `mapstructure:"url"`
	}

	// Redis -.
	Redis struct {
		DSN string `mapstructure:"dsn"`
	}

	// HTTP -.
	HTTP struct {
		Port        string `mapstructure:"port"`
		BackendPort string `mapstructure:"backend_port"`
	}

	// Log -.
	Log struct {
		Level string `mapstructure:"log_level"`
		LogID string `mapstructure:"log_id"`
	}

	Enviroment struct {
		// dev: 開發環境, prod: 正式環境, test: 測試環境
		Env string
		// backend: 後端, frontend: 前端
		Service string
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 取得環境變數
	cfg.Enviroment.Env = viper.GetString("env")
	cfg.Enviroment.Service = viper.GetString("service")
	switch viper.GetString("env") {
	case "dev":
		viper.SetConfigName("dev")
	case "prod":
		viper.SetConfigName("prod")
	case "test":
		viper.SetConfigName("test")
	default:
		viper.SetConfigName("dev")
	}
	// 取得PG環境變數
	cfg.PG.URL = viper.GetString("pg_url")
	cfg.Redis.DSN = viper.GetString("redis_url")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", viper.GetString("gcp.credentials"))

	if p := viper.GetBool("gcp.emulator"); p {
		// 設定Firestore環境變數
		if f := viper.GetString("gcp.firestore"); f != "" {
			os.Setenv("FIRESTORE_EMULATOR_HOST", f)
		}
		// 設定Pubsub環境變數
		if p := viper.GetString("gcp.pubsub.url"); p != "" {
			os.Setenv("PUBSUB_EMULATOR_HOST", p)
		}
		// 設定Storage環境變數
		if s := viper.GetString("gcp.storage.url"); s != "" {
			os.Setenv("FIREBASE_STORAGE_EMULATOR_HOST", s)
			os.Setenv("STORAGE_EMULATOR_HOST", s)
		}
		// 設定Auth環境變數
		if a := viper.GetString("gcp.auth"); a != "" {
			os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", a)
		}
	} else {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", viper.GetString("gcp.credentials"))
	}

	viper.Unmarshal(cfg)

	// err := cleanenv.ReadConfig("./config/config.yml", cfg)
	// if err != nil {
	// 	return nil, fmt.Errorf("config error: %w", err)
	// }

	// err = cleanenv.ReadEnv(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	return cfg, nil
}
