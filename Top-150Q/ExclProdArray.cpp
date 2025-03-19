#include <bits/stdc++.h> 
#include <vector>
using namespace std;

class Solution {
  public: 
    // time complexity: O(n^2)
    // space complexity: O(n)
    vector<int> brute(vector<int>& nums) {
      int size = nums.size();
      vector<int> newVec(size, 0);

      for(int i = 0; i < size; i++){
        int prod = 1;
        for(int j = 0; j < size; j++){
          if(i==j) continue;
          prod *= nums[j];
        }
        newVec[i] = prod;
      }

      return newVec;
    }

    // time complexity: O(n x 3)
    // space complexity: O(n)
    vector<int> best(vector<int>& nums) {
      int size = nums.size();
      vector<int> prefix(size, 1);
      vector<int> postfix(size, 1);
      vector<int> result(size, 1);

      for (int i = 1; i < size; i++) {
        prefix[i] = prefix[i - 1] * nums[i - 1];
      }

      for (int i = size - 2; i >= 0; i--) {
        postfix[i] = postfix[i + 1] * nums[i + 1];
      }

      for (int i = 0; i < size; i++) {
        result[i] = prefix[i] * postfix[i];
      }

      return result;
    }

    // time complexity: O(n x 2)
    // space complexity: O(1)
    vector<int> optimal(vector<int>& nums){
      int size = nums.size();
      vector<int> result(size, 1);

      int prefix = 1;
      for (int i = 0; i < size; i++) {
        result[i] = prefix; 
        prefix *= nums[i];  
      }

      int postfix = 1;
      for (int i = size - 1; i >= 0; i--) {
        result[i] *= postfix;
        postfix *= nums[i];   
      }

      return result;
    }
};

int main() {
  Solution sol;

  vector<int> arr = { 1, 2, 3, 4 };

  cout << "Brute sol : ";
  vector<int> result = sol.brute(arr);
  for(int r: result) {
    cout << r << ", ";
  }
  cout << endl;

  cout << "Best sol : ";
  result = sol.best(arr);
  for(int r: result) {
    cout << r << ", ";
  }
  cout << endl;

  cout << "Optimal sol : ";
  result = sol.optimal(arr);
  for(int r: result) {
    cout << r << ", ";
  }
  cout << endl;

  return 0;
}

/*
output: 
  Brute sol : 24, 12, 8, 6, 
  Best sol : 24, 12, 8, 6, 
  Optimal sol : 24, 12, 8, 6, 
*/
