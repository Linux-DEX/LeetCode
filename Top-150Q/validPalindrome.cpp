/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

#include <iostream> 
#include <string>
using namespace std;

class Solution {
  public:
    // Time complexity : O(n) , Space complexity : O(n)
    bool brute(string s) {
      string clean = "";
      for (char c : s) {
        if (isalnum(c)) clean += tolower(c);
      }
      string rev = clean;
      reverse(rev.begin(), rev.end());
      return clean == rev;
    }

    // Time complexity : O(n) , Space complexity : O(1)
    bool optimal(string s) {
      int left = 0, right = s.length() - 1;
      while (left < right) {
        while (left < right && !isalnum(s[left])) left++;
        while (left < right && !isalnum(s[right])) right--;
        if (tolower(s[left]) != tolower(s[right])) return false;
        left++; right--;
      }
      return true;
    }

    // Time complexity : O(n) , Space complexity : O(1)
    bool best(string s) {
      for (int l = 0, r = s.length() - 1; l < r; ) {
        if (!isalnum(s[l])) { l++; continue; }
        if (!isalnum(s[r])) { r--; continue; }
        if (tolower(s[l]) != tolower(s[r])) return false;
        l++; r--;
      }
      return true;
    }
};

int main() {
  Solution sol;

  string str1 = "A man, a plan, a canal: Panama";
  string str2 = "race a car";
  string str3 = "";

  cout << "Brute:   " << (sol.brute(str1) ? "true" : "false") << endl;  
  cout << "Optimal: " << (sol.optimal(str2) ? "true" : "false") << endl; 
  cout << "Best:    " << (sol.best(str3) ? "true" : "false") << endl;    

  return 0;
}

