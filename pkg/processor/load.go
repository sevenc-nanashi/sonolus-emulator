package processor

import (
	"fmt"
	ioutil "io/ioutil"

	"github.com/sevenc-nanashi/sonolus-emulator/pkg/sonolus"
	log "github.com/sirupsen/logrus"
)

type ProcessorData struct {
	LevelItem sonolus.LevelItem
	LevelData sonolus.LevelData

	EngineItem          sonolus.EngineItem
	EngineData          sonolus.EngineData
	EngineConfiguration sonolus.EngineConfiguration
	EngineRom           []int
}

func (p *ProcessorData) loadLevel(serverUrl string, levelName string) error {
	var levelDetails sonolus.ItemDetails[sonolus.LevelItem]
	itemUrl := fmt.Sprintf("%s/sonolus/levels/%s", serverUrl, levelName)
	log.Infof("Fetching level item from %s", itemUrl)
	err := FetchAndUnmarshal(itemUrl, &levelDetails)
	if err != nil {
		log.Errorf("Failed to fetch level item: %s", err)
		return err
	}
	level := levelDetails.Item
	log.Infof("Fetched level item: %s (#%s)", level.Title, level.Name)

	var levelData sonolus.LevelData
	dataUrl, err := sonolus.JoinUrl(serverUrl, level.Data.Url)
	log.Infof("Fetching level data from %s", dataUrl)
	if err != nil {
		log.Errorf("Failed to fetch level data: %s", err)
		return err
	}
	err = FetchAndUnmarshalWithGunzip(dataUrl, &levelData)
	if err != nil {
		log.Errorf("Failed to fetch level data: %s", err)
		return err
	}

	p.LevelItem = level
	p.LevelData = levelData

	return nil
}

func (p *ProcessorData) loadEngine(serverUrl string, engine sonolus.EngineItem) error {
	log.Infof("Fetching engine %s (#%s)", engine.Title, engine.Name)
	var engineData sonolus.EngineData
	dataUrl, err := sonolus.JoinUrl(serverUrl, engine.Data.Url)
	log.Infof("Fetching engine data from %s", dataUrl)
	if err != nil {
		log.Errorf("Failed to fetch engine data: %s", err)
		return err
	}
	err = FetchAndUnmarshalWithGunzip(dataUrl, &engineData)
	if err != nil {
		log.Errorf("Failed to fetch engine data: %s", err)
		return err
	}
	log.Infof("Fetched engine data")

	var engineConfiguration sonolus.EngineConfiguration
	configurationUrl, err := sonolus.JoinUrl(serverUrl, engine.Configuration.Url)
	log.Infof("Fetching engine configuration from %s", configurationUrl)
	if err != nil {
		log.Errorf("Failed to fetch engine configuration: %s", err)
		return err
	}
	err = FetchAndUnmarshalWithGunzip(configurationUrl, &engineConfiguration)
	if err != nil {
		log.Errorf("Failed to fetch engine configuration: %s", err)
		return err
	}
	log.Infof("Fetched engine configuration")

	var engineRom []int
	if engine.Rom != nil {
		romUrl, err := sonolus.JoinUrl(serverUrl, engine.Rom.Url)
		log.Infof("Fetching engine rom from %s", romUrl)
		if err != nil {
			log.Errorf("Failed to fetch engine rom: %s", err)
			return err
		}
		engineRomReader, err := Fetch(romUrl)
		if err != nil {
			log.Errorf("Failed to fetch engine rom: %s", err)
			return err
		}
    engineRomBytes, err := ioutil.ReadAll(engineRomReader.Body)
		if err != nil {
			log.Errorf("Failed to fetch engine rom: %s", err)
			return err
		}
		defer engineRomReader.Body.Close()
		log.Infof("Fetched engine rom")

		engineRom = make([]int, len(engineRomBytes))
		for i, b := range engineRomBytes {
			engineRom[i] = int(b)
		}
	} else {
		log.Infof("Engine rom is not available")
		engineRom = make([]int, 0)
	}
	p.EngineItem = engine
	p.EngineData = engineData
	p.EngineConfiguration = engineConfiguration
	p.EngineRom = engineRom

	return nil
}
