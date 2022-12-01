package data

import (
	"math/rand"
	"time"
)

var (
	Angka1, Angka2, Jawab, Benar, Salah int
	Menu                                string
	Hasil                               interface{}
)

func Random() (int, int) {
	rand.Seed(time.Now().UTC().Unix())
	angka1 := rand.Intn(90)
	angka2 := rand.Intn(85)

	return angka1, angka2
}
