package settings

import (
	"encoding/json"
	"os"
)

type EngineSettings struct {
	Screen *ScreenSettings
	Grid   *GridSettings
}

func Load() (*EngineSettings, error) {
	configPath := "settings.json"
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	s := &EngineSettings{}
	err = json.Unmarshal(configFile, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func Save(s *EngineSettings) error {
	configPath := "settings.json"
	configFile, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
