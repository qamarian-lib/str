package str

// PrependTillN repeatedly prepends a prepend, to a subject, until the length of the subject is at
// least N chars in size.
//
// INPUTS
// input 0: The subject.
// input 1: The prepend. If value is an empty string, input 0 is returned.
// input 2: Size N.
//
// OUTPTS
// outpt 0: Result of operation.
func PrependTillN (subject, prepend string, minLen int) (string) {

	if prepend == "" {
		return subject
	}

        for {
                if len (subject) >= minLen {
                        break
                }
                subject = prepend + subject
        }
        return subject
}
