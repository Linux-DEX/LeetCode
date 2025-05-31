/*
Given a string s consisting of words and spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only. 
*/

#include <iostream> 
#include <string> 
#include <sstream>
using namespace std;

class Solution {
  public:
    // Time complexity : O(n)
    // Space complexity : O(n)
    int brute(string s) {
      stringstream ss(s);
      string word, lastWord;

      while (ss >> word) {
        lastWord = word;
      }

      return lastWord.length();
    }

    // Time complexity : O(n)
    // Space complexity : O(1)
    int best(string s) {
      int i = s.length() - 1;

      while (i >= 0 && s[i] == ' ') i--;

      int length = 0;
      while (i >= 0 && s[i] != ' ') {
        length++;
        i--;
      }

      return length;
    }
};

int main() {
  Solution sol;
  // string input = "   Hello World   ";
  // string input = "   fly me   to   the moon  ";
  string input = "luffy is still joyboy";
  cout << "Brute Sol : " << sol.brute(input) << endl;

  cout << "Best Sol : " << sol.best(input) << endl;

  return 0;
}
