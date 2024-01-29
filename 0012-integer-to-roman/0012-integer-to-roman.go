func intToRoman(num int) string {
    // Define the Roman numeral symbols and their corresponding values
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    result := ""
    // Iterate through the symbols and values
    for i := 0; num > 0; i++ {
        // Repeatedly subtract the largest possible value and append the corresponding symbol
        for num >= values[i] {
            num -= values[i]
            result += symbols[i]
        }
    }

    return result
}
