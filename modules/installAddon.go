package modules

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func InstallAddon() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	fmt.Println(colorGreen, "Downloading addon, please wait...", colorReset)
	workshopItemIdFinal, err := os.ReadFile("workshopitemid.data")
	if err != nil {
		log.Fatalf(colorRed, "[!] Unable to read file: %v", err, colorReset)
	}
	//fmt.Println(string(workshopItemIdFinal)) //delete later !!!
	workshopItemIdFinalSTRING := string(workshopItemIdFinal)
	faggotDownload := exec.Command("tools\\steamcmd\\steamcmd.exe", "+login anonymous", "+workshop_download_item", "4000", workshopItemIdFinalSTRING, "+quit")
	faggotDownload.Run()
	UnpackAddon()
}
