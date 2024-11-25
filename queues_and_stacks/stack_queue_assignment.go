package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

type queue struct {
	val []int
}

type stack struct {
	val []int
}

func (Q *queue) Enqueue(data int) {
	Q.val = append(Q.val, data)
}

func (Q *queue) Dequeue() {
	Q.val = Q.val[1:]
}

func (S *stack) Push(data int) []int {
	c := S.val[0:]
	S.val = []int{data}
	for _, v := range c {
		S.val = append(S.val, v)
	}
	return S.val
}

func (S *stack) Pop() {
	S.val = S.val[1:]
}

func main() {

	queue := queue{}
	stack := stack{}

	input := bufio.NewReader(os.Stdin)
	r, err := input.ReadString('\n')
	var promptsendtoqueueandstack string
	

	if err != nil {
		return
	}
	
	for k,v := range r {
		_, e := strconv.Atoi(string(v))
        if e == nil {
            promptsendtoqueueandstack += string(r[k])
        } else {
			promptint, err := strconv.Atoi(promptsendtoqueueandstack)
			if err != nil {
				fmt.Print(err)
				return
			}
			queue.Enqueue(promptint)
			stack.Push(promptint)
			promptsendtoqueueandstack = ""
		}
	}

	fmt.Print("Stack output (LIFO): ")
	for _, v := range queue.val {
		if len(queue.val) > 1 {
		fmt.Print(v, ", ")
		queue.Dequeue()
		} else {
			fmt.Print(v)
			queue.Dequeue()
		}
	}

	fmt.Println("")
}

