/*
   Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

   You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

   Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

   For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:
- A word is defined as a character sequence consisting of non-space characters only.
- Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
- The input array words contains at least one word.
*/

#include <iostream> 
#include <string>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<string> brute(vector<string>& words, int maxWidth) {
      vector<string> res;
      int i = 0, n = words.size();

      while (i < n) {
        int j = i + 1, lineLen = words[i].length();

        while (j < n && lineLen + words[j].length() + (j - i) <= maxWidth) {
          lineLen += words[j].length();
          ++j;
        }

        int spaceSlots = j - i - 1;
        string line = words[i];

        if (j == n || spaceSlots == 0) {
          for (int k = i + 1; k < j; ++k)
            line += " " + words[k];
          line += string(maxWidth - line.length(), ' ');
        } else {
          int totalSpaces = maxWidth - lineLen;
          int space = totalSpaces / spaceSlots;
          int extra = totalSpaces % spaceSlots;

          for (int k = i + 1; k < j; ++k) {
            line += string(space + (extra-- > 0 ? 1 : 0), ' ') + words[k];
          }
        }

        res.push_back(line);
        i = j;
      }

      return res;
    }
};

int main() {
  Solution sol;
  vector<string> words = {"This", "is", "an", "example", "of", "text", "justification."};
  int maxWidth = 16;

  vector<string> result = sol.brute(words, maxWidth);
  for (auto& line : result) cout << "\"" << line << "\"" << endl;

  return 0;
}
