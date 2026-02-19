package sistem

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

type InformasiRAM struct {
	Terpakai uint64
	Total    uint64
}

func DeteksiRAM() (InformasiRAM, error) {
	infR := InformasiRAM{}

	v, err := mem.VirtualMemory()
	if err != nil {
		return infR, err
	}

	infR.Total = v.Total
	infR.Terpakai = v.Used

	return infR, nil
}

func (infR InformasiRAM) FormatRAM() string {
	totalGiB := float64(infR.Total) / (1024 * 1024 * 1024)
	terpakaiGiB := float64(infR.Terpakai) / (1024 * 1024 * 1024)
	persen := (float64(infR.Terpakai) / float64(infR.Total)) * 100
	return fmt.Sprintf("%.2f GiB / %.2f GiB (%.0f%%)", terpakaiGiB, totalGiB, persen)
}
