// Reverse a String
func reverseString(s string) string {
    rev := ""
    for i := len(s) - 1; i >= 0; i-- {
        rev += string(s[i])
    }
    return rev
}