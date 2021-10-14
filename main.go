package main

import (
	"flag"
	"github.com/chadius/terosbattleserver"
	"io"
	"log"
	"os"
)

func main() {
	scriptFilename := "scripts/battle.yml"
	flag.StringVar(&scriptFilename, "f", "scripts/battle.yml", "The filename of the script file. Defaults to scripts/battle.yml")
	flag.Parse()
	squaddieRepositoryFilename := "data/squaddieDatabase.yml"
	powerRepositoryFilename := "data/powerDatabase.yml"

	scriptFile := loadScript(scriptFilename)
	squaddieFile := loadSquaddieRepoYAML(squaddieRepositoryFilename)
	powerFile := loadPowerRepoYAML(powerRepositoryFilename)

	replayErr := terosbattleserver.ReplayBattleScript(scriptFile, squaddieFile, powerFile, os.Stdout)
	if replayErr != nil {
		println(replayErr.Error())
		log.Fatal(replayErr)
	}
}

func loadScript(scriptFilename string) io.Reader {
	scriptFile, err := os.Open(scriptFilename)
	if err != nil {
		println(err.Error())
		log.Fatal(err)
	}
	return scriptFile
}

func loadSquaddieRepoYAML(squaddieRepositoryFilename string) io.Reader {
	squaddieYamlData, err := os.Open(squaddieRepositoryFilename)
	if err != nil {
		log.Fatal(err)
	}

	return squaddieYamlData
}

func loadPowerRepoYAML(powerRepositoryFilename string) io.Reader {
	powerYamlData, err := os.Open(powerRepositoryFilename)
	if err != nil {
		log.Fatal(err)
	}

	return powerYamlData
}