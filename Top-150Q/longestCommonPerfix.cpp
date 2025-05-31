/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string ""
*/

#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
  public:
    // Time Complexity : O(N*M)
    string brute(vector<string>& strs) {
      if (strs.empty()) return "";

      string prefix = strs[0];
      for( int i = 1; i < strs.size(); ++i ) {
        int j = 0;
        while( j < prefix.size() && j < strs[i].size() && prefix[j] == strs[i][j]) {
          ++j;
        }
        prefix = prefix.substr(0,j);
        if (prefix.empty()) return "";
      }

      return prefix;
    }

    // Time Complexity : O(n*L)
    string optimal(vector<string>& strs) {
      if (strs.empty()) return "";

      for (int i = 0; i < strs[0].size(); ++i) {
        char c = strs[0][i];
        for (int j = 1; j < strs.size(); ++j) {
          if (i >= strs[j].size() || strs[j][i] != c) {
            return strs[0].substr(0, i);
          }
        }
      }
      return strs[0];
    }
};

int main() {
  Solution sol;

  // vector<string> input = { "flower", "flow", "flight" };
  vector<string> input = {"dog","racecar","car"};

  cout << "Brute force : " << sol.brute(input) << endl;
  cout << "Optimal force : " << sol.optimal(input) << endl;

  return 0;
}
