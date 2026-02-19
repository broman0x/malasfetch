package sistem

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

func DeteksiUptime() (string, error) {
	u, err := host.Uptime()
	if err != nil {
		return "Gak tau", err
	}

	jam := u / 3600
	menit := (u % 3600) / 60

	if jam > 0 {
		return fmt.Sprintf("%d hours, %d mins", jam, menit), nil
	}
	return fmt.Sprintf("%d mins", menit), nil
}
