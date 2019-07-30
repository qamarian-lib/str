package str

// MonthToInt () returns the numerical value of a month.
//
// INPUTS
// input 0: The full name of the month (capitalizing only the first letter of the month name, e.g.
// April).
//
// OUTPTS
// outpt 0: The numerical value of the month. If input 0 is invalid, value would be 0.
func MonthToInt (month string) (int) {
        switch month {
                case "January":   return 1
                case "February":  return 2
                case "March":     return 3
                case "April":     return 4
                case "May":       return 5
                case "June":      return 6
                case "July":      return 7
                case "August":    return 8
                case "September": return 9
                case "October":   return 10
                case "November":  return 11
                case "December":  return 12
        }
        return 0
}
