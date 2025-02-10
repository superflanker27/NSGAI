package modules

import (
	"fmt"
	"os"
	"regexp"
)

func GetAddonInterface() {

	colorReset := "\033[0m"
	//colorRed := "\033[31m"
	colorGreen := "\033[32m"
	fmt.Println(colorGreen, "Enter link of workshop item:", colorReset)
	GetAddonID()
}

func GetAddonID() {
	var addonLink string
	fmt.Scan(&addonLink)
	reg := regexp.MustCompile("[0-9]+")
	match := reg.FindAllString(addonLink, -1)
	itemid := ""
	for _, finalmatch := range match {
		itemid += finalmatch
	}
	//fmt.Println("(debugggg) ITEM ID - ", itemid)
	filename := "workshopitemid.data"
	os.WriteFile(filename, []byte(itemid), 0644) // 0644 permissions
	InstallAddon()
}
