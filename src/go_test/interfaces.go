package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	/*	sort.Sort(s)
		str := "["
		for i, elem := range s {
			fmt.Println(i)
			fmt.Println(elem)
			if i > 0 {
				str += " "
			}
			str += fmt.Sprint(elem)
			fmt.Println(str)
		}
		return str + "]"
	*/
	//sort.Sort(s)
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main() {
	v := Sequence{1, 2, 3}
	fmt.Println(v)
}
