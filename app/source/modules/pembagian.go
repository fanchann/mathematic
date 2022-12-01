package modules

import (
	"fmt"
	"time"

	"project-mtk/app/source/data"
)

func Bagi() {
	data.Angka1, data.Angka2 = data.Random()
	data.Hasil = data.Angka1 / data.Angka2
	fmt.Printf("Hasil dari %d : %d = ", data.Angka1, data.Angka2)
	fmt.Scan(&data.Jawab)

	if data.Jawab == data.Hasil {
		fmt.Printf("Jawaban nya adalah %d, dan jawabanmu benar\n", data.Jawab)
		data.Benar += 1
		time.Sleep(2 * time.Second)

	} else {
		fmt.Printf("Jawabanmu salah !, Jawabannya adalah [%d]\n", data.Hasil)
		data.Salah += 1
		time.Sleep(2 * time.Second)
	}

}
