/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

#include <bits/stdc++.h>
using namespace std;

class Solution {
  public: 
    int best(vector<int>& prices) {
      int mini = prices[0];
      int maxProfit = 0;
      int n = prices.size();

      for(int i=0; i<n; i++) {
        int cost = prices[i] - mini;
        maxProfit = max(maxProfit, cost);
        mini = min(mini, prices[i]);
      }

      return maxProfit;
    }
};

int main() {
  Solution sol;

  vector<int> prices = {7,1,5,3,6,4};
  // vector<int> prices = {7,6,4,3,1};

  cout << "Best sol : " ;
  int bestRes = sol.best(prices); 
  cout << bestRes << endl;

  return 0;
}

/*
output: 
  Best sol : 5
*/
