package perintah

import (
	"encoding/json"
	"fmt"
	"malasfetch/ascii"
	"malasfetch/konfigurasi"
	"malasfetch/sistem"
	"malasfetch/tampilan"
	"os"
	"runtime"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
)

var alamatRoot = &cobra.Command{
	Use:   "malasfetch",
	Short: "malasfetch itu alat fetch info sistem yang simpel",
	Long:  `malasfetch nampilin spek sama logo distro buat pamer ke temen atau postingan fesnuk.`,
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("versi")
		if v {
			fmt.Printf("malasfetch versi %s\n", konfigurasi.PengaturanGlobal.Versi)
			return
		}
		jalankanUtama()
	},
}

func init() {
	alamatRoot.Flags().BoolP("versi", "v", false, "Cek versi aplikasinya")
	alamatRoot.Flags().BoolVar(&konfigurasi.PengaturanGlobal.FormatJSON, "json", false, "Output-nya jadi format JSON")
	alamatRoot.Flags().BoolVar(&konfigurasi.PengaturanGlobal.TanpaWarna, "tanpa-warna", false, "Gak usah pake warna-warni")
}

type dataSistem struct {
	NamaUser   string `json:"nama_user"`
	NamaHost   string `json:"nama_host"`
	OS         string `json:"os"`
	Kernel     string `json:"kernel"`
	Uptime     string `json:"uptime"`
	CPU        string `json:"cpu"`
	GPU        string `json:"gpu"`
	RAM        string `json:"ram"`
	Swap       string `json:"swap"`
	Disk       string `json:"disk"`
	Shell      string `json:"shell"`
	Terminal   string `json:"terminal"`
	DE         string `json:"de"`
	WM         string `json:"wm"`
	Paket      string `json:"paket"`
	Resolusi   string `json:"resolusi"`
	VersiGo    string `json:"versi_go"`
	VersiFetch string `json:"versi_fetch"`
	IDDistro   string `json:"id_distro"`
	LocalIP    string `json:"local_ip"`
	Battery    string `json:"battery"`
	Locale     string `json:"locale"`
	Theme      string `json:"theme"`
	Icons      string `json:"icons"`
	Font       string `json:"font"`
	Cursor     string `json:"cursor"`
}

