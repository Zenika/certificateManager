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
0.100           2023.04.16              near-config-aware
0.200			2023.04.18				ca subcommands
`)
	fmt.Println()
}
