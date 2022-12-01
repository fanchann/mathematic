package modules

import (
	"fmt"
	"time"

	"project-mtk/app/source/data"
)

func Logo() {
	hariIni := time.Now()
	fmt.Printf(`
      [%v/%v/%v]
   Ayo Belajar matematika
   +---------------------+
   [1] Perkalian
   [2] Pertambahan
   [3] Pengurangan
   [4] Pembagian
   [5] Soal Cerita [ coming soon ]
   +---------------------+
   Benar = %d  |  Salah = %d
   +---------------------+
  Matematika game by FanChann
`, hariIni.Month(), hariIni.Day(), hariIni.Year(), data.Benar, data.Salah)
}
