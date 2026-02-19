package sistem

import (
	"fmt"
	"strings"
	"github.com/shirou/gopsutil/v3/cpu"
)

type InformasiCPU struct {
	Model      string
	JumlahCore int
	Speed      float64
}

func DeteksiCPU() (InformasiCPU, error) {
	infC := InformasiCPU{Model: "Unknown", JumlahCore: 0}

	cpus, err := cpu.Info()
	if err != nil {
		return infC, err
	}

	if len(cpus) > 0 {
		infC.Model = cpus[0].ModelName
		infC.JumlahCore = len(cpus)
		infC.Speed = cpus[0].Mhz
	}

	return infC, nil
}

func (infC InformasiCPU) FormatCPU() string {
	model := strings.TrimSpace(infC.Model)
	return fmt.Sprintf("%s (%d) @ %.2f GHz", model, infC.JumlahCore, infC.Speed/1000)
}
