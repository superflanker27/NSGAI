package main

import (
	"fmt"
	"os/exec"
	"project/modules"
)

func main() {

	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	var input int
	fmt.Println(colorGreen, "Non-Steam gmod addon installer", colorReset)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(colorGreen, "Select an option", colorReset)
	fmt.Println(" ")
	fmt.Println(colorGreen, "[1] Download addon", colorReset)
	fmt.Println(colorGreen, "[2] Patch steamcmd", colorReset)
	fmt.Scan(&input)

	switch input {
	case 1:
		modules.GetAddonInterface()

	case 2:
		cleansteamcmd := exec.Command("cmd", "/c", "cleansteamcmd.bat")
		cleansteamcmd.Run()

	default:
		fmt.Println(colorRed, "[!] Please select a valid option!", colorReset)
	}
}
