/*
mplement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.

Example 1:

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	valueToIndex map[int]int
	values       []int
}

// -----------------------------------
// Constructor
// -----------------------------------
func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano()) // seed randomness
	return RandomizedSet{
		valueToIndex: make(map[int]int),
		values:       []int{},
	}
}

// -----------------------------------
// Insert - O(1)
// -----------------------------------
func (r *RandomizedSet) Insert(val int) bool {
	if _, exists := r.valueToIndex[val]; exists {
		return false
	}

	r.values = append(r.values, val)
	r.valueToIndex[val] = len(r.values) - 1
	return true
}

// -----------------------------------
// Remove - O(1)
// -----------------------------------
func (r *RandomizedSet) Remove(val int) bool {
	index, exists := r.valueToIndex[val]
	if !exists {
		return false
	}

	lastElement := r.values[len(r.values)-1]

	// swap with last
	r.values[index] = lastElement
	r.valueToIndex[lastElement] = index

	// remove last
	r.values = r.values[:len(r.values)-1]
	delete(r.valueToIndex, val)

	return true
}

// -----------------------------------
// GetRandom - O(1)
// -----------------------------------
func (r *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(r.values))
	return r.values[randomIndex]
}

// -----------------------------------
// Main
// -----------------------------------
func main() {
	sol := Constructor()

	fmt.Println(sol.Insert(1))
	fmt.Println(sol.Insert(2))
	fmt.Println(sol.Insert(2))
	fmt.Println(sol.Remove(1))
	fmt.Println(sol.GetRandom())
}
