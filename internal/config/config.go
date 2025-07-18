package config

import (
	"os"

	"github.com/joho/godotenv"
)


type ConfigI interface {
	GetPort() string
	GetDbHost() string
	GetDbUser() string
	GetDbPassword() string
	GetDbName() string
	GetDbPort() string
	GetDbSslMode() string
}

type config struct {

}

func (c *config) GetPort() string {
	return os.Getenv("PORT")
}

func (c *config) GetDbHost() string {
	return os.Getenv("DB_HOST")
}
func (c *config) GetDbUser() string {
	return os.Getenv("DB_USER")
}
func (c *config) GetDbPassword() string {
	return os.Getenv("DB_PASSWORD")
}
func (c *config) GetDbName() string {
	return os.Getenv("DB_NAME")
}
func (c *config) GetDbPort() string {
	return os.Getenv("DB_PORT")
}

func (c *config) GetDbSslMode () string {
	return os.Getenv("disable")
}

func NewConfigInstance(envFile string) ConfigI {
	
	if err:= godotenv.Load(envFile); err != nil {
		panic(err.Error())
	}
	
	return &config{}
}
