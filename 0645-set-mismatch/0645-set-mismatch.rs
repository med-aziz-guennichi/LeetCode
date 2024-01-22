impl Solution {
    pub fn find_error_nums(nums: Vec<i32>) -> Vec<i32> {
    let mut result = vec![];

    let n = nums.len() as i32;
    let mut seen = vec![false; n as usize];
    let mut duplicate = 0;
    let mut missing = 0;

    for &num in nums.iter() {
        if seen[num as usize - 1] {
            duplicate = num;
        } else {
            seen[num as usize - 1] = true;
        }
    }

    for (i, &is_seen) in seen.iter().enumerate() {
        if !is_seen {
            missing = (i + 1) as i32;
            break;
        }
    }

    result.push(duplicate);
    result.push(missing);

    result
}
}