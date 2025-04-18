package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcelArt/polygo/components"
	"github.com/MarcelArt/polygo/scaffold"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var projectName string
	step1 := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &projectName,
		Placeholder: "my-golang-backend",
		Label:       "Enter project name",
	}))
	if _, err := step1.Run(); err != nil {
		log.Printf("Error creating project name: %v", err)
		os.Exit(1)
	}

	var port string
	step2 := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &port,
		Placeholder: "8080",
		Label:       "Enter preferred port",
	}))
	if _, err := step2.Run(); err != nil {
		log.Printf("Error creating port: %v", err)
		os.Exit(1)
	}

	var framework string
	step3 := tea.NewProgram(components.SingleSelect{
		Choices: []string{"Fiber", "Gin", "Echo"},
		Label:   "Select a framework",
		Value:   &framework,
	})
	if _, err := step3.Run(); err != nil {
		log.Printf("Error choosing framework: %v", err)
		os.Exit(1)
	}

	var db string
	step4 := tea.NewProgram(components.SingleSelect{
		Choices: []string{"PostgreSQL", "MySQL", "SQLite"},
		Label:   "Select a database",
		Value:   &db,
	})
	if _, err := step4.Run(); err != nil {
		log.Printf("Error choosing database: %v", err)
		os.Exit(1)
	}

	log.Println(projectName, port, framework, db)

	moduleName, err := scaffold.GenerateModuleName(projectName)
	if err != nil {
		log.Fatal(err)
	}
	step5 := tea.NewProgram(components.NewTextInput(components.TextInput{
		Value:       &moduleName,
		Placeholder: moduleName,
		Label:       fmt.Sprintf("Use this module name? (%s)", moduleName),
	}))
	if _, err := step5.Run(); err != nil {
		log.Printf("Error naming module: %v", err)
		os.Exit(1)
	}

	scaffold.CreateProjectDir(projectName)
	scaffold.CreateENVFile(
		scaffold.ENVVar{
			Port: port,
		},
		projectName,
	)

	if err := scaffold.CreateProjectBasedOnChoice(framework, db, projectName); err != nil {
		log.Fatal(err)
	}
}
