import "math"

func myAtoi(s string) int {
    if len(s) == 0 {
        return 0
    }

    i := 0
    // Step 1: Ignore leading whitespace
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // Step 2: Check for sign
    sign := 1
    if i < len(s) && (s[i] == '-' || s[i] == '+') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    // Step 3: Read in digits until a non-digit character or the end of input is reached
    result := 0
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        // Check for overflow
        if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
            if sign == 1 {
                return math.MaxInt32
            } else {
                return math.MinInt32
            }
        }

        // Update the result
        result = result*10 + digit
        i++
    }

    // Apply the sign to the result
    return result * sign
}
