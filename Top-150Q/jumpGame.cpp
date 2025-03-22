/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

#include <bits/stdc++.h> 
using namespace std;

class Solution {
  public:
    bool best(vector<int>& nums) {
      int maxJump = 0;
      int size = nums.size();

      for(int i = 0; i < size; i++) {
        if(i > maxJump) return false;
        maxJump = max(maxJump, i + nums[i]);
      }

      return true;
    }
};

int main() {
  Solution sol;

  vector<int> nums = {2,3,1,1,4};
  // vector<int> nums = {3,2,1,0,4};

  cout << "Best sol : " << (sol.best(nums) ? "True" : "False") << endl;

  return 0;
}

/*
output: 
  Best sol : True
*/
