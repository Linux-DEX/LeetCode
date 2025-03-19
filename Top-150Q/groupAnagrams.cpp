#include <bits/stdc++.h>
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
  vector<vector<string>> best(vector<string>& strs) {
    unordered_map<string, vector<string>> anagramGroups;

    for (const string& str : strs) {
      string sortedStr = str;  
      sort(sortedStr.begin(), sortedStr.end()); 

      anagramGroups[sortedStr].push_back(str); 
    }

    vector<vector<string>> result;

    for (const auto& pair : anagramGroups) {
      result.push_back(pair.second);
    }

    return result;
  }

  vector<vector<string>> optimal(vector<string>& strs) {
    if (strs.empty()) {
      return {};
    }

    map<string, vector<string>> ansMap;

    for (const string& s : strs) {
      int count[26] = {0}; 

      for (char c : s) {
        count[c - 'a']++; 
      }

      string key;
      for (int i = 0; i < 26; i++) {
        key += "#" + to_string(count[i]);
      }

      ansMap[key].push_back(s);
    }

    vector<vector<string>> result;
    for (const auto& pair : ansMap) {
      result.push_back(pair.second);
    }

    return result;
  }

  void printArr(const vector<vector<string>> result){
    for(const auto& group: result) {
      for(const auto& word: group){
        cout << word << ", ";
      }
      cout << endl;
    }
  }
};

int main(){
  Solution sol;

  vector<string> arr = {"eat","tea","tan","ate","nat","bat"};

  cout << "Best solution : \n";
  vector<vector<string>> resultBest = sol.best(arr);
  sol.printArr(resultBest);

  cout << "Optimal sol : \n";
  vector<vector<string>> resultOpt = sol.optimal(arr);
  sol.printArr(resultOpt);

  return 0;
}
