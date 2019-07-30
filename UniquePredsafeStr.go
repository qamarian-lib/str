package str

import (
        "crypto/rand"
        "encoding/base32"
	"errors"
        "strconv"
        "time"
)

// UniquePredsafeStr () generates a string that is unique, and also safe from prediction (i.e. it is
// difficult to precisely predict what a string generated at anytime (by this function) would be).
//
// INPUTS
// input 0: The length of the string to be generated. Value can not be less than 24.
//
// OUTPTS
// outpt 0: If generation succeeds, value would be the generated string. If generation should fail,
// value could be anything. Use outpt 1 to determine success or failure.
//
// outpt 1: If generation succeeds, value would be a nil error. If generation should fail, value
// would be an error describing the reason for the failure. If value of input 0 is less than 24,
// value would be error str.UPSErrTooSmall.
//
// CURRENT IMPLEMENTATION INFO
// This function relies on your system's local time (golang's package "time"). If your system should
// go back in time, all strings generated till your system goes back to the present, may be unique,
// but not guaranteed to be unique.
//
// In short, for this function to generate unique strings, at all time (till the end of the year
// 9999), your system's time must be accurate at all time.
//
// This function builds upon golang's package "crypto/rand", so its prediction-safeness is
// dependent on the package.
func UniquePredsafeStr (length int) (newStr string, err error) { /* This function uses the current
	local time and random numbers, to generated the strings it generates. Time helps ensure the
	strings generated are unique, while random numbers help ensure the strings are
	prediction-safe. */

	// Error handling, due to input string length being less than 24
	if length < 24 {
		return newStr, UPSErrTooSmall
	}

	/* This part generates a 19-character string which would be used as the unqiue part of the
	string to be generated. This string would be based on the current local time of the machine.
	Format: SSNNNNNNNNNDDMMYYYY. SS refers to the second, NNNNNNNNN refers to the nanosecond, DD
	refers to the day, MM refers to the month, and YYYY refers to the year.
	{ ... */
		currentTime := time.Now ()

                // Adding of the second part { ...
                sec := currentTime.Second ()
                secString := PrependTillN (strconv.Itoa (sec), "0", 2)
                newStr += secString
                // ... }

                // Adding of the nanosecond part { ...
                nanoSec := currentTime.Nanosecond ()
                nanoSecString := PrependTillN (strconv.Itoa (nanoSec), "0", 9)
                newStr += nanoSecString
                // ... }

                // Adding of the day part { ...
                day := currentTime.Day ()
                dayString := PrependTillN (strconv.Itoa (day), "0", 2)
                newStr += dayString
                // ... }

                // Adding of the months part { ...
                month := MonthToInt (currentTime.Month ().String ())
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
	if len (newStr) != 19 {
		return newStr, upsErrBuggy
	}

        // This part generates the prediction-safe part of the string. { ...
        safeBytes := make ([]byte, length - 19)
        _, err1 := rand.Read (safeBytes)
        if err1 != nil {
                return newStr, err1
        }
        safeStr := base32.StdEncoding.EncodeToString (safeBytes)
        newStr = safeStr[0:length - 19] + newStr
        // ... }

        return newStr, err
}

var (
	UPSErrTooSmall error = errors.New ("This function does not generate strings less than 24 " +
		"characters.")
	upsErrBuggy error = errors.New ("This function has a bug.")
)
