package konfigurasi

type Pengaturan struct {
	Versi      string
	FormatJSON bool
	TanpaWarna bool
}
var PengaturanGlobal = Pengaturan{}
