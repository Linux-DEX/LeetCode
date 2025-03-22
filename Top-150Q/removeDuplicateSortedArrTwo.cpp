/*
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    // time complexity : O(n)
    // space complexity: O(n)
    int brute(vector<int>& nums) {
      if(nums.empty()) return 0;

      vector<int> result;

      for(int num: nums){
        if(result.size() < 2 || result[result.size() - 2] != num) {
          result.push_back(num);
        }
      }

      for(int i=0; i<result.size(); i++) {
        nums[i] = result[i];
      }

      return result.size();
    }

    // time complexity : O(n)
    // space complexity : O(1)
    int optimal(vector<int>& nums) {
      if(nums.empty()) return 0;

      int i = 1;  
      int count = 1;  

      for (int j = 1; j < nums.size(); j++) {
        if (nums[j] == nums[j - 1]) {
          count++;
        } else {
          count = 1;
        }

        if (count <= 2) {
          nums[i] = nums[j];
          i++;
        }
      }

      return i;
    }
};

int main() {
  Solution sol;

  // vector<int> nums = {1,1,1,2,2,3};
  vector<int> nums = {0,0,1,1,1,1,2,3,3};

  cout << "Brute sol : ";
  int bruteRes = sol.brute(nums);
  cout << bruteRes << " =>  ";
  for(int n : nums) {
    cout << n << ", ";
  }
  cout << endl;

  // nums = {1,1,1,2,2,3};
  nums = {0,0,1,1,1,1,2,3,3};

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
  Brute sol : 7 =>  0, 0, 1, 1, 2, 3, 3, 3, 3, 
  Optimal sol : 7 =>  0, 0, 1, 1, 2, 3, 3, 3, 3, 
*/