func jalankanUtama() {
	pengguna, nHost, _ := sistem.DeteksiUser()
	osInfo, _ := sistem.DeteksiOS()
	uptime, _ := sistem.DeteksiUptime()
	cpuInfo, _ := sistem.DeteksiCPU()
	gpuInfo, _ := sistem.DeteksiGPU()
	ramInfo, _ := sistem.DeteksiRAM()
	diskInfo, _ := sistem.DeteksiDisk()
	shell, _ := sistem.DeteksiShell()
	terminal, _ := sistem.DeteksiTerminal()
	de, wm := sistem.DeteksiDEWM()
	paket := sistem.DeteksiPaket()
	resolusi := sistem.DeteksiResolusi()
	lanjutan := sistem.DeteksiLanjutan()

	h, _ := host.Info()
	kernel := "Linux"
	if h != nil {
		kernel = fmt.Sprintf("%s %s %s", h.OS, h.KernelVersion, h.KernelArch)
	}

	data := dataSistem{
		NamaUser:   pengguna,
		NamaHost:   nHost,
		OS:         osInfo.NamaCantik,
		IDDistro:   osInfo.ID,
		Kernel:     kernel,
		Uptime:     uptime,
		CPU:        cpuInfo.FormatCPU(),
		GPU:        gpuInfo,
		RAM:        ramInfo.FormatRAM(),
		Swap:       lanjutan.Swap,
		Disk:       diskInfo.FormatDisk(),
		Shell:      shell,
		Terminal:   terminal,
		DE:         de,
		WM:         wm,
		Paket:      paket,
		Resolusi:   resolusi,
		VersiGo:    runtime.Version(),
		VersiFetch: konfigurasi.PengaturanGlobal.Versi,
		LocalIP:    lanjutan.LocalIP,
		Battery:    lanjutan.Battery,
		Locale:     lanjutan.Locale,
		Theme:      lanjutan.Theme,
		Icons:      lanjutan.Icons,
		Font:       lanjutan.Font,
		Cursor:     lanjutan.Cursor,
	}

	if konfigurasi.PengaturanGlobal.FormatJSON {
		hasil, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(hasil))
		return
	}

	logo, warna := ascii.AmbilLogoBerdasarkanID(data.IDDistro)
	perender := &tampilan.Perender{
		TanpaWarna: konfigurasi.PengaturanGlobal.TanpaWarna,
	}

	w := tampilan.Warnai
	c := warna
	tw := konfigurasi.PengaturanGlobal.TanpaWarna

	garisPemisah := ""
	for i := 0; i < len(data.NamaUser)+len(data.NamaHost)+1; i++ {
		garisPemisah += "-"
	}

	lbel := func(s string) string {
		return w(s, c, tw)
	}

	info := []string{
		fmt.Sprintf("%s%s%s", lbel(data.NamaUser), w("@", tampilan.Putih, tw), lbel(data.NamaHost)),
		w(garisPemisah, tampilan.Putih, tw),
		fmt.Sprintf("%s %s", lbel("OS:"), data.OS),
		fmt.Sprintf("%s %s", lbel("Host:"), data.NamaHost),
		fmt.Sprintf("%s %s", lbel("Kernel:"), data.Kernel),
		fmt.Sprintf("%s %s", lbel("Uptime:"), data.Uptime),
		fmt.Sprintf("%s %s", lbel("Packages:"), data.Paket),
		fmt.Sprintf("%s %s", lbel("Shell:"), data.Shell),
		fmt.Sprintf("%s %s", lbel("Display:"), data.Resolusi),
		fmt.Sprintf("%s %s", lbel("DE:"), data.DE),
		fmt.Sprintf("%s %s", lbel("WM:"), data.WM),
		fmt.Sprintf("%s %s", lbel("Theme:"), data.Theme),
		fmt.Sprintf("%s %s", lbel("Icons:"), data.Icons),
		fmt.Sprintf("%s %s", lbel("Font:"), data.Font),
		fmt.Sprintf("%s %s", lbel("Cursor:"), data.Cursor),
		fmt.Sprintf("%s %s", lbel("Terminal:"), data.Terminal),
		fmt.Sprintf("%s %s", lbel("CPU:"), data.CPU),
		fmt.Sprintf("%s %s", lbel("GPU:"), data.GPU),
		fmt.Sprintf("%s %s", lbel("Memory:"), data.RAM),
		fmt.Sprintf("%s %s", lbel("Swap:"), data.Swap),
		fmt.Sprintf("%s %s", lbel("Disk (/):"), data.Disk),
		fmt.Sprintf("%s %s", lbel("Local IP:"), data.LocalIP),
		fmt.Sprintf("%s %s", lbel("Battery:"), data.Battery),
		fmt.Sprintf("%s %s", lbel("Locale:"), data.Locale),
		"",
		tampilan.Hitam + "██" + tampilan.Merah + "██" + tampilan.Hijau + "██" + tampilan.Kuning + "██" + tampilan.Biru + "██" + tampilan.Ungu + "██" + tampilan.Sian + "██" + tampilan.Putih + "██" +
			tampilan.HitamCerah + "██" + tampilan.MerahCerah + "██" + tampilan.HijauCerah + "██" + tampilan.KuningCerah + "██" + tampilan.BiruCerah + "██" + tampilan.UnguCerah + "██" + tampilan.SianCerah + "██" + tampilan.PutihCerah + "██",
	}
	perender.Tampilkan(logo, warna, info)
}

var alamatUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "Hapus malasfetch dari sistem",
	Run: func(cmd *cobra.Command, args []string) {
		target := "/usr/local/bin/malasfetch"
		if _, err := os.Stat(target); os.IsNotExist(err) {
			fmt.Println("malasfetch emang gak ada di /usr/local/bin")
			return
		}

		err := os.Remove(target)
		if err != nil {
			fmt.Printf("Gagal hapus: %v\nCoba pake 'sudo malasfetch uninstall'\n", err)
			return
		}

		fmt.Println("malasfetch udah resmi dihapus dari sistem!")
	},
}

func Eksekusi(versi string) {
	alamatRoot.AddCommand(alamatUninstall)
	konfigurasi.PengaturanGlobal.Versi = versi
	if err := alamatRoot.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
