package str

// PrependTillN repeatedly prepends variable "prepend", to variable "subject", until variable
// "subject" is at least [minLen] chars in size. If length of input "subject" is already equal to or
// greater than "minLen", the function returns a copy of "subject."
func PrependTillN (subject, prepend string, minLen int) (string) {
        for {
                if len (subject) >= minLen {
                        break
                }
                subject = prepend + subject
        }
        return subject
}
