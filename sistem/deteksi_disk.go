package sistem

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

type InformasiDisk struct {
	Total    uint64
	Terpakai uint64
}

func DeteksiDisk() (InformasiDisk, error) {
	infD := InformasiDisk{}
	d, err := disk.Usage("/")
	if err != nil {
		return infD, err
	}

	infD.Total = d.Total
	infD.Terpakai = d.Used

	return infD, nil
}

func (infD InformasiDisk) FormatDisk() string {
	totalGB := float64(infD.Total) / (1024 * 1024 * 1024)
	terpakaiGB := float64(infD.Terpakai) / (1024 * 1024 * 1024)
	persen := (float64(infD.Terpakai) / float64(infD.Total)) * 100
	return fmt.Sprintf("%.1f GB / %.1f GB (%.1f%%)", terpakaiGB, totalGB, persen)
}
