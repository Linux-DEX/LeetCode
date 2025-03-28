/*
   There are n gas stations along a circular route, where the amount of gas at the
   ith station is gas[i].

   You have a car with an unlimited gas tank and it costs cost[i] of gas to travel
   from the ith station to its next (i + 1)th station. You begin the journey with
   an empty tank at one of the gas stations.

   Given two integer arrays gas and cost, return the starting gas station's index
   if you can travel around the circuit once in the clockwise direction, otherwise
   return -1. If there exists a solution, it is guaranteed to be unique.
   */

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    // time complexity : O(n^2)
    int best(vector<int> &gas, vector<int> &cost) {
      int n = gas.size();
      for (int start = 0; start < n; start++) {
        int currentGas = 0;
        bool canComplete = true;
        for (int i = 0; i < n; ++i) {
          int idx = (start + i) % n;
          currentGas += gas[idx];
          currentGas -= cost[idx];

          if (currentGas < 0) {
            canComplete = false;
            break;
          }
        }
        if (canComplete) {
          return start;
        }
      }
      return -1;
    }

    // time complexity : O(n)
    int optimal(vector<int> &gas, vector<int> &cost) {
      int n = gas.size();
      int totalGas = 0, totalCost = 0;
      int currentGas = 0;
      int start = 0;

      for (int i = 0; i < n; ++i) {
        totalGas += gas[i];
        totalCost += cost[i];
        currentGas += gas[i] - cost[i];

        if (currentGas < 0) {
          start = i + 1;
          currentGas = 0;
        }
      }

      return (totalGas >= totalCost) ? start : -1;
    }
};

int main() {
  Solution sol;

  vector<int> gas = {1, 2, 3, 4, 5};
  vector<int> cost = {3, 4, 5, 1, 2};

  cout << "Best sol : ";
  int bestRes = sol.best(gas, cost);
  cout << bestRes << endl;

  cout << "Optimal sol : ";
  int optimalRes = sol.optimal(gas, cost);
  cout << optimalRes << endl;
  return 0;
}

/*
output: 
  Best sol : 3
  Optimal sol : 3
*/
