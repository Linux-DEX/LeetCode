#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    bool brute(string s, int slen, string t, int tlen) {
      if (slen != tlen) return false;

      vector<bool> used(tlen, false);

      for (int i = 0; i < slen; i++) {
        bool found = false;
        for (int j = 0; j < tlen; j++) {
          if (!used[j] && s[i] == t[j]) {  
            used[j] = true;  
            found = true;
            break;
          }
        }
        if (!found) {
          return false;  
        }
      }

      return true;  
    }

    bool optimal(string s , int slen, string t, int tlen){
      if(slen != tlen) return false;

      int arr[26] = {0};
      const int a = 97;

      for(int i = 0 ; i < slen ; i++){
        arr[int(s[i]) - a]--;
        arr[int(t[i]) - a]++;
      }

      for(int a : arr) {
        if (a != 0){
          return false;
        }
      }

      return true;
    }
};

int main() {
  Solution sol;

  string s = "cat";
  string t = "tax";

  int slen = s.length();
  int tlen = t.length();

  cout << "Brute force : ";
  if(sol.brute(s, slen, t, tlen)){
    cout << "yes anagram" << endl;
  } else {
    cout << "not anagram" << endl;
  }

  cout << "Optimal solution : ";
  if(sol.optimal(s, slen, t, tlen)){
    cout << "yes anagram" << endl;
  } else {
    cout << "not anagram" << endl;
  }

  return 0;
}
