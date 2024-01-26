class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        // Map to store each number along with its index
        std::unordered_map<int, int> numMap;

        for (int i = 0; i < nums.size(); ++i) {
            int complement = target - nums[i];

            // Check if the complement is already in the map
            if (numMap.find(complement) != numMap.end()) {
                // Return the indices of the two numbers
                return {numMap[complement], i};
            }

            // Add the current number to the map
            numMap[nums[i]] = i;
        }

        // No solution found
        return {};
    }
};