import "math"

func reverse(x int) int {
    result := 0

    for x != 0 {
        // Extract the last digit
        digit := x % 10

        // Check for overflow before updating the result
        if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
            return 0
        }
        if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
            return 0
        }

        // Update the result by adding the reversed digit
        result = result*10 + digit

        // Move to the next digit in the input
        x /= 10
    }

    return result
}
