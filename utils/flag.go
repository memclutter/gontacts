package utils

import (
	"flag"
	"os"
)

func FlagStringVarEnv(variable *string, name, defaultValue, usage, envName string) {
	dv, ok := os.LookupEnv(envName)
	if !ok {
		dv = defaultValue
	}

	flag.StringVar(variable, name, dv, usage)
}
