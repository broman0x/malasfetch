package ascii

import "strings"

var DaftarLogo = map[string]FungsiASCII{
	"arch":   DapatkanLogoArch,
	"ubuntu": DapatkanLogoUbuntu,
}

func AmbilLogoBerdasarkanID(id string) ([]string, string) {
	id = strings.ToLower(id)

	if fungsi, ada := DaftarLogo[id]; ada {
		return fungsi()
	}

	for kunci, fungsi := range DaftarLogo {
		if strings.Contains(id, kunci) {
			return fungsi()
		}
	}

	return DapatkanLogoDefault()
}
