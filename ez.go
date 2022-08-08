package ez

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Ez struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
}

type config struct {
	port     string
	renderer string
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

	errorLog, infoLog := e.startLoggers()

	e.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	e.Version = version
	e.ErrorLog = errorLog
	e.InfoLog = infoLog
	e.RootPath = rootPath

	e.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
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

func (e *Ez) startLoggers() (*log.Logger, *log.Logger) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	return errorLog, infoLog
}
