package str

import (
	"fmt"
	"gopkg.in/mgutz/ansi.v0"
)

func PrintETR (output, oType, printer string) {
	if oType != "err" && oType != "wrn" {
		oType = "std"
	}
	colourCode := map[string]string {"std": "28+b", "wrn": "226+b", "err": "196+b"}
	code := colourCode [oType]
	fmt.Println (fmt.Sprintf ("\n    $ (%s): %s", ansi.Color (printer, code),
		output))
}
