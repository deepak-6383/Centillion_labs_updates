// Palindrome Check
func isPalindrome(s string) bool {
    rev := ""
    for i := len(s) - 1; i >= 0; i-- {
        rev += string(s[i])
    }
    return s == rev
}
