/*
- Minimum size subarray sum

- Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
*/

#include<iostream>
#include<vector>
#include<climits> 

using namespace std;

int minSubArrayLen(int target, vector<int>& nums) {
    int n = nums.size();
    int minLength = INT_MAX;
    int sum = 0, left = 0;

    for (int right = 0; right < n; ++right) {
        sum += nums[right];

        while (sum >= target) {
            minLength = min(minLength, right - left + 1);
            sum -= nums[left];
            ++left;
        }
    }

    return (minLength == INT_MAX) ? 0 : minLength;
}

int main() {
    vector<int> nums = {2, 3, 1, 2, 4, 3};
    int target = 7;

    int result = minSubArrayLen(target, nums);
    cout << "Minimum length of subarray: " << result << endl;

    return 0;
}

