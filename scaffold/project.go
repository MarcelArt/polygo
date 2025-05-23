package scaffold

import (
	"fmt"
	"os"
	"strings"

	"github.com/MarcelArt/polygo/models"
	"github.com/MarcelArt/polygo/scaffold/fiber"
	"github.com/pelletier/go-toml/v2"
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

const gitIgnoreTemplate = `# Generated file
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

docs/*

.env

tmp/*

*.log
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

func CreateGitIgnoreFile(projectPath string) error {
	path := fmt.Sprintf("%s/.gitignore", projectPath)
	return os.WriteFile(path, []byte(gitIgnoreTemplate), 0644)
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

func CreatePolygoTOML(moduleName string, projectPath string) error {
	polygoTOML := models.PolygoTOML{
		ModuleName: moduleName,
	}

	buf, err := toml.Marshal(polygoTOML)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/polygo.toml", projectPath)
	return os.WriteFile(path, buf, 0644)
}
