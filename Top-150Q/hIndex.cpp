/*
   Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.

   According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
   */

#include <bits/stdc++.h> 
#include <vector>
using namespace std;

class Solution {
  public:
    // time complexity: O(n log(n))
    // space complexity : O(1)
    int brute(vector<int>& citations) {
      int n = citations.size();
      sort(citations.begin(), citations.end(), greater<int>());  

      for (int h = 0; h < n; ++h) {
        if (citations[h] < h + 1) {
          return h;  
        }
      }

      return n;
    }

    // time complexity: O(n log(n))
    // space complexity : O(1)
    int best(vector<int>& citations) {
      int n = citations.size();
      sort(citations.begin(), citations.end());  

      int left = 0, right = n - 1;
      while (left <= right) {
        int mid = left + (right - left) / 2;
        if (citations[mid] >= n - mid) {
          right = mid - 1;  
        } else {
          left = mid + 1;
        }
      }

      return n - left;
    }

    int optimal(vector<int>& citations) {
      int n = citations.size();
      vector<int> freq(n + 1, 0);  

      for (int citation : citations) {
        if (citation >= n) {
          freq[n]++;
        } else {
          freq[citation]++;
        }
      }

      int count = 0;
      for (int h = n; h >= 0; --h) {
        count += freq[h];
        if (count >= h) {
          return h;  
        }
      }

      return 0;
    }
};

int main() {
  Solution sol;

  vector<int> citations = {3,0,6,1,5};
  // vector<int> citations = {1,3,1};
  cout << "Brute sol : " << sol.brute(citations) << endl;

  citations = {3,0,6,1,5};
  // citations = {1,3,1};
  cout << "Best sol : " << sol.best(citations) << endl;

  citations = {3,0,6,1,5};
  // citations = {1,3,1};
  cout << "Optimal sol : " << sol.optimal(citations) << endl;

  return 0;
}

/*
output: 
  Brute sol : 3
  Best sol : 3
  Optimal sol : 3
*/
