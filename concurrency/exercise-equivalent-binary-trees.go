package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*

1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false

*/

func Walk(t *tree.Tree, ch chan int) {
	walkRecur(t, ch)
	close(ch)
}

func walkRecur(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkRecur(t.Left, ch)
	ch <- t.Value
	walkRecur(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}
	return true
}

func main5() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for n := range ch {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
