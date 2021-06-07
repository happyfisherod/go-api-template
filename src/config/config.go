package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvStruct struct {

	// Config_file
	ConfigFile string `envconfig:"CONFIG_FILE" required:"false" default:"config.api.test"`
	ConfigType string `envconfig:"CONFIG_TYPE" required:"false" default:"yaml"`
	ConfigPath string `envconfig:"CONFIG_PATH" required:"false" default:"../envfiles"`
}

type ConfigStruct struct {

	// Versioning
	Version string
	Name    string

	// Ports
	Port        string
	HealthPort  string
	MetricsPort string

	// Prefix
	RestPrefix      string
	WebsocketPrefix string
	HealthPrefix    string
	MetricsPrefix   string

	// Monitoring
	HealthPollingInterval int
	LogLevel              string
	LogToFile             bool
	NetworkName           string

	// Kafka
	KafkaBrokerURL    string
	SchemaRegistryURL string
	KafkaGroupID      string

	// Topics
	//ConsumerTopics []string
	//ProducerTopics []string
	//SchemaNames    map[string]string

	// Kafka topics
	Topics TopicsStruct `mapstructure:"Topics"`

	// Postgres
	Postgres PostgresConfigStruct `mapstructure:"Postgres"`

	// Mongo
	Mongo MongoStruct `mapstructure:"Mongo"`
}

type TopicsStruct struct {
	ConsumerTopics []string          `json:"consumer_topics,omitempty"`
	ProducerTopics []string          `json:"producer_topics,omitempty"`
	SchemaNames    map[string]string `json:"schema_names,omitempty"`
}

type PostgresConfigStruct struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Dbname   string `json:"dbname,omitempty"`
	Sslmode  string `json:"sslmode,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type MongoStruct struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

var Vars EnvStruct
var Config ConfigStruct

// GetEnvironment Run once on main.go
func GetEnvironment() {
	//Get environment variable file
	env_file := os.Getenv("ENV_FILE")
	if env_file != "" {
		_ = godotenv.Load(env_file)
	} else {
		_ = godotenv.Load()
	}

	err := envconfig.Process("", &Vars)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}

	// Fill Config struct
	ConfigInit()
}

func ConfigInit() ConfigStruct {
	viper.SetConfigName(Vars.ConfigFile)
	viper.SetConfigType(Vars.ConfigType)
	viper.AddConfigPath(Vars.ConfigPath)

	// Set Defaults
	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v\n", err))
	}

	return Config
}

func setDefaults() {
	viper.SetDefault("Version", "v0.0.0")
	viper.SetDefault("Name", "blocks service")

	viper.SetDefault("Port", "8000")
	viper.SetDefault("HealthPort", "8080")
	viper.SetDefault("MetricsPort", "9400")

	viper.SetDefault("RestPrefix", "/rest")
	viper.SetDefault("WebsocketPrefix", "/ws")
	viper.SetDefault("HealthPrefix", "/healthcheck")
	viper.SetDefault("MetricsPrefix", "/metrics")

	viper.SetDefault("HealthPollingInterval", 10)
	viper.SetDefault("LogLevel", "INFO")
	viper.SetDefault("LogToFile", false)
	viper.SetDefault("NetworkName", "mainnet")

	viper.SetDefault("KafkaBrokerURL", "")
	viper.SetDefault("SchemaRegistryURL", "")
	viper.SetDefault("KafkaGroupID", "websocket-group")

	//viper.SetDefault("ConsumerTopics", "blocks")
	//viper.SetDefault("ProducerTopics", "blocks-ws")
	//viper.SetDefault("SchemaNames", "blocks:block_raw")

	viper.SetDefault("Topics.ConsumerTopics", "blocks")
	viper.SetDefault("Topics.ProducerTopics", "blocks-ws")
	viper.SetDefault("Topics.SchemaNames", map[string]string{"blocks": "block_raw", "blocks-ws": "block_raw"})

	viper.SetDefault("Postgres.Host", "localhost")
	viper.SetDefault("Postgres.Port", "5432")
	viper.SetDefault("Postgres.User", "postgres")
	viper.SetDefault("Postgres.Password", "changeme")
	viper.SetDefault("Postgres.Dbname", "test_db")
	viper.SetDefault("Postgres.Sslmode", "disable")
	viper.SetDefault("Postgres.Timezone", "UTC")

	viper.SetDefault("Mongo.Host", "localhost")
	viper.SetDefault("Mongo.Port", "27017")
}
