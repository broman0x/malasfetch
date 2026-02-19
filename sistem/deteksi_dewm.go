package sistem

import (
	"os"
	"strings"
)

func DeteksiDEWM() (string, string) {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de == "" {
		de = os.Getenv("DESKTOP_SESSION")
	}

	wm := "Unknown"
	if strings.Contains(strings.ToLower(de), "gnome") {
		wm = "gnome"
	}

	return de, wm
}
