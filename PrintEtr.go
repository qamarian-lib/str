package str

import (
	"fmt"
	"gopkg.in/mgutz/ansi.v0"
)

// PrintEtr () which stands for "print easier-to-read", helps print texts  that would be
// easier to read.
//
// Inputs
//
// input 0: The item to the be printed.
//
// input 1: The type of data input 0 is. Possible values are: "std" which represents a
// normal data, "wrn" which represents a warning data, and "err" which represents an
// error data. If value is invalid, "std" would be used.
//
// input 2: The name or ID of the printer of the text.
func PrintEtr (output interface {}, oType, printer string) {
	if oType != "err" && oType != "wrn" {
		oType = "std"
	}
	colourCode := map[string]string {"std": "28+b", "wrn": "226+b", "err": "196+b"}
	code := colourCode [oType]
	meta := fmt.Sprintf ("\n    $ (%s):", ansi.Color (printer, code))
	fmt.Println (meta, output)
}
