/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
  public:
    vector<vector<int>> threeSum(vector<int>& nums) {
      sort(nums.begin(), nums.end());
      vector<vector<int>> result;

      for (int i = 0; i < nums.size(); ++i) {
        if (i > 0 && nums[i] == nums[i - 1]) continue;

        int left = i + 1;
        int right = nums.size() - 1;

        while (left < right) {
          int sum = nums[i] + nums[left] + nums[right];

          if (sum == 0) {
            result.push_back({nums[i], nums[left], nums[right]});

            while (left < right && nums[left] == nums[left + 1]) left++;
            while (left < right && nums[right] == nums[right - 1]) right--;

            left++;
            right--;
          } else if (sum < 0) {
            left++; 
          } else {
            right--; 
          }
        }
      }

      return result;
    }
};

int main() {
  Solution sol;
  vector<int> nums = {-1, 0, 1, 2, -1, -4};

  vector<vector<int>> result = sol.threeSum(nums);

  cout << "Unique triplets:\n";
  for (auto& triplet : result) {
    cout << "[" << triplet[0] << ", " << triplet[1] << ", " << triplet[2] << "]\n";
  }

  return 0;
}

