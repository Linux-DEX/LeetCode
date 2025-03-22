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
    // space complexity : O(1)
    int best(vector<int>& prices) {
      int profit = 0;
      int size = prices.size();
      for (int i = 1; i < size; i++) {
        if (prices[i] > prices[i - 1]) {
          profit += prices[i] - prices[i - 1]; 
        }
      }
      return profit;
    }
};

int main() {
  Solution sol;

  // vector<int> prices = {7,1,5,3,6,4};
  vector<int> prices = {1,2,3,4,5};

  cout << "Best sol : " ;
  int bestRes = sol.best(prices); 
  cout << bestRes << endl;

  return 0;
}

/*
output:
  Best sol : 4
*/
