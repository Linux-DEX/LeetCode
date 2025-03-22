/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    int best(vector<int>& nums, int val) {
      int j = 0;

      for(int i = 0 ; i < nums.size() ; i++) {
        if(nums[i] != val){
          nums[j++] = nums[i];
        }
      }

      return j;
    }

    int optimal(vector<int>& nums, int val) {
      int size = nums.size();

      for(int i=0; i < size;) {
        if(nums[i] == val){
          nums[i] = nums[size-1];
          size--;
        } else {
          i++;
        }
      }

      return size;
    }
};

int main() {
  Solution sol;

  // vector<int> nums = { 3, 2, 2, 3 };
  // int val = 3;

  vector<int> nums = {0,1,2,2,3,0,4,2};
  int val = 2;

  cout << "Best sol : ";
  cout << sol.best(nums, val) << " -> ";
  for(int n: nums){
    cout << n << ", ";
  }
  cout << endl;

  nums = {0,1,2,2,3,0,4,2};
  cout << "Optimal sol : ";
  cout << sol.optimal(nums, val) << " -> ";
  for(int n: nums){
    cout << n << ", ";
  }
  cout << endl;

  return 0;
}

/*
output: 
  Best sol : 5 -> 0, 1, 3, 0, 4, 0, 4, 2, 
  Optimal sol : 5 -> 0, 1, 4, 0, 3, 0, 4, 2, 
*/
