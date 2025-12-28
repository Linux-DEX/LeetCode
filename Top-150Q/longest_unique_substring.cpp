/*
Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without duplicate characters.
*/

#include <iostream>
#include <unordered_set>
#include <string>
using namespace std;

int lengthOfLongestSubstring(string s) {
    unordered_set<char> seen;
    int left = 0, maxLength = 0;

    for (int right = 0; right < s.length(); ++right) {
        while (seen.count(s[right])) {
            seen.erase(s[left]);
            ++left;
        }
        seen.insert(s[right]);
        maxLength = max(maxLength, right - left + 1);
    }

    return maxLength;
}

int main() {
    string input;
    cout << "Enter a string: ";
    cin >> input;

    int result = lengthOfLongestSubstring(input);
    cout << "Length of longest substring without repeating characters: " << result << endl;

    return 0;
}
