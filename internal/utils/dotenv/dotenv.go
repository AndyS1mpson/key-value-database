package dotenv

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type (
	params struct {
		filename string
		override bool
	}
)

const (
	dotEnvFileName = ".env"
)

// Load загружает данные из .env файла
func Load() {
	if err := readFile(dotEnvFileName); err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
		default:
			panic(fmt.Errorf("Error loading envs: %s", err))
		}
	}
}

func readFile(filename string) error {
	var dir string

	dir = findFileDir(filename, searchCallerFile())

	envFile := filepath.Join(dir, filename)
	_, err := os.Stat(envFile)
	if err != nil {
		return err
	}

	load := godotenv.Load

	if err := load(envFile); err != nil {
		return fmt.Errorf("loading %s file, %s", envFile, err)
	}

	return nil
}

func findFileDir(filename string, from string) string {
	dir := filepath.Dir(from)
	gopath := filepath.Clean(os.Getenv("GOPATH"))
	for dir != "/" && dir != gopath {
		path := filepath.Join(dir, filename)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			dir = filepath.Dir(dir)
			continue
		}
		return dir
	}
	return ""
}

func searchCallerFile() string {
	_, file, _, _ := runtime.Caller(1)
	currentFile := file
	for i := 2; file == currentFile; i++ {
		_, file, _, _ = runtime.Caller(i)
	}

	return file
}
