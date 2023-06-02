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

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", viper.GetString("gcp.credentials"))

	if !viper.GetBool("gcp.emulator") {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", viper.GetString("gcp.credentials"))

		return cfg, nil
	}

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

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
