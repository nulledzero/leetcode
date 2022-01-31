class Solution {
    public int searchInsert(int[] nums, int target) {
        int min, mid;
        int max = nums.length;
        min = mid = 0;

        while (min <= max) {
            mid = (min + max) / 2;
            if (target < nums[mid]) {
                max = mid - 1;
            } else if (target > nums[mid]) {
                min = mid + 1;
            } else {
                return mid;
            }
        }
        if (nums[mid] <= target) {
            return mid + 1;
        } else {
            return mid;
        }
    }
}