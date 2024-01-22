use std::cmp::{max, min};

impl Solution {
    fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (nums1, nums2) = if nums1.len() > nums2.len() {
            (nums2, nums1)
        } else {
            (nums1, nums2)
        };

        let (x, y) = (nums1.len(), nums2.len());
        let (mut low, mut high) = (0, x);

        while low <= high {
            let partition_x = (low + high) / 2;
            let partition_y = (x + y + 1) / 2 - partition_x;

            let max_x = if partition_x == 0 { i32::MIN } else { nums1[partition_x - 1] };
            let min_x = if partition_x == x { i32::MAX } else { nums1[partition_x] };

            let max_y = if partition_y == 0 { i32::MIN } else { nums2[partition_y - 1] };
            let min_y = if partition_y == y { i32::MAX } else { nums2[partition_y] };

            if max_x <= min_y && max_y <= min_x {
                if (x + y) % 2 == 0 {
                    return f64::from(max(max_x, max_y) + min(min_x, min_y)) / 2.0;
                } else {
                    return f64::from(max(max_x, max_y));
                }
            } else if max_x > min_y {
                high = partition_x - 1;
            } else {
                low = partition_x + 1;
            }
        }

        0.0 // Default value, it won't be reached
    }
}
