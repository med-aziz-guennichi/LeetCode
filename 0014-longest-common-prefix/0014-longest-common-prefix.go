func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    // Iterate through characters of the first string
    for i := 0; i < len(strs[0]); i++ {
        currentChar := strs[0][i]

        // Compare with the same position in other strings
        for j := 1; j < len(strs); j++ {
            if i == len(strs[j]) || strs[j][i] != currentChar {
                // If characters are not equal or we reach the end of any string, return the prefix found so far
                return strs[0][:i]
            }
        }
    }

    // If all strings have been compared and no mismatch found, return the entire first string as the common prefix
    return strs[0]
}
