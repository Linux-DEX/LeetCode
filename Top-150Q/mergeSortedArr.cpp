#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    // time complexity : O(n + m) + O(n + m)
    // space complexity : O(n + m)
    void best(vector<int>& nums1, int m, vector<int>& nums2, int n){
      vector<int> nums3((m+n), 0);
      int left = 0;
      int right = 0;
      int index = 0;

      while(left < m && right < n) {
        if(nums1[left] <= nums2[right]){
          nums3[index++] = nums1[left++];
        }
        else {
          nums3[index++] = nums2[right++];
        }
      }

      while(left < m) {
        nums3[index++] = nums1[left++];
      }

      while(right < n) {
        nums3[index++] = nums2[right++];
      }

      for(int i = 0; i < m+n; i++){
        nums1[i] = nums3[i];
      }
    }

    // time complexity : O(min(m, n)) + O(n log(n)) + O(m log(m))
    // space complexity : O(1)
    void optimal(vector<int>& nums1, int m, vector<int>& nums2, int n){
      int i = m - 1;
      int j = n - 1;
      int k = m + n - 1;

      while(i >= 0 && j >= 0){
        if(nums1[i] > nums2[j]) {
          nums1[k--] = nums1[i--];
        } else {
          nums1[k--] = nums2[j--];
        }
      }

      while(j >= 0){
        nums1[k--] = nums2[j--];
      }
    }
};

int main() {
  Solution sol;

  vector<int> nums1 = {1,2,3,0,0,0};
  vector<int> nums2 = {2, 5, 6};
  int m = 3;
  int n = 3;

  cout << "Best sol : ";
  sol.best(nums1, m, nums2, n);

  for(int i : nums1){
    cout << i << ", ";
  }
  cout << endl;

  nums1 = {1,2,3,0,0,0};
  nums2 = {2, 5, 6};

  cout << "Optimal sol : ";
  sol.optimal(nums1, m, nums2, n);

  for(int i : nums1){
    cout << i << ", ";
  }
  cout << endl;

  return 0;
}


/*
output: 
  Best sol : 1, 2, 2, 3, 5, 6, 
  Optimal sol : 1, 2, 2, 3, 5, 6, 
*/
