/*
   Implement the RandomizedSet class:

   - RandomizedSet() Initializes the RandomizedSet object.
   - bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
   - bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
   - int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
   You must implement the functions of the class such that each function works in average O(1) time complexity.
   */

#include <bits/stdc++.h> 
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
  public:
    unordered_map<int, int> valueToIndex;
    vector<int> values;

    Solution() { }

    bool insert(int val) {
      if (valueToIndex.find(val) != valueToIndex.end()) {
        return false;
      }

      values.push_back(val);
      valueToIndex[val] = values.size() - 1;
      return true;
    }

    bool remove(int val) {
      if (valueToIndex.find(val) == valueToIndex.end()) {
        return false;
      }

      int index = valueToIndex[val];
      int lastElement = values.back();
      values[index] = lastElement;
      valueToIndex[lastElement] = index;

      values.pop_back();
      valueToIndex.erase(val);
      return true;
    }

    int getRandom() {
      int randomIndex = rand() % values.size();
      return values[randomIndex];
    }
};

int main() {
  Solution sol;

  cout << sol.insert(1) << endl;
  cout << sol.insert(2) << endl;
  cout << sol.insert(2) << endl;
  cout << sol.remove(1) << endl;
  cout << sol.getRandom() << endl;

  return 0;
}

/*
output:
    1
    1
    0
    1
    2
*/
