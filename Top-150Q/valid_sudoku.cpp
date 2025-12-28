/*
   Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without
repetition. Note:

A Sudoku board (partially filled) could be valid but is not necessarily
solvable. Only the filled cells need to be validated according to the mentioned
rules.
   */

#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
public:
  // Brute force approach
  bool brute(vector<vector<char>> &board) {
    // Check rows
    for (int i = 0; i < 9; i++) {
      unordered_set<char> seen;
      for (int j = 0; j < 9; j++) {
        if (board[i][j] == '.')
          continue;
        if (seen.count(board[i][j]))
          return false;
        seen.insert(board[i][j]);
      }
    }

    // Check columns
    for (int j = 0; j < 9; j++) {
      unordered_set<char> seen;
      for (int i = 0; i < 9; i++) {
        if (board[i][j] == '.')
          continue;
        if (seen.count(board[i][j]))
          return false;
        seen.insert(board[i][j]);
      }
    }

    // Check 3x3 boxes
    for (int boxRow = 0; boxRow < 9; boxRow += 3) {
      for (int boxCol = 0; boxCol < 9; boxCol += 3) {
        unordered_set<char> seen;
        for (int i = 0; i < 3; i++) {
          for (int j = 0; j < 3; j++) {
            char val = board[boxRow + i][boxCol + j];
            if (val == '.')
              continue;
            if (seen.count(val))
              return false;
            seen.insert(val);
          }
        }
      }
    }

    return true;
  }

  // Optimal approach (single pass)
  bool optimal(vector<vector<char>> &board) {
    unordered_set<string> seen;

    for (int i = 0; i < 9; i++) {
      for (int j = 0; j < 9; j++) {
        if (board[i][j] == '.')
          continue;

        char num = board[i][j];

        string rowKey = string(1, num) + " in row " + to_string(i);
        string colKey = string(1, num) + " in col " + to_string(j);
        string boxKey = string(1, num) + " in box " + to_string(i / 3) + "-" +
                        to_string(j / 3);

        if (seen.count(rowKey) || seen.count(colKey) || seen.count(boxKey))
          return false;

        seen.insert(rowKey);
        seen.insert(colKey);
        seen.insert(boxKey);
      }
    }
    return true;
  }
};

int main() {
  Solution sol;

  vector<vector<char>> board = {{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
                                {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                                {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                                {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                                {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                                {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                                {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                                {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                                {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};

  vector<vector<char>> board2 = {{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
                                 {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                                 {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                                 {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                                 {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                                 {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                                 {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                                 {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                                 {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};

  bool ans1 = sol.brute(board);
  cout << "Valid Sudoku brute : ";
  if (ans1)
    cout << "true" << endl;
  else
    cout << "false" << endl;

  bool ans2 = sol.optimal(board2);
  cout << "Valid Sudoku optimal : ";
  if (ans2)
    cout << "true" << endl;
  else
    cout << "false" << endl;

  return 0;
}
