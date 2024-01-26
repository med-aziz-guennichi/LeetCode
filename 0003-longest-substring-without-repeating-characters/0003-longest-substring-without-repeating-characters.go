func lengthOfLongestSubstring(s string) int {
    n := len(s)
    seen := make(map[byte]int)
    maxLen := 0

    for left, right := 0, 0; right < n; right++ {
        if idx, exists := seen[s[right]]; exists && idx >= left {
            left = idx + 1
        }
        seen[s[right]] = right
        if currLen := right - left + 1; currLen > maxLen {
            maxLen = currLen
        }
    }

    return maxLen
}