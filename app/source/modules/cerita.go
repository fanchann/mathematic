package modules

import (
	"fmt"
	"math/rand"
	"time"

	"project-mtk/app/source/data"
)

func Cerita() {
	soal := map[int]string{
		0: "aok",
		1: "Andi membeli sebanyak %d kelereng, Dia membagikan kepada Dina sebanyak %d kelereng, Berapa sisa kelereng andi sekarang",
		2: "kontol anjing",
		3: "awokwokwok",
		4: "eihejn",
	}
	rand.Seed(time.Now().UTC().Unix())
	ack := rand.Intn(4)
	fmt.Println(soal[ack])
	fmt.Println(data.Angka1, data.Angka2)
}
