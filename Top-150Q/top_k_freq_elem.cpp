#include <bits/stdc++.h> 
#include <queue>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<int> best(vector<int>& nums, int k) {
      unordered_map<int, int> freqMap;
      for (int num : nums) {
        freqMap[num]++;
      }

      vector<vector<int>> buckets(nums.size() + 1); 
      for (const auto& entry : freqMap) {
        buckets[entry.second].push_back(entry.first);
      }

      vector<int> result;
      for (int i = buckets.size() - 1; i >= 0 && result.size() < k; --i) {
        for (int num : buckets[i]) {
          result.push_back(num);
          if (result.size() == k) break;
        }
      }
      return result;
    }

    vector<int> optimal(vector<int>& nums, int k) {
      unordered_map<int, int> freqMap;
      for (int num : nums) {
        freqMap[num]++;
      }

      priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> minHeap;

      for (const auto& entry : freqMap) {
        minHeap.push({entry.second, entry.first});
        if (minHeap.size() > k) {
          minHeap.pop();
        }
      }

      vector<int> result;
      while (!minHeap.empty()) {
        result.push_back(minHeap.top().second);
        minHeap.pop();
      }

      return result;
    }
};

int main() {
  Solution sol;

  vector<int> nums = {1,1,1,2,2,3};
  int k = 2;

  cout << "Best sol : ";
  vector<int> ans = sol.best(nums, k);

  for(int n : ans){
    cout << n << ", ";
  }
  cout << endl;

  vector<int> optimalAns = sol.optimal(nums, k);
  cout << "Optimal sol : ";
  for (int n : optimalAns) {
    cout << n << " ";
  }
  cout << endl;

  return 0;
}
