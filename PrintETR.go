package str

import (
	"fmt"
	"gopkg.in/mgutz/ansi.v0"
)

// PrintETR () which stands for "print easier-to-read text", helps print texts that
// would be easier to read.
//
// Inputs
//
// input 0: The text to the be printed.
//
// input 1: The type of text input 0 is. Possible values are: "std" which represents a
// normal text, "wrn" which represents a warning text, and "err" which represents an
// error text. If value is invalid, "std" would be used.
//
// input 2: The name or ID of the printer of the text.
func PrintETR (output, oType, printer string) {
	if oType != "err" && oType != "wrn" {
		oType = "std"
	}
	colourCode := map[string]string {"std": "28+b", "wrn": "226+b", "err": "196+b"}
	code := colourCode [oType]
	fmt.Println (fmt.Sprintf ("\n    $ (%s): %s", ansi.Color (printer, code),
		output))
}
