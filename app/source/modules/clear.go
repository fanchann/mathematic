package modules

import (
	"fmt"
	"os/exec"
)

func Clear() {
	clear, _ := exec.Command("clear").Output()
	fmt.Printf(string(clear))
}
