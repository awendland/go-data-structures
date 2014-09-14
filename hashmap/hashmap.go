package main

import (
	"fmt"
	"github.com/awendland/linkedlist"
	"time"
)

func hash(s string) int {
	hashcode, length := 0, len(s)
	for i := 0; i < length; i++ {
		hashcode = 31*hashcode + int(s[i])
	}
	return hashcode
}

type HashMap struct {
	buckets  []linkedlist.LinkedList
	size     int
	capacity float64
}

func New(length int, capacity float64) *HashMap {
	h := new(HashMap)
	h.buckets = make([]linkedlist.LinkedList, length, length)
	h.capacity = capacity
	return h
}

type HashMapEntry struct {
	key, val interface{}
}

func (h *HashMap) Put(key string, val interface{}) {
	hashcode := hash(key)
	index := hashcode % len(h.buckets)
	bucket := &h.buckets[index]
	existingKeyIndex := -1
	for index := 0; index < bucket.Length(); index++ {
		if bucket.Get(index).(HashMapEntry).key == key {
			fmt.Println("Key overwrite")
			existingKeyIndex = index
			break
		}
	}
	hmapEntry := HashMapEntry{key, val}
	if existingKeyIndex > -1 {
		bucket.Remove(existingKeyIndex)
		bucket.Insert(existingKeyIndex, hmapEntry)
	} else {
		bucket.Add(hmapEntry)
	}
}

func (h *HashMap) Get(key string) interface{} {
	hashcode := hash(key)
	index := hashcode % len(h.buckets)
	ll := &h.buckets[index]
	for i := 0; i < ll.Length(); i++ {
		entry := ll.Get(i).(HashMapEntry)
		if entry.key == key {
			return entry.val
		}
	}
	return nil
}

func main() {
	start := time.Now()
	h := New(12, .75)
	for i := 0; i < 37; i++ {
		fmt.Println(i)
		h.Put(string(i%25+65), i%25+65)
	}
	fmt.Println(h.buckets[5].TestGet(0))
	fmt.Println(h.buckets[5].ToString())
	// For some reason this causes the program to crash
	// h.Put(string(77), 87)
	fmt.Println(h)
	fmt.Println(h.Get(string(10 + 65)))
	fmt.Println(time.Since(start))
}
