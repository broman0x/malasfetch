package sistem

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"github.com/shirou/gopsutil/v3/mem"
)

type InformasiLanjutan struct {
	Swap    string
	LocalIP string
	Battery string
	Locale  string
	Theme   string
	Icons   string
	Font    string
	Cursor  string
}

func DeteksiLanjutan() InformasiLanjutan {
	info := InformasiLanjutan{}
	s, err := mem.SwapMemory()
	if err == nil {
		persen := 0.0
		if s.Total > 0 {
			persen = float64(s.Used) / float64(s.Total) * 100
		}
		info.Swap = fmt.Sprintf("%s / %s (%.0f%%)", formatSize(s.Used), formatSize(s.Total), persen)
	}

	info.LocalIP = deteksiIP()
	info.Battery = deteksiBaterai()
	info.Locale = os.Getenv("LANG")
	if info.Locale == "" {
		info.Locale = "en_US.UTF-8"
	}
	info.Theme = ambilGSetting("org.gnome.desktop.interface", "gtk-theme")
	info.Icons = ambilGSetting("org.gnome.desktop.interface", "icon-theme")
	info.Font = ambilGSetting("org.gnome.desktop.interface", "font-name")
	info.Cursor = ambilGSetting("org.gnome.desktop.interface", "cursor-theme")

	return info
}

func formatSize(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func deteksiIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "Unknown"
	}
	for _, i := range ifaces {
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return fmt.Sprintf("%s (%s)", ip.String(), i.Name)
		}
	}
	return "Disconnected"
}

func deteksiBaterai() string {
	path := "/sys/class/power_supply"
	files, err := os.ReadDir(path)
	if err != nil {
		return "N/A"
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), "BAT") {
			cap, _ := os.ReadFile(fmt.Sprintf("%s/%s/capacity", path, f.Name()))
			status, _ := os.ReadFile(fmt.Sprintf("%s/%s/status", path, f.Name()))

			c := strings.TrimSpace(string(cap))
			s := strings.TrimSpace(string(status))

			if c != "" && s != "" {
				return fmt.Sprintf("%s%% [%s]", c, s)
			}
		}
	}
	return "N/A"
}

func ambilGSetting(schema, key string) string {
	out, err := exec.Command("gsettings", "get", schema, key).Output()
	if err != nil {
		return "Unknown"
	}
	return strings.Trim(strings.TrimSpace(string(out)), "'")
}
