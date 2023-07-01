package config

import (
	// "log"
	"os"
	"strconv"

	// "github.com/joho/godotenv"
)

type Config struct {
	AppName              string
	MailName             string
	MailAdminId          string
	MailPassword         string
	AppEnv               string
	DBUserName           string
	DBHostReader         string
	DBHostWriter         string
	DBPort               string
	DBPassword           string
	DBName               string
	DBMaxIdleConnections int
	DBMaxOpenConnections int
	ServerPort           string
	RedisAddress         string
	RedisPassword        string
	CacheEnabled         bool
	AccessSecret         string
	RefreshSecret        string

	CustomTemplateID          string
	MailVerifictionTemplateID string
	PasswordResetTemplateID   string
	SendGridApiKey            string

	DiscordWebHookId             string
	DiscordWebHookToken          string
	DiscordShouldWaitForResponse string
	DiscordUrl                   string
}

var config Config

// init() Should run at the very beginning, before any other package or code.
func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	// configFilePath := "./config/.env.dev"

	// switch appEnv {
	// case "production":
	// 	configFilePath = "./config/.env.prod"
	// 	break
	// case "stage":
	// 	configFilePath = "./config/.env.stage"
	// 	break
	// }
	// log.Println("reading env from: " + configFilePath)

	// e := godotenv.Load(configFilePath)

	// if e != nil {
	// 	log.Println("error loading .env: ", e)
	// 	panic(e.Error())
	// }

	config.AppName = os.Getenv("SERVICE_NAME")
	config.MailName = os.Getenv("MAIL_NAME")
	config.MailAdminId = os.Getenv("MAIL_ADMIN_ID")
	config.MailPassword = os.Getenv("MAIL_PASSWORD")
	config.AppEnv = appEnv
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.DBMaxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONENCTION"))
	config.DBMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	config.CacheEnabled, _ = strconv.ParseBool(os.Getenv("CACHE_ENABLED"))
	config.AccessSecret = os.Getenv("ACCESS_SECRET")
	config.RefreshSecret = os.Getenv("REFRESH_SECRET")
	config.MailVerifictionTemplateID = os.Getenv("MAIL_VERIFICATION_TEMPLATE_ID")
	config.PasswordResetTemplateID = os.Getenv("PASSWORD_RESET_TEMPLATE_ID")
	config.CustomTemplateID = os.Getenv("CUSTOM_TEMPLATE_ID")
	config.SendGridApiKey = os.Getenv("SENDGRID_API_KEY")
	config.DiscordWebHookId = os.Getenv("DISCORD_WEBHOOK_ID")
	config.DiscordWebHookToken = os.Getenv("DISCORD_WEBHOOK_TOKEN")
	config.DiscordShouldWaitForResponse = os.Getenv("DISCORD_SHOULD_WAIT_FOR_RESPONSE")
	config.DiscordUrl = os.Getenv("DISCORD_URL")
}

func Get() Config {
	return config
}

func IsProduction() bool {
	return config.AppEnv == "production"
}
