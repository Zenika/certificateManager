// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/misc/main.go
// 4/16/23 21:35:03

package misc

import "fmt"

func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION         DATE                    COMMENT
-------         ----                    -------
0.100           2023.04.16              near-configs-aware
`)
	fmt.Println()
}