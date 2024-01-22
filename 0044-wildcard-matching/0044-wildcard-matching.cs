public class Solution {
    public bool IsMatch(string s, string p) {
        int m = s.Length;
        int n = p.Length;

        // dp[i, j] represents whether the first i characters in s match the first j characters in p
        bool[,] dp = new bool[m + 1, n + 1];
        dp[0, 0] = true;

        // Handle patterns with '*' at the beginning
        for (int j = 1; j <= n; j++) {
            if (p[j - 1] == '*') {
                dp[0, j] = dp[0, j - 1];
            }
        }

        // Build the dp array
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (p[j - 1] == s[i - 1] || p[j - 1] == '?') {
                    dp[i, j] = dp[i - 1, j - 1];
                } else if (p[j - 1] == '*') {
                    // '*' can either match one character or be empty
                    dp[i, j] = dp[i, j - 1] || dp[i - 1, j];
                }
            }
        }

        return dp[m, n];
    }
}
