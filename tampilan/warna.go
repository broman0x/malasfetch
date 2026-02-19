package tampilan

import "fmt"

const (
	Reset = "\033[0m"
	Tebal = "\033[1m"
	Hitam  = "\033[30m"
	Merah  = "\033[31m"
	Hijau  = "\033[32m"
	Kuning = "\033[33m"
	Biru   = "\033[34m"
	Ungu   = "\033[35m"
	Sian   = "\033[36m"
	Putih  = "\033[37m"
	HitamCerah  = "\033[90m"
	MerahCerah  = "\033[91m"
	HijauCerah  = "\033[92m"
	KuningCerah = "\033[93m"
	BiruCerah   = "\033[94m"
	UnguCerah   = "\033[95m"
	SianCerah   = "\033[96m"
	PutihCerah  = "\033[97m"
)

func Warnai(teks string, kodeWarna string, tanpaWarna bool) string {
	if tanpaWarna {
		return teks
	}
	return fmt.Sprintf("%s%s%s%s", Tebal, kodeWarna, teks, Reset)
}
