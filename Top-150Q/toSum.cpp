#include <bits/stdc++.h>
#include <cstdio>
#include <map>
using namespace std;

class Solution {
public:
  pair<int, int> brute(vector<int> &arr, int target){
    int size = arr.size();
    for(int i = 0 ; i < size ; i++){
      if(arr[i] > target){
        continue;
      }

      int val = target - arr[i];
      for(int j = i + 1; j < size; j++){
        if(arr[j] == val){
          return { i, j };
        }
      }
    }

    return { -1, -1 };
  }

  static bool comp(int a, int b) {
    return a < b;
  }

  static int binarySearch(vector<int> arr, int start, int target) {
    int left = start, right = arr.size() - 1;
    while (left <= right) {
      int mid = left + (right - left) / 2;
      if (arr[mid] == target) {
        return mid;
      } else if (arr[mid] < target) {
        left = mid + 1;
      } else {
        right = mid - 1;
      }
    }
    return -1;
  }


  pair<int, int> best(vector<int> arr, int target){
    int size = arr.size();
    sort(arr.begin(), arr.end(), comp);

    for(int i = 0 ; i < size; i++){
      if(arr[i] > target){
        continue;
      }

      int val = target - arr[i];
      int index = binarySearch(arr, i+1, val);
      if(index != -1){
        return { i, index };
      }
    }

    return { -1, -1 };
  }

  pair<int, int> optimal(vector<int> arr, int target){
    int size = arr.size();
    unordered_map<int, int> hashmap;

    for(int i = 0; i < size; i++){
      int complement = target - arr[i];

      if(hashmap.find(complement) != hashmap.end()){
        return {hashmap[complement], i};
      }

      hashmap[arr[i]] = i;
    }

    return { -1, -1 };
  }
};

int main() {
  Solution sol;

  vector<int> arr = { 15, 7, 2, 11 };
  int target = 9;

  cout << "Brute force : ";
  pair<int, int> result = sol.brute(arr, target);
  cout << result.first << " " << result.second << endl;

  cout << "Best sol : ";
  result = sol.best(arr, target);
  cout << result.first << " " << result.second << endl;

  cout << "Optimal sol : ";
  result = sol.optimal(arr, target);
  printf("%d %d\n", result.first, result.second);

  return 0;
}
