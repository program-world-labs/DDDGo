package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		App          `mapstructure:"app"`
		Swagger      `mapstructure:"swagger"`
		SQL          `mapstructure:"sql"`
		NoSQL        `mapstructure:"nosql"`
		Redis        `mapstructure:"redis"`
		Storage      `mapstructure:"storage"`
		HTTP         `mapstructure:"http"`
		Log          `mapstructure:"logger"`
		Kafka        `mapstructure:"kafka"`
		Telemetry    `mapstructure:"telemetry"`
		EventStoreDB `mapstructure:"esdb"`
		Env
	}

	// App -.
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	}
	// Swagger -.
	Swagger struct {
		Host string `mapstructure:"host"`
	}

	// SQL -.
	SQL struct {
		PoolMax      int           `mapstructure:"pool_max"`
		Host         string        `mapstructure:"host"`
		Port         int           `mapstructure:"port"`
		User         string        `mapstructure:"user"`
		Password     string        `mapstructure:"password"`
		DB           string        `mapstructure:"db"`
		ConnAttempts int           `mapstructure:"connection_attempts"`
		ConnTimeout  time.Duration `mapstructure:"connection_timeout"`
		Type         string        `mapstructure:"type" validate:"oneof=postgres"`
	}

	NoSQL struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DB       string `mapstructure:"db"`
		Type     string `mapstructure:"type" validate:"oneof=firestore"`
	}

	// Redis -.
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}

	// Storage -.
	Storage struct {
		Host   string `mapstructure:"host"`
		Bucket string `mapstructure:"bucket"`
		Type   string `mapstructure:"type" validate:"oneof=gcp"`
	}

	// HTTP -.
	HTTP struct {
		Port string `mapstructure:"port"`
	}

	// Log -.
	Log struct {
		Project string `mapstructure:"project"`
		Level   string `mapstructure:"log_level"`
		LogID   string `mapstructure:"log_id"`
	}

	// Telemetry -.
	Telemetry struct {
		Host       string  `mapstructure:"host"`
		Port       int     `mapstructure:"port"`
		Batcher    string  `mapstructure:"batcher" validate:"oneof=gcp"`
		SampleRate float64 `mapstructure:"sample_rate"`
		Enabled    bool    `mapstructure:"enabled"`
	}

	// Kafka -.
	Kafka struct {
		Brokers []string `mapstructure:"brokers"`
		GroupID string   `mapstructure:"group_id"`
	}

	// EventStoreDB -.
	EventStoreDB struct {
		Host string `mapstructure:"host"`
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

	setViperEnv()
	setConfigValues(cfg)

	err := readAndParseConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = applyEnvSetting(cfg)
	if err != nil {
		return nil, err
	}

	setGCPEnv(cfg)

	return cfg, nil
}

func setViperEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func setConfigValues(cfg *Config) {
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
}

func readAndParseConfig(cfg *Config) error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	return nil
}

func applyEnvSetting(cfg *Config) error {
	const postgresType = "postgres"

	const mysqlType = "mysql"

	sqlDSN := viper.GetString("sql_url")

	if sqlDSN != "" {
		// postgres://user:password@host:port/db 解析成 host:port:username:password:db
		// mysql://user:password@host:port/db 解析成 host:port:username:password:db
		u, err := url.Parse(sqlDSN)
		if err != nil {
			return err
		}

		cfg.SQL.Host = u.Hostname()
		cfg.SQL.Port, err = strconv.Atoi(u.Port())

		if err != nil {
			return err
		}

		switch strings.Split(sqlDSN, ":")[0] {
		case postgresType:
			cfg.SQL.Type = "postgresql"
		case mysqlType:
			cfg.SQL.Type = "mysql"
		}

		cfg.SQL.User = u.User.Username()
		cfg.SQL.Password, _ = u.User.Password()
		cfg.SQL.DB = strings.TrimLeft(u.Path, "/")
	}

	redisDSN := viper.GetString("redis_url")

	if redisDSN != "" {
		// redis://user:password@host:port/db 解析成 host:port:username:password:db
		u, err := url.Parse(redisDSN)
		if err != nil {
			return err
		}

		cfg.Redis.Host = u.Hostname()
		cfg.Redis.Port, err = strconv.Atoi(u.Port())

		if err != nil {
			return err
		}

		cfg.Redis.Password, _ = u.User.Password()
		cfg.Redis.DB, err = strconv.Atoi(strings.TrimLeft(u.Path, "/"))

		if err != nil {
			return err
		}
	}

	kafkaDSN := viper.GetString("kafka_url")

	if kafkaDSN != "" {
		cfg.Kafka.Brokers = []string{kafkaDSN}
	}

	// Log

	logProject := viper.GetString("log_project")
	if logProject != "" {
		cfg.Log.Project = logProject
	}

	return nil
}

func setGCPEnv(cfg *Config) {
	// Read GCP credentials from env
	if viper.GetString("gcp_credentials") != "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", viper.GetString("gcp_credentials"))
	}

	// Set GCP emulator env
	if cfg.Storage.Type == "gcp" {
		if cfg.Storage.Host != "" {
			os.Setenv("FIREBASE_STORAGE_EMULATOR_HOST", cfg.Storage.Host)
			os.Setenv("STORAGE_EMULATOR_HOST", cfg.Storage.Host)
		}
	}
}
