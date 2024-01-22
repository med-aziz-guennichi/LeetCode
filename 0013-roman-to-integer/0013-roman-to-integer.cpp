class Solution {
public:
    int romanToInt(string s) {
    std::unordered_map<char, int> romanNumerals = {
        {'I', 1},
        {'V', 5},
        {'X', 10},
        {'L', 50},
        {'C', 100},
        {'D', 500},
        {'M', 1000}
    };

    int result = 0;

    for (int i = 0; i < s.length(); ++i) {
        int currentNumeral = romanNumerals[s[i]];
        int nextNumeral = (i + 1 < s.length()) ? romanNumerals[s[i + 1]] : 0;

        if (nextNumeral > currentNumeral) {
            // If the current numeral is less than the next numeral, subtract it
            result += nextNumeral - currentNumeral;
            ++i; // Skip the next numeral since it has been considered
        } else {
            // Otherwise, add the current numeral to the result
            result += currentNumeral;
        }
    }

    return result;
}

};