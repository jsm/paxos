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

func (a itemSet) Len() int           { return len(a) }
func (a itemSet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a itemSet) Less(i, j int) bool { return a[i].cost < a[j].cost }

func main() {

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

	fmt.Printf("%s %d, %s %d\n", i1.name, i1.cost, i2.name, i2.cost)
}

func importItems(filename string) []item {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	items := []item{}

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

	best := itemPair{
		item1: items[0],
		item2: items[1],
		cost:  items[0].cost + items[1].cost,
	}

	if best.cost > desiredSum {
		return item{}, item{}, false
	}

	for p1 < p2 {
		i1, i2 := items[p1], items[p2]
		sum := i1.cost + i2.cost
		if sum == desiredSum {
			return i1, i2, true
		} else if sum < desiredSum {
			if sum > best.cost {
				best = itemPair{
					item1: i1,
					item2: i2,
					cost:  sum,
				}
			}
			p1 += 1
			continue
		} else if sum > desiredSum {
			p2 -= 1
			continue
		}
	}

	return best.item1, best.item2, true
}
