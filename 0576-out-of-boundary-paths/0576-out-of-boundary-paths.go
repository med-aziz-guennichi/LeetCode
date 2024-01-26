func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    MOD := 1000000007
    dp := make([][][]int, maxMove+1)
    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }

    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for move := 1; move <= maxMove; move++ {
        for row := 0; row < m; row++ {
            for col := 0; col < n; col++ {
                for _, dir := range directions {
                    newRow, newCol := row+dir[0], col+dir[1]
                    if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n {
                        dp[move][row][col]++
                    } else {
                        dp[move][row][col] = (dp[move][row][col] + dp[move-1][newRow][newCol]) % MOD
                    }
                }
            }
        }
    }

    return dp[maxMove][startRow][startColumn]
}