package scaffold

import (
	"fmt"
	"os"
	"strings"

	"github.com/MarcelArt/polygo/scaffold/fiber"
)

const envTemplate = `# Generated file
PORT=${port}
`

type ENVVar struct {
	Port string
}

func CreateProjectDir(projectName string) error {
	return os.Mkdir(projectName, 0755)
}

func CreateENVFile(env ENVVar, projectPath string) error {
	envFile := strings.ReplaceAll(envTemplate, "${port}", env.Port)
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

func CreateProjectBasedOnChoice(framework string, db string, projectName string) error {
	switch framework {
	case "Fiber":
		fp := fiber.FiberProject{
			ProjectName: projectName,
			ModuleName:  projectName,
			Directory:   projectName,
		}

		return fp.GenerateFiberProject()
	}

	return nil
}
