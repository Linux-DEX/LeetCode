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
