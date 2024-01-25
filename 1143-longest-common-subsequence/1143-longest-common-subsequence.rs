impl Solution {
    pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
        let m = text1.len();
        let n = text2.len();

        // Create a 2D array to store the length of common subsequences
        let mut dp = vec![vec![0; n + 1]; m + 1];

        // Fill the dp array using bottom-up dynamic programming
        for i in 1..=m {
            for j in 1..=n {
                if text1.chars().nth(i - 1).unwrap() == text2.chars().nth(j - 1).unwrap() {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = dp[i - 1][j].max(dp[i][j - 1]);
                }
            }
        }

        // The length of the longest common subsequence is stored in dp[m][n]
        dp[m][n]
    }
}
