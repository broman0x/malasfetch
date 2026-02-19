package sistem

import (
	"os"
	"strings"
)

func DeteksiShell() (string, error) {
	s := os.Getenv("SHELL")
	if s == "" {
		return "Unknown", nil
	}

	parts := strings.Split(s, "/")
	return parts[len(parts)-1], nil
}

func DeteksiTerminal() (string, error) {
	t := os.Getenv("TERM_PROGRAM")
	if t == "" {
		t = os.Getenv("TERMINAL")
	}
	if t == "" {
		t = os.Getenv("TERM")
	}

	if t == "" {
		return "Unknown", nil
	}

	return t, nil
}
