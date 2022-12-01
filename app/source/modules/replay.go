package modules

import (
	"fmt"
	"os"
	"time"

	"project-mtk/app/source/data"
)

func Replay() {
	fmt.Printf("[*] Apakah mau melanjutkan permainan ini ? y/n : ")
	fmt.Scan(&data.Menu)
	if data.Menu == "y" || data.Menu == "Y" {
		Clear()
		Start()
	} else if data.Menu == "n" || data.Menu == "N" {
		fmt.Printf("Sedang membersihkan terminal, Silahkan tunggu sebentar ^-^\n")
		time.Sleep(2 * time.Second)
		Clear()
		fmt.Println("Semoga harimu menyenangkan ^-^")
		os.Exit(1)
	} else {
		fmt.Println("Apa yang kamu ketik?")
		time.Sleep(2 * time.Second)
		Replay()
	}
}
