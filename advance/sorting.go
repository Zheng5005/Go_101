package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func sort_func() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	fmt.Println("Unsorted by age:", people)
	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(ageAsc).Sort(people)
	fmt.Println("Sorted by age (ascending):", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sorted by age (descending):", people)
	By(name).Sort(people)
	fmt.Println("Sorted by name:", people)
	By(lenName).Sort(people)
	fmt.Println("Sorted by length of name:", people)

	// ====== SORT.SLICE
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)

	//===================================
	// sort.Sort(ByAge(people))
	// numbers := []int{5, 3, 4, 1, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)

}

// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
