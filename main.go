package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcelArt/polygo/components"
	"github.com/MarcelArt/polygo/scaffold"
	"github.com/MarcelArt/polygo/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var projectName string
	projectNameStep := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &projectName,
		Placeholder: "my-golang-backend",
		Label:       "Enter project name",
	}))
	if _, err := projectNameStep.Run(); err != nil {
		log.Printf("Error creating project name: %v", err)
		os.Exit(1)
	}
	fmt.Println("")

	var port string
	portStep := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &port,
		Placeholder: "8080",
		Label:       "Enter preferred port",
	}))
	if _, err := portStep.Run(); err != nil {
		log.Printf("Error creating port: %v", err)
		os.Exit(1)
	}
	fmt.Println("")

	var framework string
	frameworkStep := tea.NewProgram(components.SingleSelect{
		Choices: []string{"Fiber", "Gin", "Echo"},
		Label:   "Select a framework",
		Value:   &framework,
	})
	if _, err := frameworkStep.Run(); err != nil {
		log.Printf("Error choosing framework: %v", err)
		os.Exit(1)
	}
	fmt.Println("")

	var db string
	dbStep := tea.NewProgram(components.SingleSelect{
		Choices: []string{"PostgreSQL", "MySQL", "SQLite"},
		Label:   "Select a database",
		Value:   &db,
	})
	if _, err := dbStep.Run(); err != nil {
		log.Printf("Error choosing database: %v", err)
		os.Exit(1)
	}
	fmt.Println("")

	var dbPort string
	var dbUser string
	var dbPass string
	var dbName string
	var dbHost string
	var dbSchema string
	if db == "PostgreSQL" {
		dbPortStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbPort,
			Placeholder: "5432",
			Label:       "Enter PostgreSQL port",
		}))
		if _, err := dbPortStep.Run(); err != nil {
			log.Printf("Error input DB Port: %v", err)
			os.Exit(1)
		}
		fmt.Println("")

		dbUserStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbUser,
			Placeholder: "postgres",
			Label:       "Enter PostgreSQL username",
		}))
		if _, err := dbUserStep.Run(); err != nil {
			log.Printf("Error input DB username: %v", err)
			os.Exit(1)
		}
		fmt.Println("")

		dbPassStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbPass,
			Placeholder: "postgres",
			Label:       "Enter PostgreSQL password",
			Type:        "password",
		}))
		if _, err := dbPassStep.Run(); err != nil {
			log.Printf("Error input DB password: %v", err)
			os.Exit(1)
		}
		fmt.Println("")

		dbNameStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbName,
			Placeholder: projectName,
			Label:       "Enter PostgreSQL database name",
		}))
		if _, err := dbNameStep.Run(); err != nil {
			log.Printf("Error input DB name: %v", err)
			os.Exit(1)
		}
		fmt.Println("")

		dbHostStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbHost,
			Placeholder: "localhost",
			Label:       "Enter PostgreSQL host",
		}))
		if _, err := dbHostStep.Run(); err != nil {
			log.Printf("Error input DB host: %v", err)
			os.Exit(1)
		}
		fmt.Println("")

		dbSchemaStep := tea.NewProgram(components.NewTextInput(components.TextInput{
			Value:       &dbSchema,
			Placeholder: "public",
			Label:       "Enter PostgreSQL DB schema",
		}))
		if _, err := dbSchemaStep.Run(); err != nil {
			log.Printf("Error input DB schema: %v", err)
			os.Exit(1)
		}
		fmt.Println("")
	} else {
		fmt.Printf("\n%s is not supported yet.", db)
		os.Exit(0)
	}

	moduleName, err := scaffold.GenerateModuleName(projectName)
	if err != nil {
		log.Fatal(err)
	}
	moduleStep := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &moduleName,
		Placeholder: moduleName,
		Label:       fmt.Sprintf("Use this module name? (%s)", moduleName),
	}))
	if _, err := moduleStep.Run(); err != nil {
		log.Printf("Error naming module: %v", err)
		os.Exit(1)
	}
	fmt.Println("")

	scaffold.CreateProjectDir(projectName)
	scaffold.CreateENVFile(
		scaffold.ENVVar{
			Port:       port,
			DBPort:     dbPort,
			DBUser:     dbUser,
			DBPassword: dbPass,
			DBName:     dbName,
			DBHost:     dbHost,
			DBSchema:   dbSchema,
			JWTSecret:  utils.RandString(32),
		},
		projectName,
	)

	if err := scaffold.CreateProjectBasedOnChoice(framework, db, projectName, moduleName); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project %s created successfully\nRun:\n", projectName)
	fmt.Printf("1. cd %s\n", projectName)
	fmt.Printf("2. go mod init\n")
	fmt.Printf("3. go mod tidy\n")
	fmt.Printf("4. Uncomment `// _ %s/docs` in main.go file\n", moduleName)
	fmt.Printf("5. Run: swag init --parseInternal --parseDependency\n")
	fmt.Printf("6. go run main.go\n")
	fmt.Println("To get started")
}
