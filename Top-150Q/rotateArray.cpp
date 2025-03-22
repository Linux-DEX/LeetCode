/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    void printArr(vector<int>& nums) {
      for (int i=0; i<nums.size(); i++){
        cout << nums[i] << ", ";
      }
      cout << endl;
    }

    // time complexity : O(n x k)
    // space complexity : O(1)
    void brute(vector<int>& nums, int k) {
      int size = nums.size();
      k = k % size;

      for(int i=0; i<k; i++) {
        int last = nums[size-1];
        for(int j=size-1; j>0; --j) {
          nums[j] = nums[j-1];
        }
        nums[0] = last;
      }
    }

    // time complexity : O(n)
    // space complexity : O(n)
    void best(vector<int>& nums, int k) {
      int size = nums.size();
      k = k % size;

      vector<int> rotate(size);

      for(int i=0; i<size; i++) {
        rotate[(i+k) % 7] = nums[i];
      }

      for(int i=0; i<size; i++) {
        nums[i] = rotate[i];
      }
    }

    void reverse(vector<int>& nums, int start, int end) {
      while (start < end) {
        swap(nums[start], nums[end]);
        start++;
        end--;
      }
    }

    // time complexity : O(n)
    // space complexity : O(1)
    void optimal(vector<int>& nums, int k) {
      int size = nums.size();
      k = k % size; 
      if (k == 0) return;  

      reverse(nums, 0, size - 1);
      reverse(nums, 0, k - 1);
      reverse(nums, k, size - 1);
    }
};

int main() {
  Solution sol;

  vector<int> nums = {1,2,3,4,5,6,7};
  int k = 3;

  cout << "Brute sol : ";
  sol.brute(nums, k);
  sol.printArr(nums);

  nums = {1,2,3,4,5,6,7};

  cout << "Best sol : ";
  sol.best(nums, k);
  sol.printArr(nums);

  nums = {1,2,3,4,5,6,7};

  cout << "Optimal sol : ";
  sol.optimal(nums, k);
  sol.printArr(nums);

  return 0;
}


/*
output: 
  Brute sol : 5, 6, 7, 1, 2, 3, 4, 
  Best sol : 5, 6, 7, 1, 2, 3, 4, 
  Optimal sol : 5, 6, 7, 1, 2, 3, 4, 
*/
