use std::collections::HashMap;

impl Solution {
    pub fn find_substring(s: String, words: Vec<String>) -> Vec<i32> {
        let word_len = words[0].len();
        let total_len = word_len * words.len();
        let mut result = Vec::new();

        if s.len() < total_len {
            return result;
        }

        let mut word_count = HashMap::new();
        for word in &words {
            *word_count.entry(word.clone()).or_insert(0) += 1;
        }

        for i in 0..word_len {
            let mut left = i;
            let mut right = i;
            let mut current_count = HashMap::new();

            while right + word_len <= s.len() {
                let current_word = &s[right..right + word_len];
                right += word_len;

                if word_count.contains_key(current_word) {
                    *current_count.entry(current_word.to_string()).or_insert(0) += 1;

                    while current_count[&current_word.to_string()] > word_count[&current_word.to_string()] {
                        let left_word = &s[left..left + word_len];
                        left += word_len;
                        *current_count.get_mut(left_word).unwrap() -= 1;
                    }

                    if right - left == total_len {
                        result.push(left as i32);
                    }
                } else {
                    current_count.clear();
                    left = right;
                }
            }
        }

        result
    }
}
