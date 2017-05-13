package main
import (
	"fmt"
	"./fileutils"
)

func main() {
var fileToOpen = "files/lang01.wl"
fmt.Println("Opening: ", fileToOpen)
c := fileutils.FileToByteslice(fileToOpen)
fmt.Println("method: ",c)
}