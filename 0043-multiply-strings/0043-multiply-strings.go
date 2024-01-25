import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	result := make([]int, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := result[i+j+1] + mul
			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	// Remove leading zeros
	var sb strings.Builder
	for _, digit := range result {
		if !(digit == 0 && sb.Len() == 0) {
			sb.WriteRune(rune(digit + '0'))
		}
	}

	if sb.Len() == 0 {
		return "0"
	}

	return sb.String()
}