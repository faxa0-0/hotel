package main

import "fmt"

type MultiplicationTable struct {
	num int
}

func NewMultiplicationTable(num int) MultiplicationTable {
	return MultiplicationTable{
		num: num,
	}
}
func (t *MultiplicationTable) DrawMultiplicationTable() {
	maxNum := t.num

	fmt.Print("  ")
	for i := 1; i <= maxNum; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	for i := 1; i <= maxNum; i++ {
		fmt.Printf("%2d", i)
		for j := 1; j <= maxNum; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func main() {
	mt := NewMultiplicationTable(5)
	mt.DrawMultiplicationTable()
}
