package configenv

import (
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

type Configuration struct {
	Database SetupDatabase
}
type SetupDatabase struct {
	Hostname      string
	Port          int
	DBDriver      string
	DBUser        string
	DBPass        string
	DBName        string
	DBHost        string
	DBPort        int
	SSHUser       string
	SSHPass       string
	SSHHost       string
	SSHPort       int
	MemoryBackend string
	MemoryIndex   string
	RedisHost     string
	RedisPort     int
	RedisUsername string
	RedisPassword string
	AccountName   string
	Issuer        string
	TempFilePath  string
	ClientOrigin  string
	SmtpEmailFrom string
	SmtpHost      string
	SmtpUser      string
	SmtpPass      string
	SmtpPort      int
	GeneralAPI    string
	AiBackendHost string
}

var EnvConfigs *SetupDatabase

func InitEnvConfigs() {
	EnvConfigs = SetupConfiguration()
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func SetupConfiguration() *SetupDatabase {
	config := &SetupDatabase{
		Hostname:      getEnv("SERVER_HOSTNAME", "localhost"),
		Port:          getEnvAsInt("SERVER_PORT", 8080),
		DBDriver:      getEnv("DB_DRIVER", "postgres"),
		DBUser:        getEnv("DB_USER", "user"),
		DBPass:        getEnv("DB_PASS", "password"),
		DBName:        getEnv("DB_NAME", "exampledb"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnvAsInt("DB_PORT", 5432),
		SSHUser:       getEnv("SSH_USER", "sshuser"),
		SSHPass:       getEnv("SSH_PASS", "sshpass"),
		SSHHost:       getEnv("SSH_HOST", "localhost"),
		SSHPort:       getEnvAsInt("SSH_PORT", 22),
		MemoryBackend: getEnv("MEMORY_BACKEND", ""),
		MemoryIndex:   getEnv("MEMORY_INDEX", ""),
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisUsername: getEnv("REDIS_USERNAME", ""),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		AccountName:   getEnv("ACCOUNT_NAME", ""),
		Issuer:        getEnv("ISSUER", ""),
		TempFilePath:  getEnv("TEMP_FILE_PATH", "/tmp"),
		ClientOrigin:  getEnv("CLIENT_ORIGIN", ""),
		SmtpEmailFrom: getEnv("SMTP_EMAIL_FROM", ""),
		SmtpHost:      getEnv("SMTP_HOST", ""),
		SmtpUser:      getEnv("SMTP_USER", ""),
		SmtpPass:      getEnv("SMTP_PASS", ""),
		SmtpPort:      getEnvAsInt("SMTP_PORT", 587),
		GeneralAPI:    getEnv("GENERAL_API", ""),
		AiBackendHost: getEnv("AI_BE_ORIGIN", ""),
	}

	log.Printf("Configuration loaded: %+v", config)
	return config
}

var JWT_KEY = []byte("dasdasdasdas")

type JWTClaim struct {
	UserName string
	//UserRole int
	UserId    int
	IsVIP     bool
	UserEmail string
	jwt.RegisteredClaims
}
