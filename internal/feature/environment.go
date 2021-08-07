package feature

import (
	"os"

	str "github.com/rubberyconf/rubberyconf/internal/stringarr"
)

func ruleEnvironment(envs []string) bool {

	currentEnv := os.Getenv("ENV")

	return str.Include(envs, currentEnv)

}
