class Solution {
    public int search(int[] nums, int target) {
        int min = 0;
        int max = nums.length - 1;

        while (min <= max) {
            int mid = (min + max) / 2;
            if (target < nums[mid]) {
                max = mid - 1;
            } else if (target > nums[mid]) {
                min = mid + 1;
            } else {
                return mid;
            }
        }

        return -1;
    }
}