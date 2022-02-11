package main

import "zura.org/mypackage"

func main() {
	imp := mypackage.NewMyPrint("Hello, MyPrint!")
	imp.Print("Yeah! ")

	impe := mypackage.NewMyPrintExtra("Hyahha! MyPrintExtra!!!", "8> ")
	impe.Print("Weeeeee ")
	impe.PrintExtra("Doooooo ")
}
