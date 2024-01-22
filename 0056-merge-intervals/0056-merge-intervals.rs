impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    if intervals.is_empty() {
        return Vec::new();
    }

    let mut intervals = intervals;
    intervals.sort_by_key(|interval| interval[0]);

    let mut result = Vec::new();
    let mut current_interval = intervals[0].clone();

    for interval in intervals.into_iter().skip(1) {
        if interval[0] <= current_interval[1] {
            // Overlapping intervals, merge them
            current_interval[1] = current_interval[1].max(interval[1]);
        } else {
            // Non-overlapping interval, push the current merged interval to the result
            result.push(current_interval.clone());
            current_interval = interval.clone();
        }
    }

    // Push the last merged interval
    result.push(current_interval);

    result
    }
}