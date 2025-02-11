package init

import (
	"autentikasi1/configs"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func init() {
	Load(".env")

	_ = configs.GetConnectionDB()
	_ = configs.GetRedis()
	_ = configs.GetSession()
}

// Load loads the environment variables from the .env file.
func Load(envFile string) {
	err := godotenv.Load(dir(envFile))
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Sprintf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
