package str

import (
        "crypto/rand"
        "encoding/base32"
	"errors"
        "strconv"
        "time"
)

// UniquePredsafeStr () generates a string that is unique, and also safe from prediction (i.e. it is
// difficult to precisely predict what a string generated, by this function, at anytime would be).
//
// Inputs
//
// input 0: The length of the string to be generated. Value can not be less than 32.
//
// Outpts
//
// outpt 0: If generation succeeds, value would be the generated string; an alpha-numeric string
// (case-insensitive). If generation should fail, value could be anything. Use outpt 1 to determine
// success or failure of the function.
//
// outpt 1: If generation succeeds, value would be a nil error. If generation should fail, value
// would be an error describing the reason for the failure. If value of input 0 is less than 32,
// value would be error str.UPSErrTooSmall.
//
// Notes on Current Implementation
//
// This function relies on your system's local time (golang's package "time"). If your system should
// go back in time, all strings generated till your system goes back to the present, may be unique,
// but not guaranteed to be unique.
//
// In short, for this function to generate unique strings, at all time (till the end of the year
// 9999), your system's time must be accurate at all times you call this function.
//
// This function builds upon golang's package "crypto/rand", so its prediction-safeness is
// dependent on the package.
func UniquePredsafeStr (length int) (newStr string, err error) { /* This function uses the current
	local time and random numbers, to generated the strings it generates. Time helps ensure the
	strings generated are unique, while random numbers help ensure the strings are
	prediction-safe. */

	// Error handling, due to input string length being less than 32
	if length < 32 {
		return newStr, UPSErrTooSmall
	}

	/* This part generates a 23-character string which would be used as the unqiue part of the
	string to be generated. This string would be based on the current local time of the machine.
	Format: NNNNNNNNNSSmmHHDDMMYYYY. NNNNNNNNN refers to the nanosecond, SS refers to the
	second, mm refers to the minute, HH refers to the hour, DD refers to the day, MM refers to
	the month, and YYYY refers to the year.
	{ ... */
		currentTime := time.Now ()

		// Adding of the nanosecond part { ...
                nanoSec := currentTime.Nanosecond ()
                nanoSecString := PrependTillN (strconv.Itoa (nanoSec), "0", 9)
                newStr += nanoSecString
                // ... }

                // Adding of the second part { ...
                sec := currentTime.Second ()
                secString := PrependTillN (strconv.Itoa (sec), "0", 2)
                newStr += secString
                // ... }

		// Adding of the minute part { ...
                min := currentTime.Minute ()
                minString := PrependTillN (strconv.Itoa (min), "0", 2)
                newStr += minString
                // ... }

		// Adding of the hour part { ...
                hour := currentTime.Hour ()
                hourString := PrependTillN (strconv.Itoa (hour), "0", 2)
                newStr += hourString
                // ... }

                // Adding of the day part { ...
                day := currentTime.Day ()
                dayString := PrependTillN (strconv.Itoa (day), "0", 2)
                newStr += dayString
                // ... }

                // Adding of the month part { ...
                month := MonthToNum (currentTime.Month ().String ())
                monthString := PrependTillN (strconv.Itoa (month), "0", 2)
                newStr += monthString
                // ... }

                // Adding of the year part { ...
                year := currentTime.Year ()
                yearString := PrependTillN (strconv.Itoa (year), "0", 4)
                newStr += yearString
                // ... }
        // ... }

	// Error handling, due to length of unique part not being 19
	if len (newStr) != 23 {
		return newStr, upsErrBuggy
	}

        // This part generates the prediction-safe part of the string. { ...
        safeBytes := make ([]byte, length - 23)
        _, err1 := rand.Read (safeBytes)
        	if err1 != nil {
                	return newStr, err1
        	}
        safeStr := base32.StdEncoding.EncodeToString (safeBytes)
		// Error handling, due to length of prediction-safe part not long as expected
		if len (safeStr) < (length - 23) {
			return newStr, upsErrBuggy
		}
        newStr = safeStr[0:length - 23] + newStr
        // ... }

        return newStr, err
}

var (
	UPSErrTooSmall error = errors.New ("This function does not generate strings less than 32 " +
		"characters.")
	upsErrBuggy error = errors.New ("This function has a bug.")
)
