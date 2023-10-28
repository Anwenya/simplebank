package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT                string        `mapstructure:"ENVIRONMENT"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	MigrationURL               string        `mapstructure:"MIGRATION_URL"`
	RedisAddress               string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress          string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailGmailSenderName       string        `mapstructure:"EMAIL_GMAIL_SENDER_NAME"`
	EmailGmailSenderAddress    string        `mapstructure:"EMAIL_GMAIL_SENDER_ADDRESS"`
	EmailGmailSenderPassword   string        `mapstructure:"EMAIL_GMAIL_SENDER_PASSWORD"`
	EmailTencentSenderName     string        `mapstructure:"EMAIL_TENCENT_SENDER_NAME"`
	EmailTencentSenderAddress  string        `mapstructure:"EMAIL_TENCENT_SENDER_ADDRESS"`
	EmailTencentSenderPassword string        `mapstructure:"EMAIL_TENCENT_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// 配置文件名称
	viper.SetConfigName("app")
	// 配置文件格式
	viper.SetConfigType("env")
	// 环境变量中的值会覆盖配置文件中的同名值
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
