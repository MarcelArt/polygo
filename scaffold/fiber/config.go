package fiber

import (
	"fmt"
	"os"
)

const envFileTemplate = `// This file is auto generated by polygo
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	PORT         string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBHost       string
	JwtSecret    string
	ServerENV    string
}

var Env *env

func SetupENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	Env = &env{
		PORT:         os.Getenv("PORT"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBHost:       os.Getenv("DB_HOST"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		ServerENV:    os.Getenv("SERVER_ENV"),
	}
}
`

func (fp FiberProject) createENVGoFile() error {
	if err := os.Mkdir(fmt.Sprintf("%s/internal/config", fp.Directory), 0755); err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/internal/config/env.go", fp.Directory), []byte(envFileTemplate), 0644)
}
