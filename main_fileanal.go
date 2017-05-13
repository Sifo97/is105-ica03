package main
import (
	"fmt"
	"./fileutils"
)

func main() {
	var fileToOpen = "files/lang01.wl"
	var fileToOpen2 = "files/lang02.wl"
	var fileToOpen3 = "files/lang03.wl"
	
	c := fileutils.FileToByteslice(fileToOpen)
	fmt.Println("01: ",c)
	
	c = fileutils.FileToByteslice(fileToOpen2)
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println("02: ",c)
	fmt.Println("------------------------------------------------------------------------------------------")
	c = fileutils.FileToByteslice(fileToOpen3)
	
	fmt.Println("03: ",c)
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println("Printed all 3 files.")
}