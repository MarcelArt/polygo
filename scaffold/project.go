package scaffold

import (
	"fmt"
	"os"
	"strings"

	"github.com/MarcelArt/polygo/scaffold/fiber"
)

const envTemplate = `# Generated file
PORT=${port}

DB_PORT=${dbPort}
DB_USER=${dbUser}
DB_PASSWORD=${dbPass}
DB_NAME=${dbName}
DB_HOST=${dbHost}
DB_SCHEMA=${dbSchema}

JWT_SECRET=${jwtSecret}
`

type ENVVar struct {
	Port       string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBSchema   string
	JWTSecret  string
}

func CreateProjectDir(projectName string) error {
	return os.Mkdir(projectName, 0755)
}

func CreateENVFile(env ENVVar, projectPath string) error {
	envFile := strings.ReplaceAll(envTemplate, "${port}", env.Port)
	envFile = strings.ReplaceAll(envFile, "${dbPort}", env.DBPort)
	envFile = strings.ReplaceAll(envFile, "${dbUser}", env.DBUser)
	envFile = strings.ReplaceAll(envFile, "${dbPass}", env.DBPassword)
	envFile = strings.ReplaceAll(envFile, "${dbName}", env.DBName)
	envFile = strings.ReplaceAll(envFile, "${dbHost}", env.DBHost)
	envFile = strings.ReplaceAll(envFile, "${dbSchema}", env.DBSchema)
	envFile = strings.ReplaceAll(envFile, "${jwtSecret}", env.JWTSecret)

	path := fmt.Sprintf("%s/.env", projectPath)
	return os.WriteFile(path, []byte(envFile), 0644)
}

func GenerateModuleName(projectName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	modulePath := ""
	splittedPath := strings.Split(wd, "/go/src/")
	if len(splittedPath) >= 2 {
		modulePath = splittedPath[1]
	}

	return fmt.Sprintf("%s/%s", modulePath, projectName), nil

}

func CreateProjectBasedOnChoice(framework string, db string, projectName string, moduleName string) error {
	switch framework {
	case "Fiber":
		fp := fiber.FiberProject{
			ProjectName: projectName,
			ModuleName:  moduleName,
			Directory:   projectName,
		}

		return fp.GenerateFiberProject()
	}

	return nil
}
