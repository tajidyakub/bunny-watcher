package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ApiCfg struct {
	Host string `yaml:"host"`
	Key  string `yaml:"key"`
}

type RemoteCfg struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type LocalCfg struct {
	Path string `yaml:"path"`
}

type Config struct {
	Api    ApiCfg    `yaml:"api"`
	Remote RemoteCfg `yaml:"remote"`
	Local  LocalCfg  `yaml:"local"`
}

func initConfig(path string) (err error) {
	InitialConfig := Config{
		Api: ApiCfg{
			Host: "storage.bunny-cdn.net",
			Key:  "replace-with-your-api-key",
		},
		Remote: RemoteCfg{
			Name: "storage_name",
			Path: "/storage/sub/path",
		},
		Local: LocalCfg{
			Path: "/your/local/path",
		},
	}

	yamlFile, err := yaml.Marshal(&InitialConfig)

	if err != nil {
		fmt.Println("Error while constructing yaml.")
		return err
	}

	dirName := filepath.Dir(path)
	errDir := os.MkdirAll(dirName, 0775)

	if errDir != nil {
		fmt.Println("Error while creating directory. " + dirName)
		return errDir
	}

	f, err := os.Create(path)

	if err != nil {

		fmt.Println("Error while creating file. " + path)

		return err
	} else {
		defer f.Close()

		_, errWrite := io.WriteString(f, string(yamlFile))

		if errWrite != nil {
			return errWrite
		}

		fmt.Println("Config file created at " + path + " re-run the app after config adjusted.")
		fmt.Println(string(yamlFile))
	}

	return nil
}

func main() {
	var config Config

	userHome := os.Getenv("HOME")
	configPath := userHome + "/.config/bunny-watcher/bunny.yaml"

	yamlFile, err := os.ReadFile(configPath)

	if err != nil {
		fmt.Println("Config file not found, creating...")

		errCreate := initConfig(configPath)

		if errCreate != nil {
			fmt.Println("You need to create the config file manually.")
			panic(errCreate)
		}

		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	fmt.Println("Validating the config file.")

	i, errExist := os.Stat(config.Local.Path)

	if errExist != nil {
		panic(errExist)
	}

	fmt.Println("Local path exists.")

	fmt.Println(i)
}
