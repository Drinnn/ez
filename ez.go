package ez

import (
	"fmt"

	"github.com/joho/godotenv"
)

// const version = "1.0.0"

type Ez struct {
	AppName string
	Debug   bool
	Version string
}

func (e *Ez) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middlewares",
		},
	}

	err := e.Init(pathConfig)
	if err != nil {
		return err
	}

	err = e.checkDotEnv(rootPath)
	if err != nil {
		return nil
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (e *Ez) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := e.CreateDirIfDoesntExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Ez) checkDotEnv(path string) error {
	err := e.CreateFileIfDoesntExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}
