package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbDriver   string
}

type TokenConfig struct {
	TokenIssue    string `json:"TokenIssue"`
	TokenSecret   []byte `json:"TokenSecret"`
	TokenExpire   time.Duration
	SigningMethod *jwt.SigningMethodHMAC
}

type ApiConfig struct {
	ApiPort string
}

type SmtpConfig struct {
	EmailName    string
	EmailAppPswd string
}

type ClientConfig struct {
	ResetPasswordURL          url.URL
	ResetPasswordHTMLTemplate string
}

type Config struct {
	DbConfig
	TokenConfig
	ApiConfig
	SmtpConfig
	ClientConfig
	Env string
}

func (c *Config) Configuration() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	c.DbConfig = DbConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbDriver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))

	c.TokenConfig = TokenConfig{
		TokenIssue:    os.Getenv("TOKEN_ISSUE"),
		TokenSecret:   []byte(os.Getenv("TOKEN_SECRET")),
		TokenExpire:   time.Hour * time.Duration(tokenExpire),
		SigningMethod: jwt.SigningMethodHS256,
	}

	c.SmtpConfig = SmtpConfig{
		EmailName:    os.Getenv("EMAIL_NAME"),
		EmailAppPswd: os.Getenv("EMAIL_APP_PASSWORD"),
	}

	c.Env = os.Getenv("ENV")
	if c.Env != "development" && c.Env != "staging" && c.Env != "production" {
		return fmt.Errorf("invalid env: %v", c.Env)
	}

	rpURL, err := url.Parse(os.Getenv("RESET_PASSWORD_URL"))
	if err != nil {
		return fmt.Errorf("error parsing reset password url: %v, env may be missing", err)
	}

	c.ClientConfig = ClientConfig{
		ResetPasswordURL:          *rpURL,
		ResetPasswordHTMLTemplate: os.Getenv("RESET_PASSWORD_HTML_TEMPLATE"),
	}
	if c.DbHost == "" || c.DbPort == "" || c.DbUser == "" || c.DbPassword == "" || c.DbName == "" || c.DbDriver == "" || c.ApiPort == "" || c.TokenIssue == "" || len(c.TokenSecret) == 0 || c.TokenExpire < 0 || c.SigningMethod == nil || c.Env == "" || c.EmailName == "" || c.EmailAppPswd == "" || c.ResetPasswordHTMLTemplate == "" {
		return fmt.Errorf("missing environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := config.Configuration(); err != nil {
		return nil, err
	}

	return config, nil
}
