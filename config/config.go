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
		Env
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
		Monitor     bool   `mapstructure:"monitor"`
		Emulator    bool   `mapstructure:"emulator"`
		Credentials string `mapstructure:"credentials"`
		Firestore   string `mapstructure:"firestore"`
		Storage     struct {
			Bucket string `mapstructure:"bucket"`
			URL    string `mapstructure:"url"`
		} `mapstructure:"storage"`
		Auth string `mapstructure:"auth"`
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

	Env struct {
		// dev: 開發環境, prod: 正式環境, test: 測試環境
		EnvName string
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
	cfg.Env.EnvName = viper.GetString("env")
	cfg.Env.Service = viper.GetString("service")

	envConfig := map[string]string{
		"dev":  "dev",
		"prod": "prod",
		"test": "test",
	}
	configName, ok := envConfig[viper.GetString("env")]

	if !ok {
		configName = "dev"
	}

	viper.SetConfigName(configName)

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

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// 設定GCP環境變數
	if cfg.GCP.Credentials != "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cfg.GCP.Credentials)
	}

	if !cfg.GCP.Emulator {
		return cfg, nil
	}

	// 設定Firestore環境變數
	if cfg.GCP.Firestore != "" {
		os.Setenv("FIRESTORE_EMULATOR_HOST", cfg.GCP.Firestore)
	}

	// 設定Storage環境變數
	if cfg.GCP.Storage.URL != "" {
		os.Setenv("FIREBASE_STORAGE_EMULATOR_HOST", cfg.GCP.Storage.URL)
		os.Setenv("STORAGE_EMULATOR_HOST", cfg.GCP.Storage.URL)
	}
	// 設定Auth環境變數
	if cfg.GCP.Auth != "" {
		os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", cfg.GCP.Auth)
	}

	return cfg, nil
}
