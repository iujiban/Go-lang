package main

import ("fmt"
		"sort"
)

type Student struct {
	Name string
	Age int
}
type Students []Student

func (s Students) Len() int {return len(s)}
func (s Students) Less(i,j int) bool {return s[i].Age < s[j].Age}

func (s Students) Swap(i ,j int) {s[i], s[j] = s[j], s[i]}

func main() {
	s := []Student{
		{"화랑", 31} , {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18}}

		sort.Sort(Students(s))
		fmt.Println(s)
	/*
		var slice = []int{1, 2, 3}
		var totalSum int
		for i := 0; i < len(slice); i++ {
			totalSum += slice[i]

		}
		slice2 := append(slice, 4)
		fmt.Println(totalSum)
		fmt.Println(slice2)
	
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]

	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println((slice))
	*/

}
