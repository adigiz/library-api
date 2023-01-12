package config

import "os"

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type AppConfig struct {
	ENV                string
	AppName            string
	JWTSecretKey       []byte
	JWTExpiryInMinutes int64
	DBConfig           dbConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	ENV:                getENV("ENV", "testing"),
	AppName:            "sea-labs-library",
	JWTSecretKey:       []byte("very-secret"),
	JWTExpiryInMinutes: 15,
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "P@ssw0rd"),
		DBName:   getENV("DB_NAME", "library"),
		Port:     getENV("DB_PORT", "5432"),
	},
}
