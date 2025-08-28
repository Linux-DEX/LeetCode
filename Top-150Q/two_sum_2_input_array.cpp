/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.
*/

#include <iostream> 
#include <vector>
using namespace std;

class Solution {
  public: 
    vector<int> brute(vector<int>& numbers, int target) {
      for(int i=0; i < numbers.size(); i++) {
        for(int j=0; j < numbers.size(); j++) {
          if(numbers[i] + numbers[j] == target) {
            return {i+1, j+1};
          }
        }
      }
      return {};
    }

    vector<int> best(vector<int>& numbers, int target) {
      for (int i = 0; i < numbers.size(); ++i) {
        int low = i + 1, high = numbers.size() - 1;
        int complement = target - numbers[i];

        while (low <= high) {
          int mid = low + (high - low) / 2;
          if (numbers[mid] == complement) {
            return {i + 1, mid + 1};
          } else if (numbers[mid] < complement) {
            low = mid + 1;
          } else {
            high = mid - 1;
          }
        }
      }
      return {};
    }

    vector<int> optimal(vector<int>& numbers, int target) {
      int left = 0, right = numbers.size() - 1;

      while (left < right) {
        int sum = numbers[left] + numbers[right];

        if (sum == target) {
          return {left + 1, right + 1};
        } else if (sum < target) {
          left++;
        } else {
          right--;
        }
      }

      return {};  
    }
};

void printResult(const string& label, const vector<int>& result) {
    cout << label << " : [";
    for (int i = 0; i < result.size(); ++i) {
        cout << result[i];
        if (i != result.size() - 1) cout << ", ";
    }
    cout << "]" << endl;
}

int main() {
  Solution sol;

  vector<int> numbers = {2, 7, 11, 15};
  int target = 9;

  printResult("Brute sol", sol.brute(numbers, target));
  printResult("Best sol", sol.best(numbers, target));
  printResult("Optimal sol", sol.optimal(numbers, target));

  return 0;
}

