package cmd

import (
	"log"
	"os"

	"github.com/MarcelArt/polygo/models"
	"github.com/MarcelArt/polygo/scaffold/fiber/scaffold"
	"github.com/MarcelArt/polygo/utils"
	"github.com/pelletier/go-toml/v2"
)

func Scaffolder(modelName string) {
	tomlBuf, err := os.ReadFile("polygo.toml")
	if err != nil {
		log.Fatalln("Failed to read polygo.toml")
	}

	var polygoTOML models.PolygoTOML
	if err := toml.Unmarshal(tomlBuf, &polygoTOML); err != nil {
		log.Fatalln("Invalid polygo.toml")
	}

	modelCamel := utils.ToCamelCase(modelName)
	modelSnake := utils.ToSeparateByCharLowered(modelCamel, '_')
	handlerRoute := utils.ToSeparateByCharLowered(modelName, '-')
	scaffold.ScaffoldModel(modelName, modelCamel, modelSnake)
	scaffold.ScaffoldRepo(modelName, modelCamel, polygoTOML.ModuleName)
	scaffold.ScaffoldHandler(modelName, handlerRoute, polygoTOML.ModuleName)
	scaffold.ScaffoldRoute(modelName, handlerRoute, polygoTOML.ModuleName)
	log.Println("Successfully scaffolded")
}
