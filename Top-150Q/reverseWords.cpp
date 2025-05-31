/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    string brute(string s) {
      stringstream ss(s);
      string word;
      vector<string> words;

      while(ss >> word) {
        words.push_back(word);
      }

      reverse(words.begin(), words.end());

      string result;
      for (int i = 0; i < words.size(); ++i) {
        result += words[i];
        if (i != words.size() - 1) result += " ";
      }

      return result;
    }

    string best(string s) {
      stringstream ss(s);
      string word;
      stack<string> st;

      while (ss >> word) {
        st.push(word);
      }

      string result;
      while (!st.empty()) {
        result += st.top();
        st.pop();
        if (!st.empty()) result += " ";
      }

      return result;
    }

    string optimal(string s) {
      int n = s.size();
      int i = 0;

      while (i < n && s[i] == ' ') i++; 
      int j = n - 1;
      while (j >= 0 && s[j] == ' ') j--; 

      string temp;
      while (i <= j) {
        if (s[i] != ' ') {
          temp += s[i];
        } else if (!temp.empty() && temp.back() != ' ') {
          temp += ' ';
        }
        i++;
      }

      reverse(temp.begin(), temp.end());

      int start = 0;
      for (int end = 0; end <= temp.size(); ++end) {
        if (end == temp.size() || temp[end] == ' ') {
          reverse(temp.begin() + start, temp.begin() + end);
          start = end + 1;
        }
      }

      return temp;
    }
};

int main() {
  Solution sol;

  // string input = "the sky is blue";
  // string input = "  hello world  ";
  string input = "a good   example";

  cout << "Brute sol : " << sol.brute(input) << endl;
  cout << "Best sol : " << sol.best(input) << endl;
  cout << "Optimal sol : " << sol.optimal(input) << endl;

  return 0;
}
