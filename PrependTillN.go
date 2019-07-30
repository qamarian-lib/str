package str

// PrependTillN () repeatedly prepends a prepend to a subject, until the length of the subject is at
// least N chars in size.
//
// INPUTS
//
// input 0: The subject.
//
// input 1: The prepend. If value is an empty string, input 0 is returned.
//
// input 2: Size N. The highest value of N is the highest value of type int. However, when the value
// of N is very high (like 1000000000 (1 billion)), this function could block indefinitely, due to
// the way Go works (behaviour has been confirmed in Go versions <= 1.12.4). If this function starts
// to block with high values of N, you can lower your N, but increase the number of times the
// function is called. So, instead of this:
//
// text = str.PrependTillN (text, "hi!", 1000000000)
//
// you can do this:
//
// text = str.PrependTillN (text, "hi!", 500000000)
// text = str.PrependTillN (text, "hi!", 500000000)
//
// OUTPTS
//
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
