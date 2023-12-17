package slmcommon

import (
	"fmt"
	"os"
)

// CheckEnvOrFlag checks if the command-line flag is set; if not, it checks for an environment variable.
// It returns the value and a nil error if set, otherwise an error if neither is set.
func CheckEnvOrFlag(flagValue, envVarName string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}

	envValue, exists := os.LookupEnv(envVarName)
	if !exists {
		return "", fmt.Errorf("no value was provided for '%s'; either provide it as a command-line argument or set an environment variable called '%s'", envVarName, envVarName)
	}

	return envValue, nil
}
