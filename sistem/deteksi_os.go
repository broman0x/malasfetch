package sistem

import (
	"os"
	"strings"
)

type InformasiOS struct {
	NamaCantik string
	ID         string
}

func DeteksiOS() (InformasiOS, error) {
	infO := InformasiOS{NamaCantik: "Linux", ID: "linux"}
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return infO, err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			infO.NamaCantik = strings.Trim(line[12:], "\"")
		}
		if strings.HasPrefix(line, "ID=") {
			infO.ID = strings.Trim(line[3:], "\"")
		}
	}

	return infO, nil
}
