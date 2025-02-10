package modules

import (
	"fmt"
	"os"
	"os/exec"
)

func UnpackAddon() {

	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	gmodDirbytes, err := os.ReadFile("gmodlocation.cfg")
	if err != nil {
		fmt.Println(colorRed, "Error!", colorReset)
	}
	gmodDir := string(gmodDirbytes)
	addonIDfnbytes, err2 := os.ReadFile("workshopitemid.data")
	if err2 != nil {
		fmt.Println(colorRed, "Error!", colorReset)
	}
	gmadDir := gmodDir + "\\bin\\gmad.exe"

	// not all addons are named gmpublisher, some have other names, so take name as user input
	var gmaname string
	fmt.Println(colorGreen, "Enter the name of the gma file: ", colorReset)
	fmt.Scan(&gmaname)

	addonIDfn := string(addonIDfnbytes)
	unpackEx := exec.Command(gmadDir, "tools\\steamcmd\\steamapps\\workshop\\content\\4000\\"+addonIDfn+"\\"+gmaname+".gma")
	unpackEx.Run()
	//fmt.Println(gmadDir, "tools\\steamcmd\\steamapps\\workshop\\content\\4000\\"+addonIDfn+"\\"+gmaname+".gma")

	// drop to addons
	gmodaddonsDir := gmodDir + "\\garrysmod\\addons"
	unpackedAddonDir := "tools\\steamcmd\\steamapps\\workshop\\content\\4000\\" + addonIDfn + "\\" + gmaname

	drop2addons := exec.Command("cmd", "/c", "move", unpackedAddonDir, gmodaddonsDir)
	drop2addons.Run()
	//fmt.Println(unpackedAddonDir, gmodaddonsDir)

	// cleanup
	unpackedAddonDirClean := "tools\\steamcmd\\steamapps\\workshop\\content\\4000\\" + addonIDfn + "\\" + gmaname + ".gma"
	cleanupgma := exec.Command("cmd", "/c", "del", unpackedAddonDirClean)
	cleanupgma.Run()
}
