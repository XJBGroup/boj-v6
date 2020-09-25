package problemconfig

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/spf13/afero"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func Select(configPaths ...string) (string, error) {
	if len(configPaths) == 0 {
		return "", errors.New("nil config files")
	}

	for _, configPath := range configPaths {
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}
	}

	return "", errors.New("no such file in the root directory")
}

func Load(config *ProblemConfig, configPath string) error {
	return LoadFS(afero.NewOsFs(), config, configPath)
}

func LoadFS(filesystem types.Filesystem, config *ProblemConfig, configPath string) error {
	for _, configX := range []struct {
		Type      string
		Unmarshal func([]byte, interface{}) error
	}{
		{".json", json.Unmarshal}, {".yml", yaml.Unmarshal},
		{".toml", toml.Unmarshal}, {".xml", xml.Unmarshal},
		{"", toml.Unmarshal}} {
		if _, err := filesystem.Stat(configPath + configX.Type); err == nil {
			f, err := filesystem.Open(configPath + configX.Type)
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				return err
			}
			err = configX.Unmarshal(b, config)
			if err != nil {
				return err
			}
			config.LoadType = configX.Type
			return nil
		}
	}

	return errors.New("no such file in the root directory")
}

func Save(config *ProblemConfig, configPath string) error {
	return SaveFS(afero.NewOsFs(), config, configPath)
}

func SaveFS(filesystem types.Filesystem, config *ProblemConfig, configPath string) error {
	var b []byte
	var err error
	switch config.LoadType {
	case ".json":
		b, err = json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	case ".yml":
		b, err = yaml.Marshal(config)
		if err != nil {
			return err
		}
	case ".toml":
		b, err = toml.Marshal(config)
		if err != nil {
			return err
		}
	case ".xml":
		b, err = xml.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
	default:
		b, err = toml.Marshal(config)
		if err != nil {
			return err
		}
	}
	if _, err := filesystem.Stat(configPath + config.LoadType); err == nil {
		f, err := filesystem.OpenFile(configPath+config.LoadType, os.O_WRONLY|os.O_TRUNC, 0333)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		f.Close()
		if err != nil {
			return err
		}
		return nil
	} else {
		f, err := filesystem.Create(configPath + config.LoadType)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		f.Close()
		if err != nil {
			return err
		}
		return nil
	}
}
