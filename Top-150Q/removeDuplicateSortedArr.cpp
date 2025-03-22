/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    // time complexity : O(n)
    // space complexity: O(n)
    int brute(vector<int>& nums) {
      if(nums.empty()) return 0;

      vector<int> unique;

      for(int i=0; i < nums.size(); i++){
        if(unique.empty() || nums[i] != unique.back()) {
          unique.push_back(nums[i]);
        }
      }

      for(int i=0; i<unique.size(); i++){
        nums[i] = unique[i];
      }

      return unique.size();
    }

    // time complexity : O(n)
    // space complexity : O(1)
    int optimal(vector<int>& nums) {
      if(nums.empty()) return 0;

      int i = 0;
      for(int j=0; j < nums.size(); j++) {
        if(nums[j] != nums[i]) {
          i++;
          nums[i] = nums[j];
        }
      }

      return i+1;
    }
};

int main() {
  Solution sol;

  // vector<int> nums = { 1,1,2 };
  vector<int> nums = {0,0,1,1,1,2,2,3,3,4};

  cout << "Brute sol : ";
  int bruteRes = sol.brute(nums);
  cout << bruteRes << " =>  ";
  for(int n : nums) {
    cout << n << ", ";
  }
  cout << endl;

  // nums = { 1, 1, 2 };
  nums = {0,0,1,1,1,2,2,3,3,4};


  cout << "Optimal sol : ";
  int optimalRes = sol.optimal(nums);
  cout << optimalRes << " =>  ";
  for(int n : nums) {
    cout << n << ", ";
  }
  cout << endl;

  return 0;
}

/*
output: 
  Brute sol : 5 =>  0, 1, 2, 3, 4, 2, 2, 3, 3, 4, 
  Optimal sol : 5 =>  0, 1, 2, 3, 4, 2, 2, 3, 3, 4, 
*/
