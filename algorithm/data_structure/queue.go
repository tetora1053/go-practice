package main

import (
	"fmt"
)

var (
	head, tail int
)

type p struct {
	name string
	time int
}

type queue []p

func (q queue) enqueue(p p) {
	q[tail] = p
	if tail + 1 == len(q) {
		tail = 0
		return
	}
	tail++
}

func (q queue) dequeue() (x p) {
	x = q[head]
	if head+1 == len(q) {
		head = 0
	} else {
		head++
	}
	return x
}

func main() {
	var (
		n, qtm, time, elapsed int
		name string
	)
	fmt.Scan(&n, &qtm)
	q := make(queue, n+1)
	for i:=1; i<=n; i++ {
		fmt.Scan(&name, &time)
		q[i].name = name
		q[i].time = time
	}
	head = 1
	tail = 0
	for {
		if head == tail {
			break
		}
		p := q.dequeue()
		if qtm >= p.time {
			elapsed += p.time
			fmt.Println(p.name, elapsed)
			continue
		}
		p.time -= qtm
		q.enqueue(p)
		elapsed += qtm
	}
}

