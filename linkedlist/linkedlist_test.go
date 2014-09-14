package linkedlist

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ll := LinkedList{}
	ll.Add(9)
	ll.Insert(1, 10)
	ll.Add("hello")
	fmt.Println(ll.ToString())
	fmt.Println(ll.Remove(0).(int))
	fmt.Println(ll.ToString())
}
