package util

import (
	"os"
)

// IsCI returns true if the current environment is CI.
func IsCI() bool {
	return os.Getenv("CI") != "" ||
		os.Getenv("CONTINUOUS_INTEGRATION") != "" ||
		os.Getenv("BUILD_NUMBER") != "" ||
		os.Getenv("RUN_ID") != ""
}
