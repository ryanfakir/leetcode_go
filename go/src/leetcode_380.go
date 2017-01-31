package main

import "math/rand"

type RandomizedSet struct {
	M    map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), nil}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.M[val]; ok {
		return false
	}
	this.M[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.M[val]; !ok {
		return false
	}
	index := this.M[val]
	lastValue := this.list[len(this.list)-1]
	this.M[lastValue] = index
	this.list[index] = lastValue
	this.list = this.list[:len(this.list)-1]
	delete(this.M, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}
