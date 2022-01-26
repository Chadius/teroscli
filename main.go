package main

import (
	"bytes"
	"flag"
	"github.com/cserrant/terosCLI/command"
	"io/ioutil"
	"log"
)

func main() {
	scriptFilename := "scripts/battle.yml"
	flag.StringVar(&scriptFilename, "script", "scripts/battle.yml", "The filename of the script file. Defaults to scripts/battle.yml")
	squaddieRepositoryFilename := "data/squaddieDatabase.yml"
	flag.StringVar(&squaddieRepositoryFilename, "squaddie", "data/squaddieDatabase.yml", "The filename of the script file. Defaults to data/squaddieDatabase.yml")
	powerRepositoryFilename := "data/powerDatabase.yml"
	flag.StringVar(&powerRepositoryFilename, "power", "data/powerDatabase.yml", "The filename of the script file. Defaults to data/powerDatabase.yml")
	flag.Parse()

	scriptData, scriptErr := ioutil.ReadFile(scriptFilename)
	if scriptErr != nil {
		log.Fatal(scriptErr)
		return
	}

	squaddieData, squaddieErr := ioutil.ReadFile(squaddieRepositoryFilename)
	if squaddieErr != nil {
		log.Fatal(squaddieErr)
		return
	}

	powerData, powerErr := ioutil.ReadFile(powerRepositoryFilename)
	if scriptErr != nil {
		log.Fatal(powerErr)
		return
	}

	var outputMessage bytes.Buffer
	commandProcessor := command.NewCommandProcessor(nil, nil)
	replayErr := commandProcessor.ApplyRulesetToData(&command.RulesetArguments{
		ScriptData:    scriptData,
		SquaddieData:  squaddieData,
		PowerData:     powerData,
		OutputMessage: &outputMessage,
	})
	if replayErr != nil {
		println(replayErr.Error())
		log.Fatal(replayErr)
	}

	println(outputMessage.String())
}
