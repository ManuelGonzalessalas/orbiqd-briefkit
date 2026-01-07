package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func ResolveRuntimeLogDir() (string, error) {
	dir := os.Getenv("BRIEFKIT_RUNTIME_LOG_DIR")
	if dir == "" {
		dir = "~/.orbiqd/briefkit/logs/runtime/"
	}

	expanded, err := homedir.Expand(dir)
	if err != nil {
		return "", fmt.Errorf("expand runtime log dir: %w", err)
	}

	abs, err := filepath.Abs(expanded)
	if err != nil {
		return "", fmt.Errorf("resolve absolute runtime log dir: %w", err)
	}

	return abs, nil
}
