func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    resultRows := make([]strings.Builder, min(numRows, len(s)))
    currentRow, goingDown := 0, false

    for _, char := range s {
        resultRows[currentRow].WriteRune(char)

        if currentRow == 0 || currentRow == numRows-1 {
            goingDown = !goingDown
        }

        if goingDown {
            currentRow++
        } else {
            currentRow--
        }
    }

    var result strings.Builder
    for _, row := range resultRows {
        result.WriteString(row.String())
    }

    return result.String()
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
