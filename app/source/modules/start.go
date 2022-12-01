package modules

import (
	"fmt"
	"time"

	"project-mtk/app/source/data"
)

func Start() {
	Clear()
	Logo()
	fmt.Printf("\nSilahkan memilih : ")
	fmt.Scanln(&data.Menu)

	switch {
	case data.Menu == "1":
		Perkalian()
		Replay()
	case data.Menu == "2":
		Tambah()
		Replay()
	case data.Menu == "3":
		Kurang()
		Replay()
	case data.Menu == "4":
		Bagi()
		Replay()
	default:
		fmt.Println("Masukan command dengan benar")
		time.Sleep(time.Second)
		Clear()
		Start()
	}
}
