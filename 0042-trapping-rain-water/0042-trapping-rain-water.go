func trap(height []int) int {
    left, right := 0, len(height)-1
    leftMax, rightMax := 0, 0
    result := 0

    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                result += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                result += rightMax - height[right]
            }
            right--
        }
    }

    return result
}