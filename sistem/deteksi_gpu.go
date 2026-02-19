package sistem

import (
	"os/exec"
	"strings"
)

func DeteksiGPU() (string, error) {
	cmd := exec.Command("sh", "-c", "lspci | grep -i vga")
	out, err := cmd.Output()
	if err != nil {
		return "Unknown", err
	}

	line := string(out)
	if strings.Contains(line, "controller:") {
		parts := strings.Split(line, "controller:")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1]), nil
		}
	}

	return "Unknown", nil
}
