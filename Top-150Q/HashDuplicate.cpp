#include <bits/stdc++.h>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
public:
  // time complexity : O(n^2)
  void brute(vector<int> arr, int size) {
    bool isDuplicate = false;
    for(int i = 0 ; i < size ; i++){
      for(int j = i + 1; j < size; j++){
        if(arr[i] == arr[j]) {
          isDuplicate = true;
          break;
        }
      }
      if(isDuplicate) break;
    }
    cout << "isDuplicate -> " << ( isDuplicate ? "True" : "False" ) << endl;
  }

  static bool comp(int a , int b){
    return a < b;
  }

  // time complexity : O(n logn)
  void best(vector<int> arr, int size) {
    bool isDuplicate = false;

    sort(arr.begin(), arr.end(), comp);

    for(int i = 0 ; i < size; i++) {
      if(arr[i] == arr[i+1]) {
        isDuplicate = true;
      }
      if(isDuplicate) break;
    }

    cout << "isDuplicate -> " << ( isDuplicate ? "True" : "False" ) << endl;
  }

  // time complexity : O(n)
  void optimal(vector<int> arr, int size) {
    unordered_set<int> duplicate;

    for( int i = 0 ; i < size ; i++ ){
      if(duplicate.find(arr[i]) != duplicate.end()) {
        cout << "isDuplicate -> True" << endl;
        return;
      } else {
        duplicate.insert(arr[i]);
      }
    }
    cout << "isDuplicate -> False" << endl;
  }
};

int main() {
  Solution sol;

  vector<int> arr = { 1, 4, 3, 4, 5 };
  int size = arr.size();

  cout << "Brute force : ";
  sol.brute(arr, size);

  cout << "Best solution : ";
  sol.best(arr, size);

  cout << "optimal solution : ";
  sol.optimal(arr, size);

  return 0;
}
