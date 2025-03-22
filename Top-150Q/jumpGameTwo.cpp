/*
   You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

   Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

   0 <= j <= nums[i] and
   i + j < n
   Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
   */

#include <bits/stdc++.h> 
using namespace std;

class Solution {
  public:
    int minJumpsHelper(vector<int>& nums, int currentPos, vector<int>& memo) {
      int n = nums.size();
      if (currentPos >= n - 1) return 0;  

      if (memo[currentPos] != -1) return memo[currentPos];  

      int jumps = INT_MAX;
      for (int i = 1; i <= nums[currentPos]; ++i) {
        if (currentPos + i < n) {
          int result = minJumpsHelper(nums, currentPos + i, memo);
          if (result != INT_MAX) {
            jumps = min(jumps, 1 + result);
          }
        }
      }

      memo[currentPos] = jumps;
      return jumps; 
    }

    // time complexity : O(n^2)
    // space complexity : O(n)
    int brute(vector<int>& nums) {
      int n = nums.size();
      vector<int> memo(n, -1);  
      return minJumpsHelper(nums, 0, memo);
    }

    // time complexity : O(n)
    // space complexity : O(1)
    int optimal(vector<int>& nums) {
      int n = nums.size();
      if (n <= 1) return 0;

      int jumps = 0;
      int current_end = 0;
      int farthest = 0;

      for (int i = 0; i < n; ++i) {
        farthest = max(farthest, i + nums[i]);

        if (i == current_end) {
          jumps++;
          current_end = farthest;

          if (current_end >= n - 1) {
            break;
          }
        }
      }

      return jumps;
    }
};

int main() {
  Solution sol;

  vector<int> nums = {2,3,1,1,4};

  // recursive DFS
  cout << "Brute sol : " << sol.brute(nums) << endl;

  // Greedy Approach
  cout << "Optimal sol : " << sol.optimal(nums) << endl;

  return 0;
}

/*
output: 
  Best sol : 2
*/
