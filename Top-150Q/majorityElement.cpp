/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/

#include <bits/stdc++.h> 
using namespace std;

class Solution {
  public:
    // time complexity : O(n^2)
    int brute(vector<int>& nums) {
      int n = nums.size();

      for(int i=0; i<n; i++) {
        int count = 0;
        for(int j=0; j<n; j++) {
          if(nums[i] == nums[j]) {
            count++;
          }
        }
        if(count > n/2) {
          return nums[i];
        }
      }

      return -1;
    }

    // time complexity : O(n)
    // space complexity : O(n)
    int best(vector<int>& nums) {
      unordered_map<int, int> count_map;
      int n = nums.size();

      for(int num: nums){ 
        count_map[num]++;
        if(count_map[num] > n/2) {
          return num;
        }
      }

      return -1;
    }

    // time complexity : O(n)
    // space complexity : O(1)
    int optimal(vector<int>& nums) {
      int count = 0;
      int candidate = -1;

      for(int num : nums) {
        if(count == 0) {
          candidate = num;
        }
        count += (num == candidate) ? 1 : -1 ;
      }

      return candidate;
    }
};

int main() {
  Solution sol;

  vector<int> nums = {3, 2, 3};
  // vector<int> nums = {4,4,1,1,1,4,4};

  cout << "Brute sol : ";
  int bruteRes = sol.brute(nums);
  cout << bruteRes << endl;

  cout << "Best sol : ";
  int bestRes = sol.best(nums);
  cout << bestRes << endl;

  cout << "Optimal sol : ";
  int optimalRes = sol.optimal(nums);
  cout << optimalRes << endl;

  return 0;
}

/*
output: 
  Brute sol : 3
  Best sol : 3
  Optimal sol : 3
*/
