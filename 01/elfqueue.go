package main

import "container/heap"

type Elf struct {
	calories int
	index    int // The index of the item in the heap.
}

type ElfQueue []*Elf

func (eq ElfQueue) Len() int { return len(eq) }
func (eq ElfQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, calories so we use greater than here.
	return eq[i].calories > eq[j].calories
}
func (eq ElfQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
	eq[i].index = i
	eq[j].index = j
}

func (eq *ElfQueue) Push(x any) {
	n := len(*eq)
	elf := x.(*Elf) // check that x is of type *Elf and extract it, panic otherwise
	elf.index = n
	*eq = append(*eq, elf)
}

func (eq *ElfQueue) Pop() any {
	old := *eq
	n := len(old)
	elf := old[n-1]
	old[n-1] = nil // avoid memory leak
	elf.index = -1 // for safety
	*eq = old[0 : n-1]
	return elf
}

// update modifies the calories and value of an Elf in the queue.
func (eq *ElfQueue) update(elf *Elf, value string, calories int) {
	elf.calories = calories
	heap.Fix(eq, elf.index)
}
