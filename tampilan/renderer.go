package tampilan

import (
	"fmt"
	"strings"
)

type Perender struct {
	TanpaWarna bool
}

func (p *Perender) Tampilkan(logo []string, warna string, info []string) {
	maxLines := len(logo)
	if len(info) > maxLines {
		maxLines = len(info)
	}

	for i := 0; i < maxLines; i++ {
		lineLogo := ""
		if i < len(logo) {
			lineLogo = logo[i]
		} else {
			if len(logo) > 0 {
				lebarLogo := p.hitungLebarBersih(logo[0])
				lineLogo = strings.Repeat(" ", lebarLogo)
			}
		}

		lineInfo := ""
		if i < len(info) {
			lineInfo = info[i]
		}

		fmt.Printf("%s    %s\n", lineLogo, lineInfo)
	}
}

func (p *Perender) hitungLebarBersih(s string) int {
	bersih := ""
	diDalamANSI := false
	for _, r := range s {
		if r == '\033' {
			diDalamANSI = true
			continue
		}
		if diDalamANSI {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				diDalamANSI = false
			}
			continue
		}
		bersih += string(r)
	}
	return len(bersih)
}
