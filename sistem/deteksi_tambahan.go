package sistem

import (
	"fmt"
	"os/exec"
	"strings"
)

func DeteksiPaket() string {
	var hasil []string
	outRpm, err := exec.Command("rpm", "-qa").Output()
	if err == nil {
		count := strings.Count(string(outRpm), "\n")
		if count > 0 {
			hasil = append(hasil, fmt.Sprintf("%d (rpm)", count))
		}
	}

	outFlat, err := exec.Command("flatpak", "list").Output()
	if err == nil {
		count := strings.Count(string(outFlat), "\n")
		if count > 0 {
			count--
			if count > 0 {
				hasil = append(hasil, fmt.Sprintf("%d (flatpak)", count))
			}
		}
	}

	if len(hasil) == 0 {
		return "Unknown"
	}

	return strings.Join(hasil, ", ")
}

func DeteksiResolusi() string {
	out, err := exec.Command("xrandr").Output()
	if err == nil {
		for _, line := range strings.Split(string(out), "\n") {
			if strings.Contains(line, " connected") && strings.Contains(line, "primary") {
				parts := strings.Fields(line)
				for _, p := range parts {
					if strings.Contains(p, "x") && strings.Contains(p, "+") {
						return strings.Split(p, "+")[0]
					}
				}
			}
		}
	}
	return "1920x1080"
}
