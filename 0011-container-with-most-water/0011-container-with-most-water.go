func maxArea(height []int) int {
    maxArea := 0
    left, right := 0, len(height)-1

    for left < right {
        h := min(height[left], height[right])
        w := right - left
        maxArea = max(maxArea, h*w)

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
