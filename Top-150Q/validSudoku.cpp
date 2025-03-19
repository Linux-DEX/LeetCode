#include <bits/stdc++.h> 
using namespace std;

class Solution {
  public: 
    // time complexity: O(n^2)
    // space complexity: O(3n) , O(n) , O(1)
    bool optimal(vector<vector<char>>& board) {
      vector<unordered_set<char>> rows(9), cols(9), boxes(9);

      for (int i = 0; i < 9; i++) {
        for (int j = 0; j < 9; j++) {
          char num = board[i][j];
          if (num == '.') continue;

          int boxIndex = (i / 3) * 3 + (j / 3);  

          if (rows[i].count(num) || cols[j].count(num) || boxes[boxIndex].count(num))
            return false;

          rows[i].insert(num);
          cols[j].insert(num);
          boxes[boxIndex].insert(num);
        }
      }
      return true;
    }
};

int main(){
  Solution sol;

  vector<vector<char>> board = {
    {'5','3','.','.','7','.','.','.','.'},
    {'6','.','.','1','9','5','.','.','.'},
    {'.','9','8','.','.','.','.','6','.'},
    {'8','.','.','.','6','.','.','.','3'},
    {'4','.','.','8','.','3','.','.','1'},
    {'7','.','.','.','2','.','.','.','6'},
    {'.','6','.','.','.','.','2','8','.'},
    {'.','.','.','4','1','9','.','.','5'},
    {'.','.','.','.','8','.','.','7','9'}
  };

  cout << "Optimal sol : ";
  bool result = sol.optimal(board);
  cout << (result ? "Valid" : "Invalid") << endl;

  return 0;
}

/*
output:
  Optimal sol : Valid
*/
