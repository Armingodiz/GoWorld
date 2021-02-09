package main

import (
	"fmt"
)

type Box struct {
	size        int
	containings []int
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	var (
		boxes []Box
	)
	for i := 0; i < n; i++ {
		tss := []int{i + 1}
		boxes = append(boxes, Box{size: 1, containings: tss})
	}
	/*for i := 0; i < n; i++ {
	  fmt.Println(boxes[i].containings)
	}*/
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		for j := 0; j < boxes[x-1].size; j++{
			boxes[y-1].containings = append(boxes[y-1].containings, boxes[x-1].containings[j])
		}
		boxes[y-1].size += boxes[x-1].size
		newBox := []int{1}
		boxes[x-1].containings = newBox
		boxes[x-1].size = 0
	}
	var d int
	_, _ = fmt.Scan(&d)
	fmt.Print(boxes[d-1].size)
	fmt.Print(" ")
	for i := 0; i < boxes[d-1].size; i++ {
		fmt.Print(boxes[d-1].containings[i])
		fmt.Print(" ")
	}
}
