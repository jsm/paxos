package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	name string
	cost uint
}

type itemPair struct {
	item1 item
	item2 item
	cost  uint
}

type itemSet []item

// Implements the sort.Sort interface
func (a itemSet) Len() int           { return len(a) }
func (a itemSet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a itemSet) Less(i, j int) bool { return a[i].cost < a[j].cost }

func main() {
	// Read command line values
	filename := os.Args[1]
	desiredSum, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}

	items := importItems(filename)
	sort.Sort(itemSet(items))

	i1, i2, success := findItemPair(items, uint(desiredSum))
	if !success {
		fmt.Println("Not possible")
		return
	}

	// Print the result
	fmt.Printf("%s %d, %s %d\n", i1.name, i1.cost, i2.name, i2.cost)
}

func importItems(filename string) []item {
	// Open File
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Read as a CSV
	r := csv.NewReader(f)

	items := []item{}

	// Iterate through each row, and convert to an item struct
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				return items
			}
			panic(err)
		}

		cost, err := strconv.ParseUint(strings.TrimLeft(record[1], " "), 10, 64)
		if err != nil {
			panic(err)
		}

		items = append(items, item{
			name: record[0],
			cost: uint(cost),
		})
	}
}

func findItemPair(items []item, desiredSum uint) (item, item, bool) {
	p1 := 0
	p2 := len(items) - 1

	// Keep track of the best pair so far in case we don't find an exact match
	// This is initialized to the first two items, as this is the worst pair
	best := itemPair{
		item1: items[0],
		item2: items[1],
		cost:  items[0].cost + items[1].cost,
	}

	// Exit early in the case that even the least pair is greater than the desired Sum
	if best.cost > desiredSum {
		return item{}, item{}, false
	}

	// Two pointers, that start and min value and max value, and work their way towards each other
	for p1 < p2 {
		i1, i2 := items[p1], items[p2]
		sum := i1.cost + i2.cost
		if sum == desiredSum {
			return i1, i2, true
		} else if sum < desiredSum {
			// If less than the disred sum, but still the closest we've encountered, keep track of it.
			if sum > best.cost {
				best = itemPair{
					item1: i1,
					item2: i2,
					cost:  sum,
				}
			}
			p1 += 1 // Increment left pointer
			continue
		} else if sum > desiredSum {
			p2 -= 1 // Decerement right pointer
			continue
		}
	}

	return best.item1, best.item2, true
}
